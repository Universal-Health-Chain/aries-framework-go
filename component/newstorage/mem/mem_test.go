/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mem_test

import (
	"testing"

	"github.com/Universal-Health-Chain/aries-framework-go/component/newstorage/mem"
	"github.com/Universal-Health-Chain/aries-framework-go/test/newstorage"
)

func TestCommon(t *testing.T) {
	provider := mem.NewProvider()

	newstorage.TestAll(t, provider)
}
