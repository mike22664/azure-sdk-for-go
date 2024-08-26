//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedEnvironmentUsagesClient contains the methods for the ManagedEnvironmentUsages group.
// Don't use this type directly, use NewManagedEnvironmentUsagesClient() instead.
type ManagedEnvironmentUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedEnvironmentUsagesClient creates a new instance of ManagedEnvironmentUsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedEnvironmentUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedEnvironmentUsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedEnvironmentUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets the current usage information as well as the limits for environment.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Environment.
//   - options - ManagedEnvironmentUsagesClientListOptions contains the optional parameters for the ManagedEnvironmentUsagesClient.NewListPager
//     method.
func (client *ManagedEnvironmentUsagesClient) NewListPager(resourceGroupName string, environmentName string, options *ManagedEnvironmentUsagesClientListOptions) *runtime.Pager[ManagedEnvironmentUsagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedEnvironmentUsagesClientListResponse]{
		More: func(page ManagedEnvironmentUsagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedEnvironmentUsagesClientListResponse) (ManagedEnvironmentUsagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedEnvironmentUsagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			}, nil)
			if err != nil {
				return ManagedEnvironmentUsagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedEnvironmentUsagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *ManagedEnvironmentUsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedEnvironmentUsagesClient) listHandleResponse(resp *http.Response) (ManagedEnvironmentUsagesClientListResponse, error) {
	result := ManagedEnvironmentUsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUsagesResult); err != nil {
		return ManagedEnvironmentUsagesClientListResponse{}, err
	}
	return result, nil
}
