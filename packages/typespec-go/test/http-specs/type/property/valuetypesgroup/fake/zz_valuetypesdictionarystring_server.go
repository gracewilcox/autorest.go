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

// ValueTypesDictionaryStringServer is a fake server for instances of the valuetypesgroup.ValueTypesDictionaryStringClient type.
type ValueTypesDictionaryStringServer struct {
	// Get is the fake for method ValueTypesDictionaryStringClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesDictionaryStringClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDictionaryStringClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesDictionaryStringClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.DictionaryStringProperty, options *valuetypesgroup.ValueTypesDictionaryStringClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDictionaryStringClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesDictionaryStringServerTransport creates a new instance of ValueTypesDictionaryStringServerTransport with the provided implementation.
// The returned ValueTypesDictionaryStringServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesDictionaryStringClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesDictionaryStringServerTransport(srv *ValueTypesDictionaryStringServer) *ValueTypesDictionaryStringServerTransport {
	return &ValueTypesDictionaryStringServerTransport{srv: srv}
}

// ValueTypesDictionaryStringServerTransport connects instances of valuetypesgroup.ValueTypesDictionaryStringClient to instances of ValueTypesDictionaryStringServer.
// Don't use this type directly, use NewValueTypesDictionaryStringServerTransport instead.
type ValueTypesDictionaryStringServerTransport struct {
	srv *ValueTypesDictionaryStringServer
}

// Do implements the policy.Transporter interface for ValueTypesDictionaryStringServerTransport.
func (v *ValueTypesDictionaryStringServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesDictionaryStringServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if valueTypesDictionaryStringServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = valueTypesDictionaryStringServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ValueTypesDictionaryStringClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "ValueTypesDictionaryStringClient.Put":
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

func (v *ValueTypesDictionaryStringServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DictionaryStringProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesDictionaryStringServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.DictionaryStringProperty](req)
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

// set this to conditionally intercept incoming requests to ValueTypesDictionaryStringServerTransport
var valueTypesDictionaryStringServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
