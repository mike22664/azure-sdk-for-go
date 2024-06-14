//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/30c95a2a424a347f27ac78b6bdefd37f71314b7e/specification/azurefleet/resource-manager/Microsoft.AzureFleet/preview/2024-05-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputefleet.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armcomputefleet.OperationListResult{
		// 	Value: []*armcomputefleet.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureFleet/fleets/read"),
		// 			Display: &armcomputefleet.OperationDisplay{
		// 				Description: to.Ptr("Get properties of Azure Fleet resource"),
		// 				Operation: to.Ptr("Get Azure Fleet"),
		// 				Provider: to.Ptr("Microsoft Azure Fleet"),
		// 				Resource: to.Ptr("Fleets"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureFleet/fleets/write"),
		// 			Display: &armcomputefleet.OperationDisplay{
		// 				Description: to.Ptr("Creates a new Azure Fleet resource or updates an existing one"),
		// 				Operation: to.Ptr("Create or Update Azure Fleet"),
		// 				Provider: to.Ptr("Microsoft Azure Fleet"),
		// 				Resource: to.Ptr("Fleets"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureFleet/fleets/delete"),
		// 			Display: &armcomputefleet.OperationDisplay{
		// 				Description: to.Ptr("Deletes all compute resources of Azure Fleet resource"),
		// 				Operation: to.Ptr("Delete Virtual Machine and Virtual Machine scale sets in a Azure Fleet resource"),
		// 				Provider: to.Ptr("Microsoft Azure Fleet"),
		// 				Resource: to.Ptr("Fleets"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureFleet/register/action"),
		// 			Display: &armcomputefleet.OperationDisplay{
		// 				Description: to.Ptr("Registers Subscription with Microsoft.AzureFleet resource provider"),
		// 				Operation: to.Ptr("Register subscription for Microsoft.AzureFleet"),
		// 				Provider: to.Ptr("Microsoft Azure Fleet"),
		// 				Resource: to.Ptr("Subscription"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AzureFleet/unregister/action"),
		// 			Display: &armcomputefleet.OperationDisplay{
		// 				Description: to.Ptr("Unregisters Subscription with Microsoft.AzureFleet resource provider"),
		// 				Operation: to.Ptr("Unregister Subscription for Microsoft.AzureFleet"),
		// 				Provider: to.Ptr("Microsoft Azure Fleet"),
		// 				Resource: to.Ptr("Subscription"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr(armcomputefleet.OriginUserSystem),
		// 	}},
		// }
	}
}
