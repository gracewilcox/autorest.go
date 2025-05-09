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

// ArrayModelValueServer is a fake server for instances of the arraygroup.ArrayModelValueClient type.
type ArrayModelValueServer struct {
	// Get is the fake for method ArrayModelValueClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, options *arraygroup.ArrayModelValueClientGetOptions) (resp azfake.Responder[arraygroup.ArrayModelValueClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ArrayModelValueClient.Put
	// HTTP status codes to indicate success: http.StatusNoContent
	Put func(ctx context.Context, body []arraygroup.InnerModel, options *arraygroup.ArrayModelValueClientPutOptions) (resp azfake.Responder[arraygroup.ArrayModelValueClientPutResponse], errResp azfake.ErrorResponder)
}

// NewArrayModelValueServerTransport creates a new instance of ArrayModelValueServerTransport with the provided implementation.
// The returned ArrayModelValueServerTransport instance is connected to an instance of arraygroup.ArrayModelValueClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewArrayModelValueServerTransport(srv *ArrayModelValueServer) *ArrayModelValueServerTransport {
	return &ArrayModelValueServerTransport{srv: srv}
}

// ArrayModelValueServerTransport connects instances of arraygroup.ArrayModelValueClient to instances of ArrayModelValueServer.
// Don't use this type directly, use NewArrayModelValueServerTransport instead.
type ArrayModelValueServerTransport struct {
	srv *ArrayModelValueServer
}

// Do implements the policy.Transporter interface for ArrayModelValueServerTransport.
func (a *ArrayModelValueServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *ArrayModelValueServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if arrayModelValueServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = arrayModelValueServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ArrayModelValueClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "ArrayModelValueClient.Put":
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

func (a *ArrayModelValueServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InnerModelArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ArrayModelValueServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if a.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[[]arraygroup.InnerModel](req)
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

// set this to conditionally intercept incoming requests to ArrayModelValueServerTransport
var arrayModelValueServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
