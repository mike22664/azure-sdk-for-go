//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TriggerRunsClient contains the methods for the TriggerRuns group.
// Don't use this type directly, use NewTriggerRunsClient() instead.
type TriggerRunsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTriggerRunsClient creates a new instance of TriggerRunsClient with the specified values.
func NewTriggerRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TriggerRunsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TriggerRunsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Cancel - Cancel a single trigger instance by runId.
// If the operation fails it returns the *CloudError error type.
func (client *TriggerRunsClient) Cancel(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *TriggerRunsCancelOptions) (TriggerRunsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, factoryName, triggerName, runID, options)
	if err != nil {
		return TriggerRunsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunsCancelResponse{}, client.cancelHandleError(resp)
	}
	return TriggerRunsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *TriggerRunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *TriggerRunsCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/triggerRuns/{runId}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *TriggerRunsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// QueryByFactory - Query trigger runs.
// If the operation fails it returns the *CloudError error type.
func (client *TriggerRunsClient) QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters, options *TriggerRunsQueryByFactoryOptions) (TriggerRunsQueryByFactoryResponse, error) {
	req, err := client.queryByFactoryCreateRequest(ctx, resourceGroupName, factoryName, filterParameters, options)
	if err != nil {
		return TriggerRunsQueryByFactoryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunsQueryByFactoryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunsQueryByFactoryResponse{}, client.queryByFactoryHandleError(resp)
	}
	return client.queryByFactoryHandleResponse(resp)
}

// queryByFactoryCreateRequest creates the QueryByFactory request.
func (client *TriggerRunsClient) queryByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, filterParameters RunFilterParameters, options *TriggerRunsQueryByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryTriggerRuns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, filterParameters)
}

// queryByFactoryHandleResponse handles the QueryByFactory response.
func (client *TriggerRunsClient) queryByFactoryHandleResponse(resp *http.Response) (TriggerRunsQueryByFactoryResponse, error) {
	result := TriggerRunsQueryByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerRunsQueryResponse); err != nil {
		return TriggerRunsQueryByFactoryResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// queryByFactoryHandleError handles the QueryByFactory error response.
func (client *TriggerRunsClient) queryByFactoryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Rerun - Rerun single trigger instance by runId.
// If the operation fails it returns the *CloudError error type.
func (client *TriggerRunsClient) Rerun(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *TriggerRunsRerunOptions) (TriggerRunsRerunResponse, error) {
	req, err := client.rerunCreateRequest(ctx, resourceGroupName, factoryName, triggerName, runID, options)
	if err != nil {
		return TriggerRunsRerunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TriggerRunsRerunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TriggerRunsRerunResponse{}, client.rerunHandleError(resp)
	}
	return TriggerRunsRerunResponse{RawResponse: resp}, nil
}

// rerunCreateRequest creates the Rerun request.
func (client *TriggerRunsClient) rerunCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string, options *TriggerRunsRerunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/triggerRuns/{runId}/rerun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// rerunHandleError handles the Rerun error response.
func (client *TriggerRunsClient) rerunHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
