//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GetVMMServer.json
func ExampleVmmServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVmmServersClient().Get(ctx, "testrg", "ContosoVMMServer", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMMServer = armscvmm.VMMServer{
	// 	Name: to.Ptr("ContosoVMMServer"),
	// 	Type: to.Ptr("Microsoft.SCVMM/VMMServers"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 	Location: to.Ptr("East US"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VMMServerProperties{
	// 		ConnectionStatus: to.Ptr("Connected"),
	// 		Fqdn: to.Ptr("VMM.contoso.com"),
	// 		Port: to.Ptr[int32](1234),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		UUID: to.Ptr("fd3c3665-1729-4b7b-9a38-238e83b0f98b"),
	// 		Version: to.Ptr("2.0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/CreateVMMServer.json
func ExampleVmmServersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVmmServersClient().BeginCreateOrUpdate(ctx, "testrg", "ContosoVMMServer", armscvmm.VMMServer{
		Location: to.Ptr("East US"),
		ExtendedLocation: &armscvmm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
			Type: to.Ptr("customLocation"),
		},
		Properties: &armscvmm.VMMServerProperties{
			Credentials: &armscvmm.VMMCredential{
				Password: to.Ptr("password"),
				Username: to.Ptr("testuser"),
			},
			Fqdn: to.Ptr("VMM.contoso.com"),
			Port: to.Ptr[int32](1234),
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
	// res.VMMServer = armscvmm.VMMServer{
	// 	Name: to.Ptr("ContosoVMMServer"),
	// 	Type: to.Ptr("Microsoft.SCVMM/VMMServers"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 	Location: to.Ptr("East US"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VMMServerProperties{
	// 		ConnectionStatus: to.Ptr("Connected"),
	// 		Fqdn: to.Ptr("VMM.contoso.com"),
	// 		Port: to.Ptr[int32](1234),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		UUID: to.Ptr("fd3c3665-1729-4b7b-9a38-238e83b0f98b"),
	// 		Version: to.Ptr("2.0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteVMMServer.json
func ExampleVmmServersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVmmServersClient().BeginDelete(ctx, "testrg", "ContosoVMMServer", &armscvmm.VmmServersClientBeginDeleteOptions{Force: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/UpdateVMMServer.json
func ExampleVmmServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVmmServersClient().BeginUpdate(ctx, "testrg", "ContosoVMMServer", armscvmm.ResourcePatch{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.VMMServer = armscvmm.VMMServer{
	// 	Name: to.Ptr("ContosoVMMServer"),
	// 	Type: to.Ptr("Microsoft.SCVMM/VMMServers"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VMMServerProperties{
	// 		ConnectionStatus: to.Ptr("Connected"),
	// 		Fqdn: to.Ptr("VMM.contoso.com"),
	// 		Port: to.Ptr[int32](1234),
	// 		ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 		UUID: to.Ptr("fd3c3665-1729-4b7b-9a38-238e83b0f98b"),
	// 		Version: to.Ptr("2.0"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/ListVMMServersByResourceGroup.json
func ExampleVmmServersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVmmServersClient().NewListByResourceGroupPager("testrg", nil)
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
		// page.VMMServerListResult = armscvmm.VMMServerListResult{
		// 	Value: []*armscvmm.VMMServer{
		// 		{
		// 			Name: to.Ptr("ContosoVMMServer"),
		// 			Type: to.Ptr("Microsoft.SCVMM/VMMServers"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
		// 			Location: to.Ptr("East US"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.VMMServerProperties{
		// 				ConnectionStatus: to.Ptr("Connected"),
		// 				Fqdn: to.Ptr("VMM.contoso.com"),
		// 				Port: to.Ptr[int32](1234),
		// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 				UUID: to.Ptr("fd3c3665-1729-4b7b-9a38-238e83b0f98b"),
		// 				Version: to.Ptr("2.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/ListVMMServersBySubscription.json
func ExampleVmmServersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVmmServersClient().NewListBySubscriptionPager(nil)
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
		// page.VMMServerListResult = armscvmm.VMMServerListResult{
		// 	Value: []*armscvmm.VMMServer{
		// 		{
		// 			Name: to.Ptr("ContosoVMMServer"),
		// 			Type: to.Ptr("Microsoft.SCVMM/VMMServers"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer"),
		// 			Location: to.Ptr("East US"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.VMMServerProperties{
		// 				ConnectionStatus: to.Ptr("Connected"),
		// 				Fqdn: to.Ptr("VMM.contoso.com"),
		// 				Port: to.Ptr[int32](1234),
		// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 				UUID: to.Ptr("fd3c3665-1729-4b7b-9a38-238e83b0f98b"),
		// 				Version: to.Ptr("2.0"),
		// 			},
		// 	}},
		// }
	}
}
