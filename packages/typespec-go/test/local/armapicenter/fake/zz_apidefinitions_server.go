// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"armapicenter"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"net/http"
	"net/url"
	"regexp"
)

// APIDefinitionsServer is a fake server for instances of the armapicenter.APIDefinitionsClient type.
type APIDefinitionsServer struct {
	// CreateOrUpdate is the fake for method APIDefinitionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload armapicenter.APIDefinition, options *armapicenter.APIDefinitionsClientCreateOrUpdateOptions) (resp azfake.Responder[armapicenter.APIDefinitionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method APIDefinitionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *armapicenter.APIDefinitionsClientDeleteOptions) (resp azfake.Responder[armapicenter.APIDefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExportSpecification is the fake for method APIDefinitionsClient.BeginExportSpecification
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportSpecification func(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *armapicenter.APIDefinitionsClientBeginExportSpecificationOptions) (resp azfake.PollerResponder[armapicenter.APIDefinitionsClientExportSpecificationResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method APIDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *armapicenter.APIDefinitionsClientGetOptions) (resp azfake.Responder[armapicenter.APIDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// Head is the fake for method APIDefinitionsClient.Head
	// HTTP status codes to indicate success: http.StatusOK
	Head func(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, options *armapicenter.APIDefinitionsClientHeadOptions) (resp azfake.Responder[armapicenter.APIDefinitionsClientHeadResponse], errResp azfake.ErrorResponder)

	// BeginImportSpecification is the fake for method APIDefinitionsClient.BeginImportSpecification
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginImportSpecification func(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, definitionName string, payload armapicenter.APISpecImportRequest, options *armapicenter.APIDefinitionsClientBeginImportSpecificationOptions) (resp azfake.PollerResponder[armapicenter.APIDefinitionsClientImportSpecificationResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method APIDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, workspaceName string, apiName string, versionName string, options *armapicenter.APIDefinitionsClientListOptions) (resp azfake.PagerResponder[armapicenter.APIDefinitionsClientListResponse])
}

// NewAPIDefinitionsServerTransport creates a new instance of APIDefinitionsServerTransport with the provided implementation.
// The returned APIDefinitionsServerTransport instance is connected to an instance of armapicenter.APIDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAPIDefinitionsServerTransport(srv *APIDefinitionsServer) *APIDefinitionsServerTransport {
	return &APIDefinitionsServerTransport{
		srv:                      srv,
		beginExportSpecification: newTracker[azfake.PollerResponder[armapicenter.APIDefinitionsClientExportSpecificationResponse]](),
		beginImportSpecification: newTracker[azfake.PollerResponder[armapicenter.APIDefinitionsClientImportSpecificationResponse]](),
		newListPager:             newTracker[azfake.PagerResponder[armapicenter.APIDefinitionsClientListResponse]](),
	}
}

// APIDefinitionsServerTransport connects instances of armapicenter.APIDefinitionsClient to instances of APIDefinitionsServer.
// Don't use this type directly, use NewAPIDefinitionsServerTransport instead.
type APIDefinitionsServerTransport struct {
	srv                      *APIDefinitionsServer
	beginExportSpecification *tracker[azfake.PollerResponder[armapicenter.APIDefinitionsClientExportSpecificationResponse]]
	beginImportSpecification *tracker[azfake.PollerResponder[armapicenter.APIDefinitionsClientImportSpecificationResponse]]
	newListPager             *tracker[azfake.PagerResponder[armapicenter.APIDefinitionsClientListResponse]]
}

// Do implements the policy.Transporter interface for APIDefinitionsServerTransport.
func (a *APIDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *APIDefinitionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if apiDefinitionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = apiDefinitionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "APIDefinitionsClient.CreateOrUpdate":
				res.resp, res.err = a.dispatchCreateOrUpdate(req)
			case "APIDefinitionsClient.Delete":
				res.resp, res.err = a.dispatchDelete(req)
			case "APIDefinitionsClient.BeginExportSpecification":
				res.resp, res.err = a.dispatchBeginExportSpecification(req)
			case "APIDefinitionsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "APIDefinitionsClient.Head":
				res.resp, res.err = a.dispatchHead(req)
			case "APIDefinitionsClient.BeginImportSpecification":
				res.resp, res.err = a.dispatchBeginImportSpecification(req)
			case "APIDefinitionsClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
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

func (a *APIDefinitionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions/(?P<definitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 8 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapicenter.APIDefinition](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	definitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("definitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, definitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).APIDefinition, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (a *APIDefinitionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions/(?P<definitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 8 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	definitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("definitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, definitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *APIDefinitionsServerTransport) dispatchBeginExportSpecification(req *http.Request) (*http.Response, error) {
	if a.srv.BeginExportSpecification == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportSpecification not implemented")}
	}
	beginExportSpecification := a.beginExportSpecification.get(req)
	if beginExportSpecification == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions/(?P<definitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportSpecification`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 8 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
		if err != nil {
			return nil, err
		}
		versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
		if err != nil {
			return nil, err
		}
		definitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("definitionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginExportSpecification(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, definitionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportSpecification = &respr
		a.beginExportSpecification.add(req, beginExportSpecification)
	}

	resp, err := server.PollerResponderNext(beginExportSpecification, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginExportSpecification.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportSpecification) {
		a.beginExportSpecification.remove(req)
	}

	return resp, nil
}

func (a *APIDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions/(?P<definitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 8 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	definitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("definitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, definitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).APIDefinition, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (a *APIDefinitionsServerTransport) dispatchHead(req *http.Request) (*http.Response, error) {
	if a.srv.Head == nil {
		return nil, &nonRetriableError{errors.New("fake for method Head not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions/(?P<definitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if len(matches) < 8 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	definitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("definitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Head(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, definitionNameParam, nil)
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

func (a *APIDefinitionsServerTransport) dispatchBeginImportSpecification(req *http.Request) (*http.Response, error) {
	if a.srv.BeginImportSpecification == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginImportSpecification not implemented")}
	}
	beginImportSpecification := a.beginImportSpecification.get(req)
	if beginImportSpecification == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions/(?P<definitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importSpecification`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 8 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapicenter.APISpecImportRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
		if err != nil {
			return nil, err
		}
		versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
		if err != nil {
			return nil, err
		}
		definitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("definitionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginImportSpecification(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, definitionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginImportSpecification = &respr
		a.beginImportSpecification.add(req, beginImportSpecification)
	}

	resp, err := server.PollerResponderNext(beginImportSpecification, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginImportSpecification.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginImportSpecification) {
		a.beginImportSpecification.remove(req)
	}

	return resp, nil
}

func (a *APIDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiCenter/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/definitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if len(matches) < 7 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		apiNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiName")])
		if err != nil {
			return nil, err
		}
		versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armapicenter.APIDefinitionsClientListOptions
		if filterParam != nil {
			options = &armapicenter.APIDefinitionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, serviceNameParam, workspaceNameParam, apiNameParam, versionNameParam, options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapicenter.APIDefinitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to APIDefinitionsServerTransport
var apiDefinitionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
