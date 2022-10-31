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

// CloudServicesServer is a fake server for instances of the armcompute.CloudServicesClient type.
type CloudServicesServer struct {
	BeginCreateOrUpdate  func(ctx context.Context, resourceGroupName string, cloudServiceName string, parameters armcompute.CloudService, options *armcompute.CloudServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse], err azfake.ErrorResponder)
	BeginDelete          func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse], err azfake.ErrorResponder)
	BeginDeleteInstances func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginDeleteInstancesOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse], err azfake.ErrorResponder)
	Get                  func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientGetOptions) (resp azfake.Responder[armcompute.CloudServicesClientGetResponse], err azfake.ErrorResponder)
	GetInstanceView      func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientGetInstanceViewOptions) (resp azfake.Responder[armcompute.CloudServicesClientGetInstanceViewResponse], err azfake.ErrorResponder)
	NewListPager         func(resourceGroupName string, options *armcompute.CloudServicesClientListOptions) (resp azfake.PagerResponder[armcompute.CloudServicesClientListResponse])
	NewListAllPager      func(options *armcompute.CloudServicesClientListAllOptions) (resp azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse])
	BeginPowerOff        func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginPowerOffOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse], err azfake.ErrorResponder)
	BeginRebuild         func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginRebuildOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse], err azfake.ErrorResponder)
	BeginReimage         func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginReimageOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse], err azfake.ErrorResponder)
	BeginRestart         func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginRestartOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse], err azfake.ErrorResponder)
	BeginStart           func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginStartOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientStartResponse], err azfake.ErrorResponder)
	BeginUpdate          func(ctx context.Context, resourceGroupName string, cloudServiceName string, parameters armcompute.CloudServiceUpdate, options *armcompute.CloudServicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse], err azfake.ErrorResponder)
}

// NewCloudServicesServerTransport creates a new instance of CloudServicesServerTransport with the provided implementation.
// The returned CloudServicesServerTransport instance is connected to an instance of armcompute.CloudServicesClient by way of the
// arm.ClientOptions.Transporter field.
func NewCloudServicesServerTransport(srv *CloudServicesServer) *CloudServicesServerTransport {
	return &CloudServicesServerTransport{srv: srv}
}

// CloudServicesServerTransport connects instances of armcompute.CloudServicesClient to instances of CloudServicesServer.
// Don't use this type directly, use NewCloudServicesServerTransport instead.
type CloudServicesServerTransport struct {
	srv                  *CloudServicesServer
	beginCreateOrUpdate  *azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse]
	beginDelete          *azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse]
	beginDeleteInstances *azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse]
	newListPager         *azfake.PagerResponder[armcompute.CloudServicesClientListResponse]
	newListAllPager      *azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse]
	beginPowerOff        *azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse]
	beginRebuild         *azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse]
	beginReimage         *azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse]
	beginRestart         *azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse]
	beginStart           *azfake.PollerResponder[armcompute.CloudServicesClientStartResponse]
	beginUpdate          *azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse]
}

