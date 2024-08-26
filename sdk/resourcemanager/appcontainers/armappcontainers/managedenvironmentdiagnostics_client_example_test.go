//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedEnvironmentDiagnostics_List.json
func ExampleManagedEnvironmentDiagnosticsClient_ListDetectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentDiagnosticsClient().ListDetectors(ctx, "mikono-workerapp-test-rg", "mikonokubeenv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticsCollection = armappcontainers.DiagnosticsCollection{
	// 	Value: []*armappcontainers.Diagnostics{
	// 		{
	// 			Name: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 			Type: to.Ptr("Microsoft.App/managedEnvironments/detectors"),
	// 			ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/managedEnvironments/mikonokubeenv/detectors/ManagedEnvAvailabilityMetrics"),
	// 			Properties: &armappcontainers.DiagnosticsProperties{
	// 				Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
	// 				},
	// 				Metadata: &armappcontainers.DiagnosticsDefinition{
	// 					Name: to.Ptr("Availability Metrics for Managed Environments"),
	// 					Type: to.Ptr("Analysis"),
	// 					Author: to.Ptr(""),
	// 					Category: to.Ptr("Availability and Performance"),
	// 					ID: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 					Score: to.Ptr[float32](0),
	// 					SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
	// 					},
	// 				},
	// 				Status: &armappcontainers.DiagnosticsStatus{
	// 					StatusID: to.Ptr[int32](4),
	// 				},
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edf14cc0a577f6b9c4e3ce018cec0c383e64b7b0/specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedEnvironmentDiagnostics_Get.json
func ExampleManagedEnvironmentDiagnosticsClient_GetDetector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedEnvironmentDiagnosticsClient().GetDetector(ctx, "mikono-workerapp-test-rg", "mikonokubeenv", "ManagedEnvAvailabilityMetrics", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Diagnostics = armappcontainers.Diagnostics{
	// 	Name: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments/detectors"),
	// 	ID: to.Ptr("/subscriptions/f07f3711-b45e-40fe-a941-4e6d93f851e6/resourceGroups/mikono-workerapp-test-rg/providers/Microsoft.App/managedEnvironments/mikonokubeenv/detectors/ManagedEnvAvailabilityMetrics"),
	// 	Properties: &armappcontainers.DiagnosticsProperties{
	// 		Dataset: []*armappcontainers.DiagnosticsDataAPIResponse{
	// 			{
	// 				RenderingProperties: &armappcontainers.DiagnosticRendering{
	// 					Type: to.Ptr[int32](8),
	// 					Description: to.Ptr(""),
	// 					IsVisible: to.Ptr(true),
	// 					Title: to.Ptr("Managed Environment Network Inbound "),
	// 				},
	// 				Table: &armappcontainers.DiagnosticDataTableResponseObject{
	// 					Columns: []*armappcontainers.DiagnosticDataTableResponseColumn{
	// 						{
	// 							ColumnName: to.Ptr("TimeStamp"),
	// 							DataType: to.Ptr("DateTime"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("Metric"),
	// 							DataType: to.Ptr("String"),
	// 						},
	// 						{
	// 							ColumnName: to.Ptr("Average"),
	// 							DataType: to.Ptr("Double"),
	// 					}},
	// 					Rows: []any{
	// 						[]any{
	// 							"2022-03-15T21:35:00",
	// 							"RxBytes",
	// 							float64(0),
	// 					}},
	// 					TableName: to.Ptr(""),
	// 				},
	// 		}},
	// 		Metadata: &armappcontainers.DiagnosticsDefinition{
	// 			Name: to.Ptr("Managed Env Netowrk Inbound and Outbound"),
	// 			Type: to.Ptr("Detector"),
	// 			Description: to.Ptr("This detector shows the Managed Environment Network Inbound and Outbound."),
	// 			Author: to.Ptr(""),
	// 			Category: to.Ptr("Availability and Performance"),
	// 			ID: to.Ptr("ManagedEnvAvailabilityMetrics"),
	// 			Score: to.Ptr[float32](0),
	// 			SupportTopicList: []*armappcontainers.DiagnosticSupportTopic{
	// 			},
	// 		},
	// 		Status: &armappcontainers.DiagnosticsStatus{
	// 			StatusID: to.Ptr[int32](3),
	// 		},
	// 	},
	// }
}
