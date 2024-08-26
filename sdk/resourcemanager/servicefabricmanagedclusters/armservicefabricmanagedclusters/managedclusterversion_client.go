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

// ManagedClusterVersionClient contains the methods for the ManagedClusterVersion group.
// Don't use this type directly, use NewManagedClusterVersionClient() instead.
type ManagedClusterVersionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedClusterVersionClient creates a new instance of ManagedClusterVersionClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedClusterVersionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedClusterVersionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedClusterVersionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets information about an available Service Fabric managed cluster code version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - clusterVersion - The cluster code version.
//   - options - ManagedClusterVersionClientGetOptions contains the optional parameters for the ManagedClusterVersionClient.Get
//     method.
func (client *ManagedClusterVersionClient) Get(ctx context.Context, location string, clusterVersion string, options *ManagedClusterVersionClientGetOptions) (ManagedClusterVersionClientGetResponse, error) {
	var err error
	const operationName = "ManagedClusterVersionClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, clusterVersion, options)
	if err != nil {
		return ManagedClusterVersionClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterVersionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterVersionClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedClusterVersionClient) getCreateRequest(ctx context.Context, location string, clusterVersion string, options *ManagedClusterVersionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions/{clusterVersion}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterVersion == "" {
		return nil, errors.New("parameter clusterVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterVersion}", url.PathEscape(clusterVersion))
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

// getHandleResponse handles the Get response.
func (client *ManagedClusterVersionClient) getHandleResponse(resp *http.Response) (ManagedClusterVersionClientGetResponse, error) {
	result := ManagedClusterVersionClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterCodeVersionResult); err != nil {
		return ManagedClusterVersionClientGetResponse{}, err
	}
	return result, nil
}

// GetByEnvironment - Gets information about an available Service Fabric cluster code version by environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - environment - The operating system of the cluster. The default means all.
//   - clusterVersion - The cluster code version.
//   - options - ManagedClusterVersionClientGetByEnvironmentOptions contains the optional parameters for the ManagedClusterVersionClient.GetByEnvironment
//     method.
func (client *ManagedClusterVersionClient) GetByEnvironment(ctx context.Context, location string, environment ManagedClusterVersionEnvironment, clusterVersion string, options *ManagedClusterVersionClientGetByEnvironmentOptions) (ManagedClusterVersionClientGetByEnvironmentResponse, error) {
	var err error
	const operationName = "ManagedClusterVersionClient.GetByEnvironment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByEnvironmentCreateRequest(ctx, location, environment, clusterVersion, options)
	if err != nil {
		return ManagedClusterVersionClientGetByEnvironmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterVersionClientGetByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterVersionClientGetByEnvironmentResponse{}, err
	}
	resp, err := client.getByEnvironmentHandleResponse(httpResp)
	return resp, err
}

// getByEnvironmentCreateRequest creates the GetByEnvironment request.
func (client *ManagedClusterVersionClient) getByEnvironmentCreateRequest(ctx context.Context, location string, environment ManagedClusterVersionEnvironment, clusterVersion string, options *ManagedClusterVersionClientGetByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions/{clusterVersion}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if environment == "" {
		return nil, errors.New("parameter environment cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environment}", url.PathEscape(string(environment)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterVersion == "" {
		return nil, errors.New("parameter clusterVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterVersion}", url.PathEscape(clusterVersion))
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

// getByEnvironmentHandleResponse handles the GetByEnvironment response.
func (client *ManagedClusterVersionClient) getByEnvironmentHandleResponse(resp *http.Response) (ManagedClusterVersionClientGetByEnvironmentResponse, error) {
	result := ManagedClusterVersionClientGetByEnvironmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterCodeVersionResult); err != nil {
		return ManagedClusterVersionClientGetByEnvironmentResponse{}, err
	}
	return result, nil
}

// List - Gets all available code versions for Service Fabric cluster resources by location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - options - ManagedClusterVersionClientListOptions contains the optional parameters for the ManagedClusterVersionClient.List
//     method.
func (client *ManagedClusterVersionClient) List(ctx context.Context, location string, options *ManagedClusterVersionClientListOptions) (ManagedClusterVersionClientListResponse, error) {
	var err error
	const operationName = "ManagedClusterVersionClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return ManagedClusterVersionClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterVersionClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterVersionClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *ManagedClusterVersionClient) listCreateRequest(ctx context.Context, location string, options *ManagedClusterVersionClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *ManagedClusterVersionClient) listHandleResponse(resp *http.Response) (ManagedClusterVersionClientListResponse, error) {
	result := ManagedClusterVersionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterCodeVersionResultArray); err != nil {
		return ManagedClusterVersionClientListResponse{}, err
	}
	return result, nil
}

// ListByEnvironment - Gets all available code versions for Service Fabric cluster resources by environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - environment - The operating system of the cluster. The default means all.
//   - options - ManagedClusterVersionClientListByEnvironmentOptions contains the optional parameters for the ManagedClusterVersionClient.ListByEnvironment
//     method.
func (client *ManagedClusterVersionClient) ListByEnvironment(ctx context.Context, location string, environment ManagedClusterVersionEnvironment, options *ManagedClusterVersionClientListByEnvironmentOptions) (ManagedClusterVersionClientListByEnvironmentResponse, error) {
	var err error
	const operationName = "ManagedClusterVersionClient.ListByEnvironment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByEnvironmentCreateRequest(ctx, location, environment, options)
	if err != nil {
		return ManagedClusterVersionClientListByEnvironmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedClusterVersionClientListByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedClusterVersionClientListByEnvironmentResponse{}, err
	}
	resp, err := client.listByEnvironmentHandleResponse(httpResp)
	return resp, err
}

// listByEnvironmentCreateRequest creates the ListByEnvironment request.
func (client *ManagedClusterVersionClient) listByEnvironmentCreateRequest(ctx context.Context, location string, environment ManagedClusterVersionEnvironment, options *ManagedClusterVersionClientListByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/managedClusterVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if environment == "" {
		return nil, errors.New("parameter environment cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environment}", url.PathEscape(string(environment)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listByEnvironmentHandleResponse handles the ListByEnvironment response.
func (client *ManagedClusterVersionClient) listByEnvironmentHandleResponse(resp *http.Response) (ManagedClusterVersionClientListByEnvironmentResponse, error) {
	result := ManagedClusterVersionClientListByEnvironmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedClusterCodeVersionResultArray); err != nil {
		return ManagedClusterVersionClientListByEnvironmentResponse{}, err
	}
	return result, nil
}
