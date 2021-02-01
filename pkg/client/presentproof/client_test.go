/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package presentproof

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/Universal-Health-Chain/aries-framework-go/pkg/didcomm/common/service"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/didcomm/protocol/presentproof"
	mocks "github.com/Universal-Health-Chain/aries-framework-go/pkg/internal/gomocks/client/presentproof"
)

const (
	Alice = "Alice"
	Bob   = "Bob"
)

func TestNew(t *testing.T) {
	const errMsg = "test err"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("get service error", func(t *testing.T) {
		provider := mocks.NewMockProvider(ctrl)
		provider.EXPECT().Service(gomock.Any()).Return(nil, errors.New(errMsg))
		_, err := New(provider)
		require.EqualError(t, err, errMsg)
	})

	t.Run("cast service error", func(t *testing.T) {
		provider := mocks.NewMockProvider(ctrl)
		provider.EXPECT().Service(gomock.Any()).Return(nil, nil)
		_, err := New(provider)
		require.EqualError(t, err, "cast service to presentproof service failed")
	})
}

func TestClient_SendRequestPresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		provider := mocks.NewMockProvider(ctrl)
		thid := uuid.New().String()

		svc := mocks.NewMockProtocolService(ctrl)
		svc.EXPECT().HandleInbound(gomock.Any(), Alice, Bob).
			DoAndReturn(func(msg service.DIDCommMsg, _, _ string) (string, error) {
				require.Equal(t, msg.Type(), presentproof.RequestPresentationMsgType)

				return thid, nil
			})

		provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
		client, err := New(provider)
		require.NoError(t, err)

		result, err := client.SendRequestPresentation(&RequestPresentation{}, Alice, Bob)
		require.NoError(t, err)
		require.Equal(t, thid, result)
	})

	t.Run("Empty Request Presentation", func(t *testing.T) {
		provider := mocks.NewMockProvider(ctrl)

		provider.EXPECT().Service(gomock.Any()).Return(mocks.NewMockProtocolService(ctrl), nil)
		client, err := New(provider)
		require.NoError(t, err)

		_, err = client.SendRequestPresentation(nil, Alice, Bob)
		require.EqualError(t, err, errEmptyRequestPresentation.Error())
	})
}

func TestClient_SendProposePresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		provider := mocks.NewMockProvider(ctrl)
		thid := uuid.New().String()

		svc := mocks.NewMockProtocolService(ctrl)
		svc.EXPECT().HandleInbound(gomock.Any(), Alice, Bob).
			DoAndReturn(func(msg service.DIDCommMsg, _, _ string) (string, error) {
				require.Equal(t, msg.Type(), presentproof.ProposePresentationMsgType)

				return thid, nil
			})

		provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
		client, err := New(provider)
		require.NoError(t, err)

		result, err := client.SendProposePresentation(&ProposePresentation{}, Alice, Bob)
		require.NoError(t, err)
		require.Equal(t, thid, result)
	})

	t.Run("Empty Request Presentation", func(t *testing.T) {
		provider := mocks.NewMockProvider(ctrl)

		provider.EXPECT().Service(gomock.Any()).Return(mocks.NewMockProtocolService(ctrl), nil)
		client, err := New(provider)
		require.NoError(t, err)

		_, err = client.SendProposePresentation(nil, Alice, Bob)
		require.EqualError(t, err, errEmptyProposePresentation.Error())
	})
}

func TestClient_AcceptRequestPresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionContinue("PIID", gomock.Any()).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.AcceptRequestPresentation("PIID", &Presentation{}))
}

func TestClient_DeclineRequestPresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionStop("PIID", errors.New("declined")).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.DeclineRequestPresentation("PIID", "declined"))
}

func TestClient_AcceptProposePresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionContinue("PIID", gomock.Any()).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.AcceptProposePresentation("PIID", &RequestPresentation{}))
}

func TestClient_DeclineProposePresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionStop("PIID", errors.New("declined")).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.DeclineProposePresentation("PIID", "declined"))
}

func TestClient_AcceptPresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionContinue("PIID", gomock.Any()).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.AcceptPresentation("PIID"))
}

func TestClient_DeclinePresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionStop("PIID", errors.New("declined")).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.DeclinePresentation("PIID", "declined"))
}

func TestClient_AcceptProblemReport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionContinue("PIID", gomock.Any()).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.AcceptProblemReport("PIID"))
}

func TestClient_NegotiateProposePresentation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := mocks.NewMockProvider(ctrl)

	svc := mocks.NewMockProtocolService(ctrl)
	svc.EXPECT().ActionContinue("PIID", gomock.Any()).Return(nil)

	provider.EXPECT().Service(gomock.Any()).Return(svc, nil)
	client, err := New(provider)
	require.NoError(t, err)

	require.NoError(t, client.NegotiateRequestPresentation("PIID", &ProposePresentation{}))
}
