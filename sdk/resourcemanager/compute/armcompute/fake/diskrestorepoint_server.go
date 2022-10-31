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

// DiskRestorePointServer is a fake server for instances of the armcompute.DiskRestorePointClient type.
type DiskRestorePointServer struct {
	Get                        func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *armcompute.DiskRestorePointClientGetOptions) (resp azfake.Responder[armcompute.DiskRestorePointClientGetResponse], err azfake.ErrorResponder)
	BeginGrantAccess           func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, grantAccessData armcompute.GrantAccessData, options *armcompute.DiskRestorePointClientBeginGrantAccessOptions) (resp azfake.PollerResponder[armcompute.DiskRestorePointClientGrantAccessResponse], err azfake.ErrorResponder)
	NewListByRestorePointPager func(resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, options *armcompute.DiskRestorePointClientListByRestorePointOptions) (resp azfake.PagerResponder[armcompute.DiskRestorePointClientListByRestorePointResponse])
	BeginRevokeAccess          func(ctx context.Context, resourceGroupName string, restorePointCollectionName string, vmRestorePointName string, diskRestorePointName string, options *armcompute.DiskRestorePointClientBeginRevokeAccessOptions) (resp azfake.PollerResponder[armcompute.DiskRestorePointClientRevokeAccessResponse], err azfake.ErrorResponder)
}

// NewDiskRestorePointServerTransport creates a new instance of DiskRestorePointServerTransport with the provided implementation.
// The returned DiskRestorePointServerTransport instance is connected to an instance of armcompute.DiskRestorePointClient by way of the
// arm.ClientOptions.Transporter field.
func NewDiskRestorePointServerTransport(srv *DiskRestorePointServer) *DiskRestorePointServerTransport {
	return &DiskRestorePointServerTransport{srv: srv}
}

// DiskRestorePointServerTransport connects instances of armcompute.DiskRestorePointClient to instances of DiskRestorePointServer.
// Don't use this type directly, use NewDiskRestorePointServerTransport instead.
type DiskRestorePointServerTransport struct {
	srv                        *DiskRestorePointServer
	beginGrantAccess           *azfake.PollerResponder[armcompute.DiskRestorePointClientGrantAccessResponse]
	newListByRestorePointPager *azfake.PagerResponder[armcompute.DiskRestorePointClientListByRestorePointResponse]
	beginRevokeAccess          *azfake.PollerResponder[armcompute.DiskRestorePointClientRevokeAccessResponse]
}

// Do implements the policy.Transporter interface for DiskRestorePointServerTransport.
func (d *DiskRestorePointServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DiskRestorePointClient.Get":
		resp, err = d.dispatchGet(req)
	case "DiskRestorePointClient.BeginGrantAccess":
		resp, err = d.dispatchBeginGrantAccess(req)
	case "DiskRestorePointClient.NewListByRestorePointPager":
		resp, err = d.dispatchNewListByRestorePointPager(req)
	case "DiskRestorePointClient.BeginRevokeAccess":
		resp, err = d.dispatchBeginRevokeAccess(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[a-zA-Z0-9-_]+)/restorePoints/(?P<vmRestorePointName>[a-zA-Z0-9-_]+)/diskRestorePoints/(?P<diskRestorePointName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resp, errResp := d.srv.Get(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("restorePointCollectionName")], matches[regex.SubexpIndex("vmRestorePointName")], matches[regex.SubexpIndex("diskRestorePointName")], nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (d *DiskRestorePointServerTransport) dispatchBeginGrantAccess(req *http.Request) (*http.Response, error) {
	if d.srv.BeginGrantAccess == nil {
		return nil, &nonRetriableError{errors.New("method BeginGrantAccess not implemented")}
	}
	if d.beginGrantAccess == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[a-zA-Z0-9-_]+)/restorePoints/(?P<vmRestorePointName>[a-zA-Z0-9-_]+)/diskRestorePoints/(?P<diskRestorePointName>[a-zA-Z0-9-_]+)/beginGetAccess"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := azfake.UnmarshalRequestAsJSON[armcompute.GrantAccessData](req)
		if err != nil {
			return nil, err
		}
		resp, errResp := d.srv.BeginGrantAccess(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("restorePointCollectionName")], matches[regex.SubexpIndex("vmRestorePointName")], matches[regex.SubexpIndex("diskRestorePointName")], body, nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		d.beginGrantAccess = &resp
	}

	resp, err := d.beginGrantAccess.Next(req)
	if err != nil {
		return nil, err
	}

	if !d.beginGrantAccess.More() {
		d.beginGrantAccess = nil
	}

	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchNewListByRestorePointPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByRestorePointPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListByRestorePointPager not implemented")}
	}
	if d.newListByRestorePointPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[a-zA-Z0-9-_]+)/restorePoints/(?P<vmRestorePointName>[a-zA-Z0-9-_]+)/diskRestorePoints"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListByRestorePointPager(matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("restorePointCollectionName")], matches[regex.SubexpIndex("vmRestorePointName")], nil)
		resp.InjectNextLinks(req, func(page *armcompute.DiskRestorePointClientListByRestorePointResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		d.newListByRestorePointPager = &resp
	}

	resp, err := d.newListByRestorePointPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !d.newListByRestorePointPager.More() {
		d.newListByRestorePointPager = nil
	}
	return resp, nil
}

func (d *DiskRestorePointServerTransport) dispatchBeginRevokeAccess(req *http.Request) (*http.Response, error) {
	if d.srv.BeginRevokeAccess == nil {
		return nil, &nonRetriableError{errors.New("method BeginRevokeAccess not implemented")}
	}
	if d.beginRevokeAccess == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/resourceGroups/(?P<resourceGroupName>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/restorePointCollections/(?P<restorePointCollectionName>[a-zA-Z0-9-_]+)/restorePoints/(?P<vmRestorePointName>[a-zA-Z0-9-_]+)/diskRestorePoints/(?P<diskRestorePointName>[a-zA-Z0-9-_]+)/endGetAccess"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp, errResp := d.srv.BeginRevokeAccess(req.Context(), matches[regex.SubexpIndex("resourceGroupName")], matches[regex.SubexpIndex("restorePointCollectionName")], matches[regex.SubexpIndex("vmRestorePointName")], matches[regex.SubexpIndex("diskRestorePointName")], nil)
		if err := errResp.Get(req); err != nil {
			return nil, err
		}
		d.beginRevokeAccess = &resp
	}

	resp, err := d.beginRevokeAccess.Next(req)
	if err != nil {
		return nil, err
	}

	if !d.beginRevokeAccess.More() {
		d.beginRevokeAccess = nil
	}

	return resp, nil
}
