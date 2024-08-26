//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkflowVersionsServer is a fake server for instances of the armappservice.WorkflowVersionsClient type.
type WorkflowVersionsServer struct {
	// Get is the fake for method WorkflowVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, workflowName string, versionID string, options *armappservice.WorkflowVersionsClientGetOptions) (resp azfake.Responder[armappservice.WorkflowVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, name string, workflowName string, options *armappservice.WorkflowVersionsClientListOptions) (resp azfake.PagerResponder[armappservice.WorkflowVersionsClientListResponse])
}

// NewWorkflowVersionsServerTransport creates a new instance of WorkflowVersionsServerTransport with the provided implementation.
// The returned WorkflowVersionsServerTransport instance is connected to an instance of armappservice.WorkflowVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowVersionsServerTransport(srv *WorkflowVersionsServer) *WorkflowVersionsServerTransport {
	return &WorkflowVersionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappservice.WorkflowVersionsClientListResponse]](),
	}
}

// WorkflowVersionsServerTransport connects instances of armappservice.WorkflowVersionsClient to instances of WorkflowVersionsServer.
// Don't use this type directly, use NewWorkflowVersionsServerTransport instead.
type WorkflowVersionsServerTransport struct {
	srv          *WorkflowVersionsServer
	newListPager *tracker[azfake.PagerResponder[armappservice.WorkflowVersionsClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkflowVersionsServerTransport.
func (w *WorkflowVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkflowVersionsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkflowVersionsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkflowVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	versionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, versionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armappservice.WorkflowVersionsClientListOptions
		if topParam != nil {
			options = &armappservice.WorkflowVersionsClientListOptions{
				Top: topParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, nameParam, workflowNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.WorkflowVersionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}
