//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

// AvailableClusterPoolVersionsClientListByLocationResponse contains the response from method AvailableClusterPoolVersionsClient.NewListByLocationPager.
type AvailableClusterPoolVersionsClientListByLocationResponse struct {
	// Represents a list of cluster pool versions.
	ClusterPoolVersionsListResult
}

// AvailableClusterVersionsClientListByLocationResponse contains the response from method AvailableClusterVersionsClient.NewListByLocationPager.
type AvailableClusterVersionsClientListByLocationResponse struct {
	// Represents a list of cluster versions.
	ClusterVersionsListResult
}

// ClusterAvailableUpgradesClientListResponse contains the response from method ClusterAvailableUpgradesClient.NewListPager.
type ClusterAvailableUpgradesClientListResponse struct {
	// Collection of cluster available upgrade.
	ClusterAvailableUpgradeList
}

// ClusterJobsClientListResponse contains the response from method ClusterJobsClient.NewListPager.
type ClusterJobsClientListResponse struct {
	// Collection of cluster job.
	ClusterJobList
}

// ClusterJobsClientRunJobResponse contains the response from method ClusterJobsClient.BeginRunJob.
type ClusterJobsClientRunJobResponse struct {
	// Cluster job.
	ClusterJob
}

// ClusterLibrariesClientListResponse contains the response from method ClusterLibrariesClient.NewListPager.
type ClusterLibrariesClientListResponse struct {
	// Collection of libraries in the cluster.
	ClusterLibraryList
}

// ClusterLibrariesClientManageLibrariesResponse contains the response from method ClusterLibrariesClient.BeginManageLibraries.
type ClusterLibrariesClientManageLibrariesResponse struct {
	// placeholder for future response values
}

// ClusterPoolAvailableUpgradesClientListResponse contains the response from method ClusterPoolAvailableUpgradesClient.NewListPager.
type ClusterPoolAvailableUpgradesClientListResponse struct {
	// collection of cluster pool available upgrade.
	ClusterPoolAvailableUpgradeList
}

// ClusterPoolUpgradeHistoriesClientListResponse contains the response from method ClusterPoolUpgradeHistoriesClient.NewListPager.
type ClusterPoolUpgradeHistoriesClientListResponse struct {
	// Represents a list of cluster pool upgrade history.
	ClusterPoolUpgradeHistoryListResult
}

// ClusterPoolsClientCreateOrUpdateResponse contains the response from method ClusterPoolsClient.BeginCreateOrUpdate.
type ClusterPoolsClientCreateOrUpdateResponse struct {
	// Cluster pool.
	ClusterPool
}

// ClusterPoolsClientDeleteResponse contains the response from method ClusterPoolsClient.BeginDelete.
type ClusterPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// ClusterPoolsClientGetResponse contains the response from method ClusterPoolsClient.Get.
type ClusterPoolsClientGetResponse struct {
	// Cluster pool.
	ClusterPool
}

// ClusterPoolsClientListByResourceGroupResponse contains the response from method ClusterPoolsClient.NewListByResourceGroupPager.
type ClusterPoolsClientListByResourceGroupResponse struct {
	// The list cluster pools operation response.
	ClusterPoolListResult
}

// ClusterPoolsClientListBySubscriptionResponse contains the response from method ClusterPoolsClient.NewListBySubscriptionPager.
type ClusterPoolsClientListBySubscriptionResponse struct {
	// The list cluster pools operation response.
	ClusterPoolListResult
}

// ClusterPoolsClientUpdateTagsResponse contains the response from method ClusterPoolsClient.BeginUpdateTags.
type ClusterPoolsClientUpdateTagsResponse struct {
	// Cluster pool.
	ClusterPool
}

// ClusterPoolsClientUpgradeResponse contains the response from method ClusterPoolsClient.BeginUpgrade.
type ClusterPoolsClientUpgradeResponse struct {
	// Cluster pool.
	ClusterPool
}

// ClusterUpgradeHistoriesClientListResponse contains the response from method ClusterUpgradeHistoriesClient.NewListPager.
type ClusterUpgradeHistoriesClientListResponse struct {
	// Represents a list of cluster upgrade history.
	ClusterUpgradeHistoryListResult
}

// ClustersClientCreateResponse contains the response from method ClustersClient.BeginCreate.
type ClustersClientCreateResponse struct {
	// The cluster.
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetInstanceViewResponse contains the response from method ClustersClient.GetInstanceView.
type ClustersClientGetInstanceViewResponse struct {
	// Cluster Instance View.
	ClusterInstanceViewResult
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	// The cluster.
	Cluster
}

// ClustersClientListByClusterPoolNameResponse contains the response from method ClustersClient.NewListByClusterPoolNamePager.
type ClustersClientListByClusterPoolNameResponse struct {
	// The list cluster operation response.
	ClusterListResult
}

// ClustersClientListInstanceViewsResponse contains the response from method ClustersClient.NewListInstanceViewsPager.
type ClustersClientListInstanceViewsResponse struct {
	// The instance view of a HDInsight Cluster.
	ClusterInstanceViewsResult
}

// ClustersClientListServiceConfigsResponse contains the response from method ClustersClient.NewListServiceConfigsPager.
type ClustersClientListServiceConfigsResponse struct {
	// Cluster instance service configs api response.
	ServiceConfigListResult
}

// ClustersClientResizeResponse contains the response from method ClustersClient.BeginResize.
type ClustersClientResizeResponse struct {
	// The cluster.
	Cluster
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.BeginUpdate.
type ClustersClientUpdateResponse struct {
	// The cluster.
	Cluster
}

// ClustersClientUpgradeManualRollbackResponse contains the response from method ClustersClient.BeginUpgradeManualRollback.
type ClustersClientUpgradeManualRollbackResponse struct {
	// The cluster.
	Cluster
}

// ClustersClientUpgradeResponse contains the response from method ClustersClient.BeginUpgrade.
type ClustersClientUpgradeResponse struct {
	// The cluster.
	Cluster
}

// LocationsClientCheckNameAvailabilityResponse contains the response from method LocationsClient.CheckNameAvailability.
type LocationsClientCheckNameAvailabilityResponse struct {
	// Result of check name availability.
	NameAvailabilityResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}
