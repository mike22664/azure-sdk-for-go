//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/IpGroupsGet.json
func ExampleIPGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPGroupsClient().Get(ctx, "myResourceGroup", "ipGroups1", &armnetwork.IPGroupsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPGroup = armnetwork.IPGroup{
	// 	Name: to.Ptr("ipGroups1"),
	// 	Type: to.Ptr("Microsoft.Network/ipGroups"),
	// 	ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups/ipGroups1"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.IPGroupPropertiesFormat{
	// 		FirewallPolicies: []*armnetwork.SubResource{
	// 		},
	// 		Firewalls: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 		}},
	// 		IPAddresses: []*string{
	// 			to.Ptr("13.64.39.16/32"),
	// 			to.Ptr("40.74.146.80/31"),
	// 			to.Ptr("40.74.147.32/28")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/IpGroupsCreate.json
func ExampleIPGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPGroupsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "ipGroups1", armnetwork.IPGroup{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.IPGroupPropertiesFormat{
			IPAddresses: []*string{
				to.Ptr("13.64.39.16/32"),
				to.Ptr("40.74.146.80/31"),
				to.Ptr("40.74.147.32/28")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPGroup = armnetwork.IPGroup{
	// 	Name: to.Ptr("ipGroups1"),
	// 	Type: to.Ptr("Microsoft.Network/ipGroups"),
	// 	ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups/ipGroups1"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.IPGroupPropertiesFormat{
	// 		FirewallPolicies: []*armnetwork.SubResource{
	// 		},
	// 		Firewalls: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 		}},
	// 		IPAddresses: []*string{
	// 			to.Ptr("13.64.39.16/32"),
	// 			to.Ptr("40.74.146.80/31"),
	// 			to.Ptr("40.74.147.32/28")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/IpGroupsUpdateTags.json
func ExampleIPGroupsClient_UpdateGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPGroupsClient().UpdateGroups(ctx, "myResourceGroup", "ipGroups1", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPGroup = armnetwork.IPGroup{
	// 	Name: to.Ptr("ipGroups1"),
	// 	Type: to.Ptr("Microsoft.Network/ipGroups"),
	// 	ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups/ipGroups1"),
	// 	Location: to.Ptr("westcentralus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.IPGroupPropertiesFormat{
	// 		FirewallPolicies: []*armnetwork.SubResource{
	// 		},
	// 		Firewalls: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
	// 		}},
	// 		IPAddresses: []*string{
	// 			to.Ptr("13.64.39.16/32"),
	// 			to.Ptr("40.74.146.80/31"),
	// 			to.Ptr("40.74.147.32/28")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/IpGroupsDelete.json
func ExampleIPGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPGroupsClient().BeginDelete(ctx, "myResourceGroup", "ipGroups1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/IpGroupsListByResourceGroup.json
func ExampleIPGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPGroupsClient().NewListByResourceGroupPager("myResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IPGroupListResult = armnetwork.IPGroupListResult{
		// 	Value: []*armnetwork.IPGroup{
		// 		{
		// 			Name: to.Ptr("ipGroups1"),
		// 			Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.IPGroupPropertiesFormat{
		// 				FirewallPolicies: []*armnetwork.SubResource{
		// 				},
		// 				Firewalls: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 				}},
		// 				IPAddresses: []*string{
		// 					to.Ptr("13.64.39.16/32"),
		// 					to.Ptr("40.74.146.80/31"),
		// 					to.Ptr("40.74.147.32/28")},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ipGroups2"),
		// 				Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup/ipGroups"),
		// 				Location: to.Ptr("centralus"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				Properties: &armnetwork.IPGroupPropertiesFormat{
		// 					FirewallPolicies: []*armnetwork.SubResource{
		// 					},
		// 					Firewalls: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 					}},
		// 					IPAddresses: []*string{
		// 						to.Ptr("14.64.39.16/32"),
		// 						to.Ptr("41.74.146.80/31"),
		// 						to.Ptr("42.74.147.32/28")},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/IpGroupsListBySubscription.json
func ExampleIPGroupsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPGroupsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IPGroupListResult = armnetwork.IPGroupListResult{
		// 	Value: []*armnetwork.IPGroup{
		// 		{
		// 			Name: to.Ptr("iptag1"),
		// 			Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup1/ipGroups"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.IPGroupPropertiesFormat{
		// 				FirewallPolicies: []*armnetwork.SubResource{
		// 				},
		// 				Firewalls: []*armnetwork.SubResource{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 				}},
		// 				IPAddresses: []*string{
		// 					to.Ptr("13.64.39.16/32"),
		// 					to.Ptr("40.74.146.80/31"),
		// 					to.Ptr("40.74.147.32/28")},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("iptag2"),
		// 				Type: to.Ptr("Microsoft.Network/ipGroups"),
		// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Network/resourceGroup/myResourceGroup2/ipGroups"),
		// 				Location: to.Ptr("centralus"),
		// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 				Properties: &armnetwork.IPGroupPropertiesFormat{
		// 					FirewallPolicies: []*armnetwork.SubResource{
		// 					},
		// 					Firewalls: []*armnetwork.SubResource{
		// 						{
		// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azurefirewall"),
		// 					}},
		// 					IPAddresses: []*string{
		// 						to.Ptr("14.64.39.16/32"),
		// 						to.Ptr("41.74.146.80/31"),
		// 						to.Ptr("42.74.147.32/28")},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					},
		// 			}},
		// 		}
	}
}
