/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package issuecredential

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	client "github.com/Universal-Health-Chain/aries-framework-go/pkg/client/issuecredential"
	"github.com/Universal-Health-Chain/aries-framework-go/pkg/controller/rest"
	mocks "github.com/Universal-Health-Chain/aries-framework-go/pkg/internal/gomocks/client/issuecredential"
	mocknotifier "github.com/Universal-Health-Chain/aries-framework-go/pkg/internal/gomocks/controller/webnotifier"
)

func provider(ctrl *gomock.Controller) client.Provider {
	service := mocks.NewMockProtocolService(ctrl)
	service.EXPECT().RegisterActionEvent(gomock.Any()).Return(nil)
	service.EXPECT().RegisterMsgEvent(gomock.Any()).Return(nil)
	service.EXPECT().ActionContinue(gomock.Any(), gomock.Any()).AnyTimes()
	service.EXPECT().ActionStop(gomock.Any(), gomock.Any()).AnyTimes()

	provider := mocks.NewMockProvider(ctrl)
	provider.EXPECT().Service(gomock.Any()).Return(service, nil)

	return provider
}

func TestOperation_AcceptProposal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("No payload", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		buf, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptProposal), nil,
			strings.Replace(AcceptProposal, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, code)
		require.Contains(t, buf.String(), "payload was not provided")
	})

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptProposal),
			bytes.NewBufferString(`{"offer_credential":{}}`),
			strings.Replace(AcceptProposal, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_AcceptOffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptOffer),
			nil,
			strings.Replace(AcceptOffer, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_AcceptProblemReport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptProblemReport),
			nil,
			strings.Replace(AcceptProblemReport, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_AcceptRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("No payload", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		buf, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptRequest), nil,
			strings.Replace(AcceptRequest, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, code)
		require.Contains(t, buf.String(), "payload was not provided")
	})

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptRequest),
			bytes.NewBufferString(`{"issue_credential":{}}`),
			strings.Replace(AcceptRequest, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_NegotiateProposal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("No payload", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		buf, code, err := sendRequestToHandler(
			handlerLookup(t, operation, NegotiateProposal), nil,
			strings.Replace(NegotiateProposal, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, code)
		require.Contains(t, buf.String(), "payload was not provided")
	})

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, NegotiateProposal),
			bytes.NewBufferString(`{"propose_credential":{}}`),
			strings.Replace(NegotiateProposal, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_AcceptCredential(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("No payload", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		buf, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptCredential), nil,
			strings.Replace(AcceptCredential, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusBadRequest, code)
		require.Contains(t, buf.String(), "payload was not provided")
	})

	t.Run("Empty payload (success)", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptCredential),
			bytes.NewBufferString(`{}`),
			strings.Replace(AcceptCredential, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, AcceptCredential),
			bytes.NewBufferString(`{"names":[]}`),
			strings.Replace(AcceptCredential, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_DeclineProposal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, DeclineProposal),
			nil,
			strings.Replace(DeclineProposal, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_DeclineOffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, DeclineOffer),
			nil,
			strings.Replace(DeclineOffer, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_DeclineRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, DeclineRequest),
			nil,
			strings.Replace(DeclineRequest, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func TestOperation_DeclineCredential(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Success", func(t *testing.T) {
		operation, err := New(provider(ctrl), mocknotifier.NewMockNotifier(nil))
		require.NoError(t, err)

		_, code, err := sendRequestToHandler(
			handlerLookup(t, operation, DeclineCredential),
			nil,
			strings.Replace(DeclineCredential, `{piid}`, "1234", 1),
		)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, code)
	})
}

func handlerLookup(t *testing.T, op *Operation, lookup string) rest.Handler {
	t.Helper()

	handlers := op.GetRESTHandlers()
	require.NotEmpty(t, handlers)

	for _, h := range handlers {
		if h.Path() == lookup {
			return h
		}
	}

	require.Fail(t, "unable to find handler")

	return nil
}

// sendRequestToHandler reads response from given http handle func.
func sendRequestToHandler(handler rest.Handler, requestBody io.Reader, path string) (*bytes.Buffer, int, error) {
	// prepare request
	req, err := http.NewRequest(handler.Method(), path, requestBody)
	if err != nil {
		return nil, 0, err
	}

	// prepare router
	router := mux.NewRouter()

	router.HandleFunc(handler.Path(), handler.Handle()).Methods(handler.Method())

	// create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// serve http on given response and request
	router.ServeHTTP(rr, req)

	return rr.Body, rr.Code, nil
}
