// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/complexgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// BasicServer is a fake server for instances of the complexgroup.BasicClient type.
type BasicServer struct {
	// GetEmpty is the fake for method BasicClient.GetEmpty
	// HTTP status codes to indicate success: http.StatusOK
	GetEmpty func(ctx context.Context, options *complexgroup.BasicClientGetEmptyOptions) (resp azfake.Responder[complexgroup.BasicClientGetEmptyResponse], errResp azfake.ErrorResponder)

	// GetInvalid is the fake for method BasicClient.GetInvalid
	// HTTP status codes to indicate success: http.StatusOK
	GetInvalid func(ctx context.Context, options *complexgroup.BasicClientGetInvalidOptions) (resp azfake.Responder[complexgroup.BasicClientGetInvalidResponse], errResp azfake.ErrorResponder)

	// GetNotProvided is the fake for method BasicClient.GetNotProvided
	// HTTP status codes to indicate success: http.StatusOK
	GetNotProvided func(ctx context.Context, options *complexgroup.BasicClientGetNotProvidedOptions) (resp azfake.Responder[complexgroup.BasicClientGetNotProvidedResponse], errResp azfake.ErrorResponder)

	// GetNull is the fake for method BasicClient.GetNull
	// HTTP status codes to indicate success: http.StatusOK
	GetNull func(ctx context.Context, options *complexgroup.BasicClientGetNullOptions) (resp azfake.Responder[complexgroup.BasicClientGetNullResponse], errResp azfake.ErrorResponder)

	// GetValid is the fake for method BasicClient.GetValid
	// HTTP status codes to indicate success: http.StatusOK
	GetValid func(ctx context.Context, options *complexgroup.BasicClientGetValidOptions) (resp azfake.Responder[complexgroup.BasicClientGetValidResponse], errResp azfake.ErrorResponder)

	// PutValid is the fake for method BasicClient.PutValid
	// HTTP status codes to indicate success: http.StatusOK
	PutValid func(ctx context.Context, complexBody complexgroup.Basic, options *complexgroup.BasicClientPutValidOptions) (resp azfake.Responder[complexgroup.BasicClientPutValidResponse], errResp azfake.ErrorResponder)
}

// NewBasicServerTransport creates a new instance of BasicServerTransport with the provided implementation.
// The returned BasicServerTransport instance is connected to an instance of complexgroup.BasicClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBasicServerTransport(srv *BasicServer) *BasicServerTransport {
	return &BasicServerTransport{srv: srv}
}

// BasicServerTransport connects instances of complexgroup.BasicClient to instances of BasicServer.
// Don't use this type directly, use NewBasicServerTransport instead.
type BasicServerTransport struct {
	srv *BasicServer
}

// Do implements the policy.Transporter interface for BasicServerTransport.
func (b *BasicServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BasicServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if basicServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = basicServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BasicClient.GetEmpty":
				res.resp, res.err = b.dispatchGetEmpty(req)
			case "BasicClient.GetInvalid":
				res.resp, res.err = b.dispatchGetInvalid(req)
			case "BasicClient.GetNotProvided":
				res.resp, res.err = b.dispatchGetNotProvided(req)
			case "BasicClient.GetNull":
				res.resp, res.err = b.dispatchGetNull(req)
			case "BasicClient.GetValid":
				res.resp, res.err = b.dispatchGetValid(req)
			case "BasicClient.PutValid":
				res.resp, res.err = b.dispatchPutValid(req)
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

func (b *BasicServerTransport) dispatchGetEmpty(req *http.Request) (*http.Response, error) {
	if b.srv.GetEmpty == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEmpty not implemented")}
	}
	respr, errRespr := b.srv.GetEmpty(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Basic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchGetInvalid(req *http.Request) (*http.Response, error) {
	if b.srv.GetInvalid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInvalid not implemented")}
	}
	respr, errRespr := b.srv.GetInvalid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Basic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchGetNotProvided(req *http.Request) (*http.Response, error) {
	if b.srv.GetNotProvided == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNotProvided not implemented")}
	}
	respr, errRespr := b.srv.GetNotProvided(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Basic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchGetNull(req *http.Request) (*http.Response, error) {
	if b.srv.GetNull == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNull not implemented")}
	}
	respr, errRespr := b.srv.GetNull(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Basic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchGetValid(req *http.Request) (*http.Response, error) {
	if b.srv.GetValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetValid not implemented")}
	}
	respr, errRespr := b.srv.GetValid(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Basic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BasicServerTransport) dispatchPutValid(req *http.Request) (*http.Response, error) {
	if b.srv.PutValid == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutValid not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[complexgroup.Basic](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.PutValid(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BasicServerTransport
var basicServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
