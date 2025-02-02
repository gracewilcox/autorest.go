//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package coreusagegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// UsageClient contains the methods for the _Specs_.Azure.ClientGenerator.Core.Usage group.
// Don't use this type directly, use a constructor function instead.
type UsageClient struct {
	internal *azcore.Client
}

// InputToInputOutput - Expected body parameter:
// ```json
// {
// "name": <any string>
// }
// ```
func (client *UsageClient) InputToInputOutput(ctx context.Context, body InputModel, options *UsageClientInputToInputOutputOptions) (UsageClientInputToInputOutputResponse, error) {
	var err error
	req, err := client.inputToInputOutputCreateRequest(ctx, body, options)
	if err != nil {
		return UsageClientInputToInputOutputResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UsageClientInputToInputOutputResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UsageClientInputToInputOutputResponse{}, err
	}
	return UsageClientInputToInputOutputResponse{}, nil
}

// inputToInputOutputCreateRequest creates the InputToInputOutput request.
func (client *UsageClient) inputToInputOutputCreateRequest(ctx context.Context, body InputModel, options *UsageClientInputToInputOutputOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/usage/inputToInputOutput"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// OutputToInputOutput - Expected response body:
// ```json
// {
// "name": <any string>
// }
// ```
func (client *UsageClient) OutputToInputOutput(ctx context.Context, options *UsageClientOutputToInputOutputOptions) (UsageClientOutputToInputOutputResponse, error) {
	var err error
	req, err := client.outputToInputOutputCreateRequest(ctx, options)
	if err != nil {
		return UsageClientOutputToInputOutputResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UsageClientOutputToInputOutputResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UsageClientOutputToInputOutputResponse{}, err
	}
	resp, err := client.outputToInputOutputHandleResponse(httpResp)
	return resp, err
}

// outputToInputOutputCreateRequest creates the OutputToInputOutput request.
func (client *UsageClient) outputToInputOutputCreateRequest(ctx context.Context, options *UsageClientOutputToInputOutputOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/usage/outputToInputOutput"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// outputToInputOutputHandleResponse handles the OutputToInputOutput response.
func (client *UsageClient) outputToInputOutputHandleResponse(resp *http.Response) (UsageClientOutputToInputOutputResponse, error) {
	result := UsageClientOutputToInputOutputResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutputModel); err != nil {
		return UsageClientOutputToInputOutputResponse{}, err
	}
	return result, nil
}
