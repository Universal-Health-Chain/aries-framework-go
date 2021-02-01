/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package jose

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/tink/go/keyset"

	cryptoapi "github.com/Universal-Health-Chain/aries-framework-go/pkg/crypto"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/crypto/tinkcrypto/primitive/composite"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/crypto/tinkcrypto/primitive/composite/api"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/crypto/tinkcrypto/primitive/composite/ecdh"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/crypto/tinkcrypto/primitive/composite/keyio"
	ecdhpb "github.com/Universal-Health-Chain/aries-framework-go/pkg/crypto/tinkcrypto/primitive/proto/ecdh_aead_go_proto"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/kms"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/storage"
)

// Decrypter interface to Decrypt JWE messages.
type Decrypter interface {
	// Decrypt a deserialized JWE, extracts the corresponding recipient key to decrypt plaintext and returns it
	Decrypt(jwe *JSONWebEncryption) ([]byte, error)
}

// JWEDecrypt is responsible for decrypting a JWE message and returns its protected plaintext.
type JWEDecrypt struct {
	// store is required for Authcrypt/ECDH1PU only (Anoncrypt doesn't need it as the sender is anonymous)
	store  storage.Store
	crypto cryptoapi.Crypto
	kms    kms.KeyManager
}

// NewJWEDecrypt creates a new JWEDecrypt instance to parse and decrypt a JWE message for a given recipient
// store is needed for Authcrypt only (to fetch sender's pre agreed upon public key), it is not needed for Anoncrypt.
func NewJWEDecrypt(store storage.Store, c cryptoapi.Crypto, k kms.KeyManager) *JWEDecrypt {
	return &JWEDecrypt{
		store:  store,
		crypto: c,
		kms:    k,
	}
}

func getECDHDecPrimitive(cek []byte, encAlg string) (api.CompositeDecrypt, error) {
	kt := ecdh.NISTPECDHAES256GCMKeyTemplateWithCEK(cek)

	if encAlg == XC20PALG {
		kt = ecdh.X25519ECDHXChachaKeyTemplateWithCEK(cek)
	}

	kh, err := keyset.NewHandle(kt)
	if err != nil {
		return nil, err
	}

	return ecdh.NewECDHDecrypt(kh)
}

// Decrypt a deserialized JWE, decrypts its protected content and returns plaintext.
func (jd *JWEDecrypt) Decrypt(jwe *JSONWebEncryption) ([]byte, error) {
	err := jd.validateAndExtractProtectedHeaders(jwe)
	if err != nil {
		return nil, fmt.Errorf("jwedecrypt: %w", err)
	}

	var senderOpt cryptoapi.WrapKeyOpts

	skid, ok := jwe.ProtectedHeaders.SenderKeyID()
	if ok && skid != "" {
		senderKH, e := jd.fetchSenderPubKey(skid)
		if e != nil {
			return nil, fmt.Errorf("jwedecrypt: failed to add sender public key for skid: %w", e)
		}

		senderOpt = cryptoapi.WithSender(senderKH)
	}

	recWK, err := buildRecipientsWrappedKey(jwe)
	if err != nil {
		return nil, fmt.Errorf("jwedecrypt: failed to build recipients WK: %w", err)
	}

	cek, err := jd.unwrapCEK(recWK, senderOpt)
	if err != nil {
		return nil, fmt.Errorf("jwedecrypt: %w", err)
	}

	return jd.decryptJWE(jwe, cek)
}

func (jd *JWEDecrypt) unwrapCEK(recWK []*cryptoapi.RecipientWrappedKey,
	senderOpt cryptoapi.WrapKeyOpts) ([]byte, error) {
	var cek []byte

	for _, rec := range recWK {
		var unwrapOpts []cryptoapi.WrapKeyOpts

		recKH, err := jd.kms.Get(rec.KID)
		if err != nil {
			continue
		}

		if rec.EPK.Type == ecdhpb.KeyType_OKP.String() {
			unwrapOpts = append(unwrapOpts, cryptoapi.WithXC20PKW())
		}

		if senderOpt != nil {
			unwrapOpts = append(unwrapOpts, senderOpt)
		}

		if len(unwrapOpts) > 0 {
			cek, err = jd.crypto.UnwrapKey(rec, recKH, unwrapOpts...)
		} else {
			cek, err = jd.crypto.UnwrapKey(rec, recKH)
		}

		if err == nil {
			break
		}
	}

	if len(cek) == 0 {
		return nil, errors.New("failed to unwrap cek")
	}

	return cek, nil
}

func (jd *JWEDecrypt) decryptJWE(jwe *JSONWebEncryption, cek []byte) ([]byte, error) {
	encAlg, ok := jwe.ProtectedHeaders.Encryption()
	if !ok {
		return nil, fmt.Errorf("jwedecrypt: JWE 'enc' protected header is missing")
	}

	decPrimitive, err := getECDHDecPrimitive(cek, encAlg)
	if err != nil {
		return nil, fmt.Errorf("jwedecrypt: failed to get decryption primitive: %w", err)
	}

	encryptedData, err := buildEncryptedData(jwe)
	if err != nil {
		return nil, fmt.Errorf("jwedecrypt: failed to build encryptedData for Decrypt(): %w", err)
	}

	authData, err := computeAuthData(jwe.ProtectedHeaders, []byte(jwe.AAD))
	if err != nil {
		return nil, err
	}

	if len(jwe.Recipients) == 1 {
		authData = []byte(jwe.OrigProtectedHders)
	}

	return decPrimitive.Decrypt(encryptedData, authData)
}

