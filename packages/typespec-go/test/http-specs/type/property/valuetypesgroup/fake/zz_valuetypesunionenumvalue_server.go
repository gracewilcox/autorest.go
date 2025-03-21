// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"valuetypesgroup"
)

// ValueTypesUnionEnumValueServer is a fake server for instances of the valuetypesgroup.ValueTypesUnionEnumValueClient type.
type ValueTypesUnionEnumValueServer struct {
	// Get is the fake for method ValueTypesUnionEnumValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesUnionEnumValueClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesUnionEnumValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesUnionEnumValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.UnionEnumValueProperty, options *valuetypesgroup.ValueTypesUnionEnumValueClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesUnionEnumValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesUnionEnumValueServerTransport creates a new instance of ValueTypesUnionEnumValueServerTransport with the provided implementation.
// The returned ValueTypesUnionEnumValueServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesUnionEnumValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesUnionEnumValueServerTransport(srv *ValueTypesUnionEnumValueServer) *ValueTypesUnionEnumValueServerTransport {
	return &ValueTypesUnionEnumValueServerTransport{srv: srv}
}

// ValueTypesUnionEnumValueServerTransport connects instances of valuetypesgroup.ValueTypesUnionEnumValueClient to instances of ValueTypesUnionEnumValueServer.
// Don't use this type directly, use NewValueTypesUnionEnumValueServerTransport instead.
type ValueTypesUnionEnumValueServerTransport struct {
	srv *ValueTypesUnionEnumValueServer
}

// Do implements the policy.Transporter interface for ValueTypesUnionEnumValueServerTransport.
func (v *ValueTypesUnionEnumValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesUnionEnumValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if valueTypesUnionEnumValueServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = valueTypesUnionEnumValueServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ValueTypesUnionEnumValueClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "ValueTypesUnionEnumValueClient.Put":
				res.resp, res.err = v.dispatchPut(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (v *ValueTypesUnionEnumValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := v.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnionEnumValueProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesUnionEnumValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.UnionEnumValueProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Put(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ValueTypesUnionEnumValueServerTransport
var valueTypesUnionEnumValueServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
