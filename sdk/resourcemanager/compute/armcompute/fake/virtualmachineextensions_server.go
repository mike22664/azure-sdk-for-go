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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"net/http"
	"regexp"
)

// VirtualMachineExtensionsServer is a fake server for instances of the armcompute.VirtualMachineExtensionsClient type.
type VirtualMachineExtensionsServer struct {
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters armcompute.VirtualMachineExtension, options *armcompute.VirtualMachineExtensionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineExtensionsClientCreateOrUpdateResponse], err azfake.ErrorResponder)
	BeginDelete         func(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *armcompute.VirtualMachineExtensionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineExtensionsClientDeleteResponse], err azfake.ErrorResponder)
	Get                 func(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *armcompute.VirtualMachineExtensionsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionsClientGetResponse], err azfake.ErrorResponder)
	List                func(ctx context.Context, resourceGroupName string, vmName string, options *armcompute.VirtualMachineExtensionsClientListOptions) (resp azfake.Responder[armcompute.VirtualMachineExtensionsClientListResponse], err azfake.ErrorResponder)
	BeginUpdate         func(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters armcompute.VirtualMachineExtensionUpdate, options *armcompute.VirtualMachineExtensionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineExtensionsClientUpdateResponse], err azfake.ErrorResponder)
}

// NewVirtualMachineExtensionsServerTransport creates a new instance of VirtualMachineExtensionsServerTransport with the provided implementation.
// The returned VirtualMachineExtensionsServerTransport instance is connected to an instance of armcompute.VirtualMachineExtensionsClient by way of the
// arm.ClientOptions.Transporter field.
func NewVirtualMachineExtensionsServerTransport(srv *VirtualMachineExtensionsServer) *VirtualMachineExtensionsServerTransport {
	return &VirtualMachineExtensionsServerTransport{srv: srv}
}

// VirtualMachineExtensionsServerTransport connects instances of armcompute.VirtualMachineExtensionsClient to instances of VirtualMachineExtensionsServer.
// Don't use this type directly, use NewVirtualMachineExtensionsServerTransport instead.
type VirtualMachineExtensionsServerTransport struct {
	srv                 *VirtualMachineExtensionsServer
	beginCreateOrUpdate *azfake.PollerResponder[armcompute.VirtualMachineExtensionsClientCreateOrUpdateResponse]
	beginDelete         *azfake.PollerResponder[armcompute.VirtualMachineExtensionsClientDeleteResponse]
	beginUpdate         *azfake.PollerResponder[armcompute.VirtualMachineExtensionsClientUpdateResponse]
}

// Do implements the policy.Transporter interface for VirtualMachineExtensionsServerTransport.
func (v *VirtualMachineExtensionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineExtensionsClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualMachineExtensionsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineExtensionsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineExtensionsClient.List":
		resp, err = v.dispatchList(req)
	case "VirtualMachineExtensionsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineExtensionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if v.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.VirtualMachineExtension](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := v.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("vmExtensionName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		v.beginCreateOrUpdate = &resp
	}

	resp, err := v.beginCreateOrUpdate.Next(req)
	if err != nil {
		return nil, err
	}

	if !v.beginCreateOrUpdate.More() {
		v.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (v *VirtualMachineExtensionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if v.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp, errResp := v.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("vmExtensionName")], nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		v.beginDelete = &resp
	}

	resp, err := v.beginDelete.Next(req)
	if err != nil {
		return nil, err
	}

	if !v.beginDelete.More() {
		v.beginDelete = nil
	}

	return resp, nil
}

func (v *VirtualMachineExtensionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	options := armcompute.VirtualMachineExtensionsClientGetOptions{
		Expand: getOptional(qp.Get("$expand")),
	}
	resp, errResp := v.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("vmExtensionName")], &options)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (v *VirtualMachineExtensionsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("method List not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/extensions"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	options := armcompute.VirtualMachineExtensionsClientListOptions{
		Expand: getOptional(qp.Get("$expand")),
	}
	resp, errResp := v.srv.List(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], &options)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (v *VirtualMachineExtensionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if v.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/virtualMachines/(?P<vmName>[a-zA-Z0-9-_]+)/extensions/(?P<vmExtensionName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.VirtualMachineExtensionUpdate](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := v.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("vmName")], matches[regex.SubexpIndex("vmExtensionName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		v.beginUpdate = &resp
	}

	resp, err := v.beginUpdate.Next(req)
	if err != nil {
		return nil, err
	}

	if !v.beginUpdate.More() {
		v.beginUpdate = nil
	}

	return resp, nil
}
