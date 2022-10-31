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

// ImagesServer is a fake server for instances of the armcompute.ImagesClient type.
type ImagesServer struct {
	BeginCreateOrUpdate         func(ctx context.Context, resourceGroupName string, imageName string, parameters armcompute.Image, options *armcompute.ImagesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.ImagesClientCreateOrUpdateResponse], err azfake.ErrorResponder)
	BeginDelete                 func(ctx context.Context, resourceGroupName string, imageName string, options *armcompute.ImagesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.ImagesClientDeleteResponse], err azfake.ErrorResponder)
	Get                         func(ctx context.Context, resourceGroupName string, imageName string, options *armcompute.ImagesClientGetOptions) (resp azfake.Responder[armcompute.ImagesClientGetResponse], err azfake.ErrorResponder)
	NewListPager                func(options *armcompute.ImagesClientListOptions) (resp azfake.PagerResponder[armcompute.ImagesClientListResponse])
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.ImagesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.ImagesClientListByResourceGroupResponse])
	BeginUpdate                 func(ctx context.Context, resourceGroupName string, imageName string, parameters armcompute.ImageUpdate, options *armcompute.ImagesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.ImagesClientUpdateResponse], err azfake.ErrorResponder)
}

// NewImagesServerTransport creates a new instance of ImagesServerTransport with the provided implementation.
// The returned ImagesServerTransport instance is connected to an instance of armcompute.ImagesClient by way of the
// arm.ClientOptions.Transporter field.
func NewImagesServerTransport(srv *ImagesServer) *ImagesServerTransport {
	return &ImagesServerTransport{srv: srv}
}

// ImagesServerTransport connects instances of armcompute.ImagesClient to instances of ImagesServer.
// Don't use this type directly, use NewImagesServerTransport instead.
type ImagesServerTransport struct {
	srv                         *ImagesServer
	beginCreateOrUpdate         *azfake.PollerResponder[armcompute.ImagesClientCreateOrUpdateResponse]
	beginDelete                 *azfake.PollerResponder[armcompute.ImagesClientDeleteResponse]
	newListPager                *azfake.PagerResponder[armcompute.ImagesClientListResponse]
	newListByResourceGroupPager *azfake.PagerResponder[armcompute.ImagesClientListByResourceGroupResponse]
	beginUpdate                 *azfake.PollerResponder[armcompute.ImagesClientUpdateResponse]
}

// Do implements the policy.Transporter interface for ImagesServerTransport.
func (i *ImagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ImagesClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "ImagesClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "ImagesClient.Get":
		resp, err = i.dispatchGet(req)
	case "ImagesClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	case "ImagesClient.NewListByResourceGroupPager":
		resp, err = i.dispatchNewListByResourceGroupPager(req)
	case "ImagesClient.BeginUpdate":
		resp, err = i.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *ImagesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginCreateOrUpdate not implemented")}
	}
	if i.beginCreateOrUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/images/(?P<imageName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.Image](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := i.srv.BeginCreateOrUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("imageName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		i.beginCreateOrUpdate = &resp
	}

	resp, err := i.beginCreateOrUpdate.Next(req)
	if err != nil {
		return nil, err
	}

	if !i.beginCreateOrUpdate.More() {
		i.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (i *ImagesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("method BeginDelete not implemented")}
	}
	if i.beginDelete == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/images/(?P<imageName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp, errResp := i.srv.BeginDelete(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("imageName")], nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		i.beginDelete = &resp
	}

	resp, err := i.beginDelete.Next(req)
	if err != nil {
		return nil, err
	}

	if !i.beginDelete.More() {
		i.beginDelete = nil
	}

	return resp, nil
}

func (i *ImagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/images/(?P<imageName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	options := armcompute.ImagesClientGetOptions{
		Expand: getOptional(qp.Get("$expand")),
	}
	resp, errResp := i.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("imageName")], &options)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (i *ImagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if i.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/images"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListPager(nil)
		resp.InjectNextLinks(req, func(page *armcompute.ImagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		i.newListPager = &resp
	}

	resp, err := i.newListPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !i.newListPager.More() {
		i.newListPager = nil
	}
	return resp, nil
}

func (i *ImagesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByResourceGroupPager not implemented")}
	}
	if i.newListByResourceGroupPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/images"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListByResourceGroupPager(matches[regex.SubexpIndex("resourceGroupName")], nil)
		resp.InjectNextLinks(req, func(page *armcompute.ImagesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		i.newListByResourceGroupPager = &resp
	}

	resp, err := i.newListByResourceGroupPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !i.newListByResourceGroupPager.More() {
		i.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (i *ImagesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("method BeginUpdate not implemented")}
	}
	if i.beginUpdate == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/images/(?P<imageName>[a-zA-Z0-9-_]+)"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.ImageUpdate](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := i.srv.BeginUpdate(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("imageName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		i.beginUpdate = &resp
	}

	resp, err := i.beginUpdate.Next(req)
	if err != nil {
		return nil, err
	}

	if !i.beginUpdate.More() {
		i.beginUpdate = nil
	}

	return resp, nil
}