// Do implements the policy.Transporter interface for CloudServicesServerTransport.
func (c *CloudServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudServicesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CloudServicesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudServicesClient.BeginDeleteInstances":
		resp, err = c.dispatchBeginDeleteInstances(req)
	case "CloudServicesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudServicesClient.GetInstanceView":
		resp, err = c.dispatchGetInstanceView(req)
	case "CloudServicesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "CloudServicesClient.NewListAllPager":
		resp, err = c.dispatchNewListAllPager(req)
	case "CloudServicesClient.BeginPowerOff":
		resp, err = c.dispatchBeginPowerOff(req)
	case "CloudServicesClient.BeginRebuild":
		resp, err = c.dispatchBeginRebuild(req)
	case "CloudServicesClient.BeginReimage":
		resp, err = c.dispatchBeginReimage(req)
	case "CloudServicesClient.BeginRestart":
		resp, err = c.dispatchBeginRestart(req)
	case "CloudServicesClient.BeginStart":
		resp, err = c.dispatchBeginStart(req)
	case "CloudServicesClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if c.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.CloudService](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := c.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginCreateOrUpdate = &resp
	}

	resp, err := c.beginCreateOrUpdate.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginCreateOrUpdate.More() {
		c.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if c.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp, errResp := c.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginDelete = &resp
	}

	resp, err := c.beginDelete.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginDelete.More() {
		c.beginDelete = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginDeleteInstances(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDeleteInstances == nil {
		return nil, &nonRetriableError{errors.New("method BeginDeleteInstances not implemented")}
	}
	if c.beginDeleteInstances == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/delete"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		options := armcompute.CloudServicesClientBeginDeleteInstancesOptions{
			Parameters: getOptional(body),
		}
		resp, errResp := c.srv.BeginDeleteInstances(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], &options)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginDeleteInstances = &resp
	}

	resp, err := c.beginDeleteInstances.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginDeleteInstances.More() {
		c.beginDeleteInstances = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resp, errResp := c.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (c *CloudServicesServerTransport) dispatchGetInstanceView(req *http.Request) (*http.Response, error) {
	if c.srv.GetInstanceView == nil {
		return nil, &nonRetriableError{errors.New("method GetInstanceView not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/instanceView"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resp, errResp := c.srv.GetInstanceView(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (c *CloudServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if c.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		resp.InjectNextLinks(req, func(page *armcompute.CloudServicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		c.newListPager = &resp
	}

	resp, err := c.newListPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !c.newListPager.More() {
		c.newListPager = nil
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListAllPager not implemented")}
	}
	if c.newListAllPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListAllPager(nil)
		resp.InjectNextLinks(req, func(page *armcompute.CloudServicesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		c.newListAllPager = &resp
	}

	resp, err := c.newListAllPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !c.newListAllPager.More() {
		c.newListAllPager = nil
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginPowerOff(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPowerOff == nil {
		return nil, &nonRetriableError{errors.New("method BeginPowerOff not implemented")}
	}
	if c.beginPowerOff == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/poweroff"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp, errResp := c.srv.BeginPowerOff(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginPowerOff = &resp
	}

	resp, err := c.beginPowerOff.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginPowerOff.More() {
		c.beginPowerOff = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginRebuild(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRebuild == nil {
		return nil, &nonRetriableError{errors.New("method BeginRebuild not implemented")}
	}
	if c.beginRebuild == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/rebuild"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		options := armcompute.CloudServicesClientBeginRebuildOptions{
			Parameters: getOptional(body),
		}
		resp, errResp := c.srv.BeginRebuild(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], &options)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginRebuild = &resp
	}

	resp, err := c.beginRebuild.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginRebuild.More() {
		c.beginRebuild = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginReimage(req *http.Request) (*http.Response, error) {
	if c.srv.BeginReimage == nil {
		return nil, &nonRetriableError{errors.New("method BeginReimage not implemented")}
	}
	if c.beginReimage == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/reimage"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		options := armcompute.CloudServicesClientBeginReimageOptions{
			Parameters: getOptional(body),
		}
		resp, errResp := c.srv.BeginReimage(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], &options)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginReimage = &resp
	}

	resp, err := c.beginReimage.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginReimage.More() {
		c.beginReimage = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("method BeginRestart not implemented")}
	}
	if c.beginRestart == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/restart"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		options := armcompute.CloudServicesClientBeginRestartOptions{
			Parameters: getOptional(body),
		}
		resp, errResp := c.srv.BeginRestart(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], &options)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginRestart = &resp
	}

	resp, err := c.beginRestart.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginRestart.More() {
		c.beginRestart = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("method BeginStart not implemented")}
	}
	if c.beginStart == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)/start"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp, errResp := c.srv.BeginStart(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginStart = &resp
	}

	resp, err := c.beginStart.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginStart.More() {
		c.beginStart = nil
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if c.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.CloudServiceUpdate](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := c.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("cloudServiceName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		c.beginUpdate = &resp
	}

	resp, err := c.beginUpdate.Next(req)
	if err != nil {
		return nil, err
	}

	if !c.beginUpdate.More() {
		c.beginUpdate = nil
	}

	return resp, nil
}
