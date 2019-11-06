// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/aries-framework-go/cmd/aries-agentd

replace github.com/hyperledger/aries-framework-go => ../..

require (
	github.com/gorilla/mux v1.7.3
	github.com/hyperledger/aries-framework-go v0.0.0
	github.com/rs/cors v1.7.0
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.3.0
)

go 1.13
