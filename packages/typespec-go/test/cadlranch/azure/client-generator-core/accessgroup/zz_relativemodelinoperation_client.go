//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// RelativeModelInOperationClient contains the methods for the _Specs_.Azure.ClientGenerator.Core.Access group.
// Don't use this type directly, use a constructor function instead.
type RelativeModelInOperationClient struct {
	internal *azcore.Client
}

// Discriminator - Expected query parameter: kind=<any string>
// Expected response body:
// ```json
// {
// "name": <any string>,
// "kind": "real"
// }
// ```
func (client *RelativeModelInOperationClient) Discriminator(ctx context.Context, kind string, options *RelativeModelInOperationClientDiscriminatorOptions) (RelativeModelInOperationClientDiscriminatorResponse, error) {
	var err error
	req, err := client.discriminatorCreateRequest(ctx, kind, options)
	if err != nil {
		return RelativeModelInOperationClientDiscriminatorResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RelativeModelInOperationClientDiscriminatorResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RelativeModelInOperationClientDiscriminatorResponse{}, err
	}
	resp, err := client.discriminatorHandleResponse(httpResp)
	return resp, err
}

// discriminatorCreateRequest creates the Discriminator request.
func (client *RelativeModelInOperationClient) discriminatorCreateRequest(ctx context.Context, kind string, options *RelativeModelInOperationClientDiscriminatorOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/relativeModelInOperation/discriminator"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("kind", kind)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// discriminatorHandleResponse handles the Discriminator response.
func (client *RelativeModelInOperationClient) discriminatorHandleResponse(resp *http.Response) (RelativeModelInOperationClientDiscriminatorResponse, error) {
	result := RelativeModelInOperationClientDiscriminatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return RelativeModelInOperationClientDiscriminatorResponse{}, err
	}
	return result, nil
}

// Operation - Expected query parameter: name=<any string>
// Expected response body:
// ```json
// {
// "name": <any string>,
// "inner":
// {
// "name": <any string>
// }
// }
// ```
func (client *RelativeModelInOperationClient) Operation(ctx context.Context, name string, options *RelativeModelInOperationClientOperationOptions) (RelativeModelInOperationClientOperationResponse, error) {
	var err error
	req, err := client.operationCreateRequest(ctx, name, options)
	if err != nil {
		return RelativeModelInOperationClientOperationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RelativeModelInOperationClientOperationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RelativeModelInOperationClientOperationResponse{}, err
	}
	resp, err := client.operationHandleResponse(httpResp)
	return resp, err
}

// operationCreateRequest creates the Operation request.
func (client *RelativeModelInOperationClient) operationCreateRequest(ctx context.Context, name string, options *RelativeModelInOperationClientOperationOptions) (*policy.Request, error) {
	urlPath := "/azure/client-generator-core/access/relativeModelInOperation/operation"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("name", name)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// operationHandleResponse handles the Operation response.
func (client *RelativeModelInOperationClient) operationHandleResponse(resp *http.Response) (RelativeModelInOperationClientOperationResponse, error) {
	result := RelativeModelInOperationClientOperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OuterModel); err != nil {
		return RelativeModelInOperationClientOperationResponse{}, err
	}
	return result, nil
}
