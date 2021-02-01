/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package formattedstore_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Universal-Health-Chain/aries-framework-go/pkg/storage/formattedstore"
)

func TestNewBatchWrite(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		macCrypto := newMACCrypto(t)

		s := formattedstore.NewBatchWrite(0, createEDVFormatter(t, macCrypto), newMockStoreProvider())

		err := s.Put(&mockStore{}, "k1", []byte("v1"))
		require.NoError(t, err)

		v, err := s.Get("k1")
		require.NoError(t, err)
		require.Equal(t, "v1", string(v))

		err = s.Flush()
		require.NoError(t, err)

		_, err = s.Get("k1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "k k1 not found")
	})

	t.Run("success batch size", func(t *testing.T) {
		macCrypto := newMACCrypto(t)

		s := formattedstore.NewBatchWrite(1, createEDVFormatter(t, macCrypto), newMockStoreProvider())

		err := s.Put(&mockStore{}, "k1", []byte("v1"))
		require.NoError(t, err)

		_, err = s.Get("k1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "k k1 not found")
	})

	t.Run("success delete", func(t *testing.T) {
		macCrypto := newMACCrypto(t)

		s := formattedstore.NewBatchWrite(10, createEDVFormatter(t, macCrypto), newMockStoreProvider())

		err := s.Put(&mockStore{}, "k1", []byte("v1"))
		require.NoError(t, err)

		v, err := s.Get("k1")
		require.NoError(t, err)
		require.Equal(t, "v1", string(v))

		err = s.Delete("k1")
		require.NoError(t, err)

		_, err = s.Get("k1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "value is deleted")
	})

	t.Run("error from flush", func(t *testing.T) {
		macCrypto := newMACCrypto(t)

		s := formattedstore.NewBatchWrite(10, createEDVFormatter(t, macCrypto),
			&mockStoreProvider{batchErr: fmt.Errorf("failed to put")})

		err := s.Put(&mockStore{}, "k2", []byte("v2"))
		require.NoError(t, err)

		v, err := s.Get("k2")
		require.NoError(t, err)
		require.Equal(t, "v2", string(v))

		err = s.Flush()
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to put")

		_, err = s.Get("k2")
		require.Error(t, err)
		require.Contains(t, err.Error(), "k k2 not found")
	})
}
