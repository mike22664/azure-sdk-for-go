//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"net/http"
	"regexp"
)

// ProximityPlacementGroupsServer is a fake server for instances of the armcompute.ProximityPlacementGroupsClient type.
type ProximityPlacementGroupsServer struct {
	CreateOrUpdate              func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters armcompute.ProximityPlacementGroup, options *armcompute.ProximityPlacementGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientCreateOrUpdateResponse], err azfake.ErrorResponder)
	Delete                      func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *armcompute.ProximityPlacementGroupsClientDeleteOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientDeleteResponse], err azfake.ErrorResponder)
	Get                         func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *armcompute.ProximityPlacementGroupsClientGetOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientGetResponse], err azfake.ErrorResponder)
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.ProximityPlacementGroupsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse])
	NewListBySubscriptionPager  func(options *armcompute.ProximityPlacementGroupsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse])
	Update                      func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters armcompute.ProximityPlacementGroupUpdate, options *armcompute.ProximityPlacementGroupsClientUpdateOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientUpdateResponse], err azfake.ErrorResponder)
}

// NewProximityPlacementGroupsServerTransport creates a new instance of ProximityPlacementGroupsServerTransport with the provided implementation.
// The returned ProximityPlacementGroupsServerTransport instance is connected to an instance of armcompute.ProximityPlacementGroupsClient by way of the
// arm.ClientOptions.Transporter field.
func NewProximityPlacementGroupsServerTransport(srv *ProximityPlacementGroupsServer) *ProximityPlacementGroupsServerTransport {
	return &ProximityPlacementGroupsServerTransport{srv: srv}
}

// ProximityPlacementGroupsServerTransport connects instances of armcompute.ProximityPlacementGroupsClient to instances of ProximityPlacementGroupsServer.
// Don't use this type directly, use NewProximityPlacementGroupsServerTransport instead.
type ProximityPlacementGroupsServerTransport struct {
	srv                         *ProximityPlacementGroupsServer
	newListByResourceGroupPager *azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse]
	newListBySubscriptionPager  *azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse]
}

// Do implements the policy.Transporter interface for ProximityPlacementGroupsServerTransport.
func (p *ProximityPlacementGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProximityPlacementGroupsClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "ProximityPlacementGroupsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "ProximityPlacementGroupsClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProximityPlacementGroupsClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "ProximityPlacementGroupsClient.NewListBySubscriptionPager":
		resp, err = p.dispatchNewListBySubscriptionPager(req)
	case "ProximityPlacementGroupsClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method CreateOrUpdate not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := azfake.UnmarshalRequestAsJSON[armcompute.ProximityPlacementGroup](req)
	if err != nil {
		return nil, err
	}
	resp, errResp := p.srv.CreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("proximityPlacementGroupName")], body, nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (p *ProximityPlacementGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("method Delete not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resp, errResp := p.srv.Delete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("proximityPlacementGroupName")], nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (p *ProximityPlacementGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	options := armcompute.ProximityPlacementGroupsClientGetOptions{
		IncludeColocationStatus: getOptional(qp.Get("includeColocationStatus")),
	}
	resp, errResp := p.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("proximityPlacementGroupName")], &options)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (p *ProximityPlacementGroupsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if p.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/proximityPlacementGroups"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		resp.InjectNextLinks(req, func(page *armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		p.newListByResourceGroupPager = &resp
	}

	resp, err := p.newListByResourceGroupPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !p.newListByResourceGroupPager.More() {
		p.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListBySubscriptionPager not implemented")}
	}
	if p.newListBySubscriptionPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/proximityPlacementGroups"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListBySubscriptionPager(nil)
		resp.InjectNextLinks(req, func(page *armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		p.newListBySubscriptionPager = &resp
	}

	resp, err := p.newListBySubscriptionPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !p.newListBySubscriptionPager.More() {
		p.newListBySubscriptionPager = nil
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("method Update not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := azfake.UnmarshalRequestAsJSON[armcompute.ProximityPlacementGroupUpdate](req)
	if err != nil {
		return nil, err
	}
	resp, errResp := p.srv.Update(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("proximityPlacementGroupName")], body, nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}
