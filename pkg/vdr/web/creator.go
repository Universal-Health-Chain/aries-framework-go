/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package web

import (
	"fmt"

	"github.com/Universal-Health-Chain/aries-framework-go/pkg/doc/did"
	vdrapi "github.com/Universal-Health-Chain/aries-framework-go/pkg/framework/aries/api/vdr"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/kms"
)

// Create creates a did:web diddoc (unsupported at the moment).
func (v *VDR) Create(keyManager kms.KeyManager, didDoc *did.Doc,
	opts ...vdrapi.DIDMethodOption) (*did.DocResolution, error) {
	return nil, fmt.Errorf("error building did:web did doc --> build not supported in http binding vdr")
}
