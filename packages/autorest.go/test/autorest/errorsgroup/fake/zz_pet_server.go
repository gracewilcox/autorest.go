// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/errorsgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
)

// PetServer is a fake server for instances of the errorsgroup.PetClient type.
type PetServer struct {
	// DoSomething is the fake for method PetClient.DoSomething
	// HTTP status codes to indicate success: http.StatusOK
	DoSomething func(ctx context.Context, whatAction string, options *errorsgroup.PetClientDoSomethingOptions) (resp azfake.Responder[errorsgroup.PetClientDoSomethingResponse], errResp azfake.ErrorResponder)

	// GetPetByID is the fake for method PetClient.GetPetByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	GetPetByID func(ctx context.Context, petID string, options *errorsgroup.PetClientGetPetByIDOptions) (resp azfake.Responder[errorsgroup.PetClientGetPetByIDResponse], errResp azfake.ErrorResponder)

	// HasModelsParam is the fake for method PetClient.HasModelsParam
	// HTTP status codes to indicate success: http.StatusOK
	HasModelsParam func(ctx context.Context, options *errorsgroup.PetClientHasModelsParamOptions) (resp azfake.Responder[errorsgroup.PetClientHasModelsParamResponse], errResp azfake.ErrorResponder)
}

// NewPetServerTransport creates a new instance of PetServerTransport with the provided implementation.
// The returned PetServerTransport instance is connected to an instance of errorsgroup.PetClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPetServerTransport(srv *PetServer) *PetServerTransport {
	return &PetServerTransport{srv: srv}
}

// PetServerTransport connects instances of errorsgroup.PetClient to instances of PetServer.
// Don't use this type directly, use NewPetServerTransport instead.
type PetServerTransport struct {
	srv *PetServer
}

// Do implements the policy.Transporter interface for PetServerTransport.
func (p *PetServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PetServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if petServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = petServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PetClient.DoSomething":
				res.resp, res.err = p.dispatchDoSomething(req)
			case "PetClient.GetPetByID":
				res.resp, res.err = p.dispatchGetPetByID(req)
			case "PetClient.HasModelsParam":
				res.resp, res.err = p.dispatchHasModelsParam(req)
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

func (p *PetServerTransport) dispatchDoSomething(req *http.Request) (*http.Response, error) {
	if p.srv.DoSomething == nil {
		return nil, &nonRetriableError{errors.New("fake for method DoSomething not implemented")}
	}
	const regexStr = `/errorStatusCodes/Pets/doSomething/(?P<whatAction>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	whatActionParam, err := url.PathUnescape(matches[regex.SubexpIndex("whatAction")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DoSomething(req.Context(), whatActionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PetAction, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PetServerTransport) dispatchGetPetByID(req *http.Request) (*http.Response, error) {
	if p.srv.GetPetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPetByID not implemented")}
	}
	const regexStr = `/errorStatusCodes/Pets/(?P<petId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/GetPet`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	petIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("petId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetPetByID(req.Context(), petIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Pet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PetServerTransport) dispatchHasModelsParam(req *http.Request) (*http.Response, error) {
	if p.srv.HasModelsParam == nil {
		return nil, &nonRetriableError{errors.New("fake for method HasModelsParam not implemented")}
	}
	qp := req.URL.Query()
	modelsUnescaped, err := url.QueryUnescape(qp.Get("models"))
	if err != nil {
		return nil, err
	}
	modelsParam := getOptional(modelsUnescaped)
	var options *errorsgroup.PetClientHasModelsParamOptions
	if modelsParam != nil {
		options = &errorsgroup.PetClientHasModelsParamOptions{
			Models: modelsParam,
		}
	}
	respr, errRespr := p.srv.HasModelsParam(req.Context(), options)
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

// set this to conditionally intercept incoming requests to PetServerTransport
var petServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
