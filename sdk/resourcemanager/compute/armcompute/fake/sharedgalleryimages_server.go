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

// SharedGalleryImagesServer is a fake server for instances of the armcompute.SharedGalleryImagesClient type.
type SharedGalleryImagesServer struct {
	Get          func(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, options *armcompute.SharedGalleryImagesClientGetOptions) (resp azfake.Responder[armcompute.SharedGalleryImagesClientGetResponse], err azfake.ErrorResponder)
	NewListPager func(location string, galleryUniqueName string, options *armcompute.SharedGalleryImagesClientListOptions) (resp azfake.PagerResponder[armcompute.SharedGalleryImagesClientListResponse])
}

// NewSharedGalleryImagesServerTransport creates a new instance of SharedGalleryImagesServerTransport with the provided implementation.
// The returned SharedGalleryImagesServerTransport instance is connected to an instance of armcompute.SharedGalleryImagesClient by way of the
// arm.ClientOptions.Transporter field.
func NewSharedGalleryImagesServerTransport(srv *SharedGalleryImagesServer) *SharedGalleryImagesServerTransport {
	return &SharedGalleryImagesServerTransport{srv: srv}
}

// SharedGalleryImagesServerTransport connects instances of armcompute.SharedGalleryImagesClient to instances of SharedGalleryImagesServer.
// Don't use this type directly, use NewSharedGalleryImagesServerTransport instead.
type SharedGalleryImagesServerTransport struct {
	srv          *SharedGalleryImagesServer
	newListPager *azfake.PagerResponder[armcompute.SharedGalleryImagesClientListResponse]
}

// Do implements the policy.Transporter interface for SharedGalleryImagesServerTransport.
func (s *SharedGalleryImagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SharedGalleryImagesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SharedGalleryImagesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SharedGalleryImagesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("method Get not implemented")}
	}
	const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/sharedGalleries/(?P<galleryUniqueName>[a-zA-Z0-9-_]+)/images/(?P<galleryImageName>[a-zA-Z0-9-_]+)"
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.Path)
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resp, errResp := s.srv.Get(req.Context(), matches[regex.SubexpIndex("location")], matches[regex.SubexpIndex("galleryUniqueName")], matches[regex.SubexpIndex("galleryImageName")], nil)
	if err := errResp.Get(req); err != nil {
		return nil, err
	}
	return resp.MarshalResponseAsJSON(req)
}

func (s *SharedGalleryImagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("method NewListPager not implemented")}
	}
	if s.newListPager == nil {
		const regexStr = "/subscriptions/(?P<subscriptionId>[a-zA-Z0-9-_]+)/providers/Microsoft.Compute/locations/(?P<location>[a-zA-Z0-9-_]+)/sharedGalleries/(?P<galleryUniqueName>[a-zA-Z0-9-_]+)/images"
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.Path)
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		options := armcompute.SharedGalleryImagesClientListOptions{
			SharedTo: getOptional(armcompute.SharedToValues(qp.Get("sharedTo"))),
		}
		resp := s.srv.NewListPager(matches[regex.SubexpIndex("location")], matches[regex.SubexpIndex("galleryUniqueName")], &options)
		resp.InjectNextLinks(req, func(page *armcompute.SharedGalleryImagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
		s.newListPager = &resp
	}

	resp, err := s.newListPager.Next(req)
	if err != nil {
		return nil, err
	}
	if !s.newListPager.More() {
		s.newListPager = nil
	}
	return resp, nil
}