func (jd *JWEDecrypt) fetchSenderPubKey(skid string) (*keyset.Handle, error) {
	mKey, err := jd.store.Get(skid)
	if err != nil {
		return nil, fmt.Errorf("failed to get sender key from DB: %w", err)
	}

	var senderKey *cryptoapi.PublicKey

	err = json.Unmarshal(mKey, &senderKey)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal sender key from DB: %w", err)
	}

	return keyio.PublicKeyToKeysetHandle(senderKey)
}

func (jd *JWEDecrypt) validateAndExtractProtectedHeaders(jwe *JSONWebEncryption) error {
	if jwe == nil {
		return fmt.Errorf("jwe is nil")
	}

	if len(jwe.ProtectedHeaders) == 0 {
		return fmt.Errorf("jwe is missing protected headers")
	}

	protectedHeaders := jwe.ProtectedHeaders

	encAlg, ok := protectedHeaders.Encryption()
	if !ok {
		return fmt.Errorf("jwe is missing encryption algorithm 'enc' header")
	}

	switch encAlg {
	case string(A256GCM), string(XC20P):
	default:
		return fmt.Errorf("encryption algorithm '%s' not supported", encAlg)
	}

	return nil
}

func buildRecipientsWrappedKey(jwe *JSONWebEncryption) ([]*cryptoapi.RecipientWrappedKey, error) {
	var recipients []*cryptoapi.RecipientWrappedKey

	if len(jwe.Recipients) == 1 { // compact serialization: it has only 1 recipient with no headers
		rHeaders, err := extractRecipientHeaders(jwe.ProtectedHeaders)
		if err != nil {
			return nil, err
		}

		rec, err := convertMarshalledJWKToRecKey(rHeaders.EPK)
		if err != nil {
			return nil, err
		}

		rec.KID = rHeaders.KID
		rec.Alg = rHeaders.Alg
		rec.EncryptedCEK = []byte(jwe.Recipients[0].EncryptedKey)

		recipients = []*cryptoapi.RecipientWrappedKey{
			rec,
		}
	} else { // full serialization
		for _, recJWE := range jwe.Recipients {
			rec, err := convertMarshalledJWKToRecKey(recJWE.Header.EPK)
			if err != nil {
				return nil, err
			}

			rec.KID = recJWE.Header.KID
			rec.Alg = recJWE.Header.Alg
			rec.EncryptedCEK = []byte(recJWE.EncryptedKey)

			recipients = append(recipients, rec)
		}
	}

	return recipients, nil
}

func buildEncryptedData(jwe *JSONWebEncryption) ([]byte, error) {
	encData := new(composite.EncryptedData)
	encData.Tag = []byte(jwe.Tag)
	encData.IV = []byte(jwe.IV)
	encData.Ciphertext = []byte(jwe.Ciphertext)

	return json.Marshal(encData)
}

// extractRecipientHeaders will extract RecipientHeaders from headers argument.
func extractRecipientHeaders(headers map[string]interface{}) (*RecipientHeaders, error) {
	// Since headers is a generic map, epk value is converted to a generic map by Serialize(), ie we lose RawMessage
	// type of epk. We need to convert epk value (generic map) to marshaled json so we can call RawMessage.Unmarshal()
	// to get the original epk value (RawMessage type).
	mapData, ok := headers[HeaderEPK].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("JSON value is not a map (%#v)", headers[HeaderEPK])
	}

	epkBytes, err := json.Marshal(mapData)
	if err != nil {
		return nil, err
	}

	epk := json.RawMessage{}

	err = epk.UnmarshalJSON(epkBytes)
	if err != nil {
		return nil, err
	}

	alg := ""
	if headers[HeaderAlgorithm] != nil {
		alg = fmt.Sprintf("%v", headers[HeaderAlgorithm])
	}

	kid := ""
	if headers[HeaderKeyID] != nil {
		kid = fmt.Sprintf("%v", headers[HeaderKeyID])
	}

	recHeaders := &RecipientHeaders{
		Alg: alg,
		KID: kid,
		EPK: epk,
	}

	// now delete from headers
	delete(headers, HeaderAlgorithm)
	delete(headers, HeaderKeyID)
	delete(headers, HeaderEPK)

	return recHeaders, nil
}

func convertMarshalledJWKToRecKey(marshalledJWK []byte) (*cryptoapi.RecipientWrappedKey, error) {
	jwk := &JWK{}

	err := jwk.UnmarshalJSON(marshalledJWK)
	if err != nil {
		return nil, err
	}

	epk := cryptoapi.PublicKey{
		Curve: jwk.Crv,
		Type:  jwk.Kty,
	}

	switch key := jwk.Key.(type) {
	case *ecdsa.PublicKey:
		epk.X = key.X.Bytes()
		epk.Y = key.Y.Bytes()
	case []byte:
		epk.X = key
	default:
		return nil, fmt.Errorf("unsupported recipient key type")
	}

	return &cryptoapi.RecipientWrappedKey{
		KID: jwk.KeyID,
		EPK: epk,
	}, nil
}
