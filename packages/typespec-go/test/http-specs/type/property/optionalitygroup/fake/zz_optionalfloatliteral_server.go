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
	"optionalitygroup"
)

// OptionalFloatLiteralServer is a fake server for instances of the optionalitygroup.OptionalFloatLiteralClient type.
type OptionalFloatLiteralServer struct {
	// GetAll is the fake for method OptionalFloatLiteralClient.GetAll
	// HTTP status codes to indicate success: http.StatusOK
	GetAll func(ctx context.Context, options *optionalitygroup.OptionalFloatLiteralClientGetAllOptions) (resp azfake.Responder[optionalitygroup.OptionalFloatLiteralClientGetAllResponse], errResp azfake.ErrorResponder)

	// GetDefault is the fake for method OptionalFloatLiteralClient.GetDefault
	// HTTP status codes to indicate success: http.StatusOK
	GetDefault func(ctx context.Context, options *optionalitygroup.OptionalFloatLiteralClientGetDefaultOptions) (resp azfake.Responder[optionalitygroup.OptionalFloatLiteralClientGetDefaultResponse], errResp azfake.ErrorResponder)

	// PutAll is the fake for method OptionalFloatLiteralClient.PutAll
	// HTTP status codes to indicate success: http.StatusNoContent
	PutAll func(ctx context.Context, body optionalitygroup.FloatLiteralProperty, options *optionalitygroup.OptionalFloatLiteralClientPutAllOptions) (resp azfake.Responder[optionalitygroup.OptionalFloatLiteralClientPutAllResponse], errResp azfake.ErrorResponder)

	// PutDefault is the fake for method OptionalFloatLiteralClient.PutDefault
	// HTTP status codes to indicate success: http.StatusNoContent
	PutDefault func(ctx context.Context, body optionalitygroup.FloatLiteralProperty, options *optionalitygroup.OptionalFloatLiteralClientPutDefaultOptions) (resp azfake.Responder[optionalitygroup.OptionalFloatLiteralClientPutDefaultResponse], errResp azfake.ErrorResponder)
}

// NewOptionalFloatLiteralServerTransport creates a new instance of OptionalFloatLiteralServerTransport with the provided implementation.
// The returned OptionalFloatLiteralServerTransport instance is connected to an instance of optionalitygroup.OptionalFloatLiteralClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOptionalFloatLiteralServerTransport(srv *OptionalFloatLiteralServer) *OptionalFloatLiteralServerTransport {
	return &OptionalFloatLiteralServerTransport{srv: srv}
}

// OptionalFloatLiteralServerTransport connects instances of optionalitygroup.OptionalFloatLiteralClient to instances of OptionalFloatLiteralServer.
// Don't use this type directly, use NewOptionalFloatLiteralServerTransport instead.
type OptionalFloatLiteralServerTransport struct {
	srv *OptionalFloatLiteralServer
}

// Do implements the policy.Transporter interface for OptionalFloatLiteralServerTransport.
func (o *OptionalFloatLiteralServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OptionalFloatLiteralServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if optionalFloatLiteralServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = optionalFloatLiteralServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OptionalFloatLiteralClient.GetAll":
				res.resp, res.err = o.dispatchGetAll(req)
			case "OptionalFloatLiteralClient.GetDefault":
				res.resp, res.err = o.dispatchGetDefault(req)
			case "OptionalFloatLiteralClient.PutAll":
				res.resp, res.err = o.dispatchPutAll(req)
			case "OptionalFloatLiteralClient.PutDefault":
				res.resp, res.err = o.dispatchPutDefault(req)
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

func (o *OptionalFloatLiteralServerTransport) dispatchGetAll(req *http.Request) (*http.Response, error) {
	if o.srv.GetAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAll not implemented")}
	}
	respr, errRespr := o.srv.GetAll(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FloatLiteralProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalFloatLiteralServerTransport) dispatchGetDefault(req *http.Request) (*http.Response, error) {
	if o.srv.GetDefault == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDefault not implemented")}
	}
	respr, errRespr := o.srv.GetDefault(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FloatLiteralProperty, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OptionalFloatLiteralServerTransport) dispatchPutAll(req *http.Request) (*http.Response, error) {
	if o.srv.PutAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutAll not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalitygroup.FloatLiteralProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.PutAll(req.Context(), body, nil)
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

func (o *OptionalFloatLiteralServerTransport) dispatchPutDefault(req *http.Request) (*http.Response, error) {
	if o.srv.PutDefault == nil {
		return nil, &nonRetriableError{errors.New("fake for method PutDefault not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[optionalitygroup.FloatLiteralProperty](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.PutDefault(req.Context(), body, nil)
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

// set this to conditionally intercept incoming requests to OptionalFloatLiteralServerTransport
var optionalFloatLiteralServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
