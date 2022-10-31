//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DedicatedHostsClient contains the methods for the DedicatedHosts group.
// Don't use this type directly, use NewDedicatedHostsClient() instead.
type DedicatedHostsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDedicatedHostsClient creates a new instance of DedicatedHostsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDedicatedHostsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DedicatedHostsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DedicatedHostsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group.
//   - hostGroupName - The name of the dedicated host group.
//   - hostName - The name of the dedicated host .
//   - parameters - Parameters supplied to the Create Dedicated Host.
//   - options - DedicatedHostsClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedHostsClient.BeginCreateOrUpdate
//     method.
func (client *DedicatedHostsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DedicatedHostsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DedicatedHostsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DedicatedHostsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DedicatedHostsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostsClient.BeginCreateOrUpdate")
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedHostsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHost, options *DedicatedHostsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a dedicated host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group.
//   - hostGroupName - The name of the dedicated host group.
//   - hostName - The name of the dedicated host.
//   - options - DedicatedHostsClientBeginDeleteOptions contains the optional parameters for the DedicatedHostsClient.BeginDelete
//     method.
func (client *DedicatedHostsClient) BeginDelete(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginDeleteOptions) (*runtime.Poller[DedicatedHostsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hostGroupName, hostName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DedicatedHostsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DedicatedHostsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a dedicated host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DedicatedHostsClient) deleteOperation(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginDeleteOptions) (*http.Response, error) {
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostsClient.BeginDelete")
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedHostsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a dedicated host.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group.
//   - hostGroupName - The name of the dedicated host group.
//   - hostName - The name of the dedicated host.
//   - options - DedicatedHostsClientGetOptions contains the optional parameters for the DedicatedHostsClient.Get method.
func (client *DedicatedHostsClient) Get(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientGetOptions) (DedicatedHostsClientGetResponse, error) {
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostsClient.Get")
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return DedicatedHostsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedHostsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedHostsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedHostsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedHostsClient) getHandleResponse(resp *http.Response) (DedicatedHostsClientGetResponse, error) {
	result := DedicatedHostsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHost); err != nil {
		return DedicatedHostsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHostGroupPager - Lists all of the dedicated hosts in the specified dedicated host group. Use the nextLink property
// in the response to get the next page of dedicated hosts.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group.
//   - hostGroupName - The name of the dedicated host group.
//   - options - DedicatedHostsClientListByHostGroupOptions contains the optional parameters for the DedicatedHostsClient.ListByHostGroup
//     method.
func (client *DedicatedHostsClient) NewListByHostGroupPager(resourceGroupName string, hostGroupName string, options *DedicatedHostsClientListByHostGroupOptions) *runtime.Pager[DedicatedHostsClientListByHostGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DedicatedHostsClientListByHostGroupResponse]{
		More: func(page DedicatedHostsClientListByHostGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DedicatedHostsClientListByHostGroupResponse) (DedicatedHostsClientListByHostGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostsClient.NewListByHostGroupPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHostGroupCreateRequest(ctx, resourceGroupName, hostGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DedicatedHostsClientListByHostGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DedicatedHostsClientListByHostGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DedicatedHostsClientListByHostGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHostGroupHandleResponse(resp)
		},
	})
}

// listByHostGroupCreateRequest creates the ListByHostGroup request.
func (client *DedicatedHostsClient) listByHostGroupCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, options *DedicatedHostsClientListByHostGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHostGroupHandleResponse handles the ListByHostGroup response.
func (client *DedicatedHostsClient) listByHostGroupHandleResponse(resp *http.Response) (DedicatedHostsClientListByHostGroupResponse, error) {
	result := DedicatedHostsClientListByHostGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedHostListResult); err != nil {
		return DedicatedHostsClientListByHostGroupResponse{}, err
	}
	return result, nil
}

// BeginRestart - Restart the dedicated host. The operation will complete successfully once the dedicated host has restarted
// and is running. To determine the health of VMs deployed on the dedicated host after the
// restart check the Resource Health Center in the Azure Portal. Please refer to https://docs.microsoft.com/azure/service-health/resource-health-overview
// for more details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group.
//   - hostGroupName - The name of the dedicated host group.
//   - hostName - The name of the dedicated host.
//   - options - DedicatedHostsClientBeginRestartOptions contains the optional parameters for the DedicatedHostsClient.BeginRestart
//     method.
func (client *DedicatedHostsClient) BeginRestart(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginRestartOptions) (*runtime.Poller[DedicatedHostsClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, hostGroupName, hostName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DedicatedHostsClientRestartResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DedicatedHostsClientRestartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Restart - Restart the dedicated host. The operation will complete successfully once the dedicated host has restarted and
// is running. To determine the health of VMs deployed on the dedicated host after the
// restart check the Resource Health Center in the Azure Portal. Please refer to https://docs.microsoft.com/azure/service-health/resource-health-overview
// for more details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DedicatedHostsClient) restart(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginRestartOptions) (*http.Response, error) {
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostsClient.BeginRestart")
	req, err := client.restartCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *DedicatedHostsClient) restartCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, options *DedicatedHostsClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}/restart"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update an dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group.
//   - hostGroupName - The name of the dedicated host group.
//   - hostName - The name of the dedicated host .
//   - parameters - Parameters supplied to the Update Dedicated Host operation.
//   - options - DedicatedHostsClientBeginUpdateOptions contains the optional parameters for the DedicatedHostsClient.BeginUpdate
//     method.
func (client *DedicatedHostsClient) BeginUpdate(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsClientBeginUpdateOptions) (*runtime.Poller[DedicatedHostsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DedicatedHostsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DedicatedHostsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update an dedicated host .
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *DedicatedHostsClient) update(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsClientBeginUpdateOptions) (*http.Response, error) {
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DedicatedHostsClient.BeginUpdate")
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostGroupName, hostName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DedicatedHostsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostGroupName string, hostName string, parameters DedicatedHostUpdate, options *DedicatedHostsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostGroupName == "" {
		return nil, errors.New("parameter hostGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostGroupName}", url.PathEscape(hostGroupName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
