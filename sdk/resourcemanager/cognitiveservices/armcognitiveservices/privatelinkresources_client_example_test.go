//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2023-10-01-preview/examples/ListPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().List(ctx, "res6977", "sto2527", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceListResult = armcognitiveservices.PrivateLinkResourceListResult{
	// 	Value: []*armcognitiveservices.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("blob"),
	// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6977/providers/Microsoft.CognitiveServices/accounts/sto2527/privateLinkResources/account"),
	// 			Properties: &armcognitiveservices.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("account"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("default")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.cognitiveservices.azure.com")},
	// 					},
	// 			}},
	// 		}
}
