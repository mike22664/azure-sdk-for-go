//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// BackupsClient contains the methods for the Backups group.
// Don't use this type directly, use NewBackupsClient() instead.
type BackupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupsClient creates a new instance of BackupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a backup under the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - backupName - The name of the backup
//   - body - Backup object supplied in the body of the operation.
//   - options - BackupsClientBeginCreateOptions contains the optional parameters for the BackupsClient.BeginCreate method.
func (client *BackupsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body Backup, options *BackupsClientBeginCreateOptions) (*runtime.Poller[BackupsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, backupVaultName, backupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a backup under the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *BackupsClient) create(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body Backup, options *BackupsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, backupName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *BackupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body Backup, options *BackupsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Backup under the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - backupName - The name of the backup
//   - options - BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
func (client *BackupsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, options *BackupsClientBeginDeleteOptions) (*runtime.Poller[BackupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, backupVaultName, backupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Backup under the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *BackupsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, options *BackupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, backupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BackupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, options *BackupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified Backup under Backup Vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - backupName - The name of the backup
//   - options - BackupsClientGetOptions contains the optional parameters for the BackupsClient.Get method.
func (client *BackupsClient) Get(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, options *BackupsClientGetOptions) (BackupsClientGetResponse, error) {
	var err error
	const operationName = "BackupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, backupName, options)
	if err != nil {
		return BackupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BackupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, options *BackupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
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

// getHandleResponse handles the Get response.
func (client *BackupsClient) getHandleResponse(resp *http.Response) (BackupsClientGetResponse, error) {
	result := BackupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Backup); err != nil {
		return BackupsClientGetResponse{}, err
	}
	return result, nil
}

// GetLatestStatus - Get the latest status of the backup for a volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - BackupsClientGetLatestStatusOptions contains the optional parameters for the BackupsClient.GetLatestStatus method.
func (client *BackupsClient) GetLatestStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetLatestStatusOptions) (BackupsClientGetLatestStatusResponse, error) {
	var err error
	const operationName = "BackupsClient.GetLatestStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getLatestStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsClientGetLatestStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupsClientGetLatestStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupsClientGetLatestStatusResponse{}, err
	}
	resp, err := client.getLatestStatusHandleResponse(httpResp)
	return resp, err
}

// getLatestStatusCreateRequest creates the GetLatestStatus request.
func (client *BackupsClient) getLatestStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetLatestStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/latestBackupStatus/current"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
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

// getLatestStatusHandleResponse handles the GetLatestStatus response.
func (client *BackupsClient) getLatestStatusHandleResponse(resp *http.Response) (BackupsClientGetLatestStatusResponse, error) {
	result := BackupsClientGetLatestStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupStatus); err != nil {
		return BackupsClientGetLatestStatusResponse{}, err
	}
	return result, nil
}

// GetVolumeLatestRestoreStatus - Get the latest status of the restore for a volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - BackupsClientGetVolumeLatestRestoreStatusOptions contains the optional parameters for the BackupsClient.GetVolumeLatestRestoreStatus
//     method.
func (client *BackupsClient) GetVolumeLatestRestoreStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetVolumeLatestRestoreStatusOptions) (BackupsClientGetVolumeLatestRestoreStatusResponse, error) {
	var err error
	const operationName = "BackupsClient.GetVolumeLatestRestoreStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getVolumeLatestRestoreStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsClientGetVolumeLatestRestoreStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupsClientGetVolumeLatestRestoreStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupsClientGetVolumeLatestRestoreStatusResponse{}, err
	}
	resp, err := client.getVolumeLatestRestoreStatusHandleResponse(httpResp)
	return resp, err
}

// getVolumeLatestRestoreStatusCreateRequest creates the GetVolumeLatestRestoreStatus request.
func (client *BackupsClient) getVolumeLatestRestoreStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetVolumeLatestRestoreStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/latestRestoreStatus/current"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
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

// getVolumeLatestRestoreStatusHandleResponse handles the GetVolumeLatestRestoreStatus response.
func (client *BackupsClient) getVolumeLatestRestoreStatusHandleResponse(resp *http.Response) (BackupsClientGetVolumeLatestRestoreStatusResponse, error) {
	result := BackupsClientGetVolumeLatestRestoreStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestoreStatus); err != nil {
		return BackupsClientGetVolumeLatestRestoreStatusResponse{}, err
	}
	return result, nil
}

// NewListByVaultPager - List all backups Under a Backup Vault
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - options - BackupsClientListByVaultOptions contains the optional parameters for the BackupsClient.NewListByVaultPager method.
func (client *BackupsClient) NewListByVaultPager(resourceGroupName string, accountName string, backupVaultName string, options *BackupsClientListByVaultOptions) *runtime.Pager[BackupsClientListByVaultResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupsClientListByVaultResponse]{
		More: func(page BackupsClientListByVaultResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupsClientListByVaultResponse) (BackupsClientListByVaultResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupsClient.NewListByVaultPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVaultCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, options)
			}, nil)
			if err != nil {
				return BackupsClientListByVaultResponse{}, err
			}
			return client.listByVaultHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVaultCreateRequest creates the ListByVault request.
func (client *BackupsClient) listByVaultCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, options *BackupsClientListByVaultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}/backups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVaultHandleResponse handles the ListByVault response.
func (client *BackupsClient) listByVaultHandleResponse(resp *http.Response) (BackupsClientListByVaultResponse, error) {
	result := BackupsClientListByVaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupsList); err != nil {
		return BackupsClientListByVaultResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch a Backup under the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - backupName - The name of the backup
//   - body - Backup object supplied in the body of the operation.
//   - options - BackupsClientBeginUpdateOptions contains the optional parameters for the BackupsClient.BeginUpdate method.
func (client *BackupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body BackupPatch, options *BackupsClientBeginUpdateOptions) (*runtime.Poller[BackupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, backupVaultName, backupName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch a Backup under the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *BackupsClient) update(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body BackupPatch, options *BackupsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, backupName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *BackupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body BackupPatch, options *BackupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}/backups/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
