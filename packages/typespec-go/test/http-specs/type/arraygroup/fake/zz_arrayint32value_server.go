// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"arraygroup"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ArrayInt32ValueServer is a fake server for instances of the arraygroup.ArrayInt32ValueClient type.
type ArrayInt32ValueServer struct {
	// Get is the fake for method ArrayInt32ValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *arraygroup.ArrayInt32ValueClientGetOptions) (resp azfake.Responder[arraygroup.ArrayInt32ValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ArrayInt32ValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body []int32, options *arraygroup.ArrayInt32ValueClientPutOptions) (resp azfake.Responder[arraygroup.ArrayInt32ValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewArrayInt32ValueServerTransport creates a new instance of ArrayInt32ValueServerTransport with the provided implementation.
// The returned ArrayInt32ValueServerTransport instance is connected to an instance of arraygroup.ArrayInt32ValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArrayInt32ValueServerTransport(srv *ArrayInt32ValueServer) *ArrayInt32ValueServerTransport {
	return &ArrayInt32ValueServerTransport{srv: srv}
}

// ArrayInt32ValueServerTransport connects instances of arraygroup.ArrayInt32ValueClient to instances of ArrayInt32ValueServer.
// Don't use this type directly, use NewArrayInt32ValueServerTransport instead.
type ArrayInt32ValueServerTransport struct {
	srv *ArrayInt32ValueServer
}

// Do implements the policy.Transporter interface for ArrayInt32ValueServerTransport.
func (a *ArrayInt32ValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ArrayInt32ValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if arrayInt32ValueServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = arrayInt32ValueServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ArrayInt32ValueClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ArrayInt32ValueClient.Put":
				res.resp, res.err = a.dispatchPut(req)
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

func (a *ArrayInt32ValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	respr, errRespr := a.srv.Get(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Int32Array, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayInt32ValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if a.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]int32](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Put(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to ArrayInt32ValueServerTransport
var arrayInt32ValueServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
