//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmanagedclusters

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

// NodeTypeSKUsClient contains the methods for the NodeTypeSKUs group.
// Don't use this type directly, use NewNodeTypeSKUsClient() instead.
type NodeTypeSKUsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNodeTypeSKUsClient creates a new instance of NodeTypeSKUsClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNodeTypeSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NodeTypeSKUsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NodeTypeSKUsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get a Service Fabric node type supported SKUs.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - nodeTypeName - The name of the node type.
//   - options - NodeTypeSKUsClientListOptions contains the optional parameters for the NodeTypeSKUsClient.NewListPager method.
func (client *NodeTypeSKUsClient) NewListPager(resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypeSKUsClientListOptions) *runtime.Pager[NodeTypeSKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[NodeTypeSKUsClientListResponse]{
		More: func(page NodeTypeSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NodeTypeSKUsClientListResponse) (NodeTypeSKUsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NodeTypeSKUsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, clusterName, nodeTypeName, options)
			}, nil)
			if err != nil {
				return NodeTypeSKUsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *NodeTypeSKUsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, nodeTypeName string, options *NodeTypeSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/managedClusters/{clusterName}/nodeTypes/{nodeTypeName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if nodeTypeName == "" {
		return nil, errors.New("parameter nodeTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeTypeName}", url.PathEscape(nodeTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NodeTypeSKUsClient) listHandleResponse(resp *http.Response) (NodeTypeSKUsClientListResponse, error) {
	result := NodeTypeSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NodeTypeListSKUResult); err != nil {
		return NodeTypeSKUsClientListResponse{}, err
	}
	return result, nil
}
