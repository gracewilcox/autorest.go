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

// ValueTypesDatetimeServer is a fake server for instances of the valuetypesgroup.ValueTypesDatetimeClient type.
type ValueTypesDatetimeServer struct {
	// Get is the fake for method ValueTypesDatetimeClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *valuetypesgroup.ValueTypesDatetimeClientGetOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDatetimeClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ValueTypesDatetimeClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body valuetypesgroup.DatetimeProperty, options *valuetypesgroup.ValueTypesDatetimeClientPutOptions) (resp azfake.Responder[valuetypesgroup.ValueTypesDatetimeClientPutResponse], errResp azfake.ErrorResponder)
}

// NewValueTypesDatetimeServerTransport creates a new instance of ValueTypesDatetimeServerTransport with the provided implementation.
// The returned ValueTypesDatetimeServerTransport instance is connected to an instance of valuetypesgroup.ValueTypesDatetimeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewValueTypesDatetimeServerTransport(srv *ValueTypesDatetimeServer) *ValueTypesDatetimeServerTransport {
	return &ValueTypesDatetimeServerTransport{srv: srv}
}

// ValueTypesDatetimeServerTransport connects instances of valuetypesgroup.ValueTypesDatetimeClient to instances of ValueTypesDatetimeServer.
// Don't use this type directly, use NewValueTypesDatetimeServerTransport instead.
type ValueTypesDatetimeServerTransport struct {
	srv *ValueTypesDatetimeServer
}

// Do implements the policy.Transporter interface for ValueTypesDatetimeServerTransport.
func (v *ValueTypesDatetimeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *ValueTypesDatetimeServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if valueTypesDatetimeServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = valueTypesDatetimeServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ValueTypesDatetimeClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "ValueTypesDatetimeClient.Put":
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

func (v *ValueTypesDatetimeServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatetimeProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *ValueTypesDatetimeServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if v.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[valuetypesgroup.DatetimeProperty](req)
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

// set this to conditionally intercept incoming requests to ValueTypesDatetimeServerTransport
var valueTypesDatetimeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
