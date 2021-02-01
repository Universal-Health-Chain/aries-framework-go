/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package didcomm

import "github.com/Universal-Health-Chain/aries-framework-go/pkg/didcomm/common/transport"

// MockAuthCrypt mock auth crypt.
type MockAuthCrypt struct {
	EncryptValue func(payload, senderPubKey []byte, recipients [][]byte) ([]byte, error)
	DecryptValue func(envelope []byte) (*transport.Envelope, error)
	Type         string
}

// Pack mock message packing.
func (m *MockAuthCrypt) Pack(payload, senderPubKey []byte,
	recipients [][]byte) ([]byte, error) {
	return m.EncryptValue(payload, senderPubKey, recipients)
}

// Unpack mock message unpacking.
func (m *MockAuthCrypt) Unpack(envelope []byte) (*transport.Envelope, error) {
	return m.DecryptValue(envelope)
}

// EncodingType mock encoding type.
func (m *MockAuthCrypt) EncodingType() string {
	return m.Type
}
