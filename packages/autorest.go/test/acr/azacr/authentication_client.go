//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azacr

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func NewAuthenticationClient(endpoint string, options *azcore.ClientOptions) (*AuthenticationClient, error) {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client, err := azcore.NewClient("azacr.AuthenticationClient", "v0.1.0", runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &AuthenticationClient{
		internal: client,
		endpoint: endpoint,
	}, nil
}
