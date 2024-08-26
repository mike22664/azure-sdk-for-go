//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListCertificates.json
func ExampleCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificatesClient().NewListPager(&armappservice.CertificatesClientListOptions{Filter: nil})
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
		// page.AppCertificateCollection = armappservice.AppCertificateCollection{
		// 	Value: []*armappservice.AppCertificate{
		// 		{
		// 			Name: to.Ptr("testc6282"),
		// 			Type: to.Ptr("Microsoft.Web/certificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc6282"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.AppCertificateProperties{
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
		// 				FriendlyName: to.Ptr(""),
		// 				HostNames: []*string{
		// 					to.Ptr("ServerCert")},
		// 					IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
		// 					Issuer: to.Ptr("CACert"),
		// 					SubjectName: to.Ptr("ServerCert"),
		// 					Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testc4912"),
		// 				Type: to.Ptr("Microsoft.Web/certificates"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc4912"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armappservice.AppCertificateProperties{
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2040-12-31T23:59:59.000Z"); return t}()),
		// 					FriendlyName: to.Ptr(""),
		// 					HostNames: []*string{
		// 						to.Ptr("ServerCert2")},
		// 						IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-12-12T23:40:25.000Z"); return t}()),
		// 						Issuer: to.Ptr("CACert"),
		// 						SubjectName: to.Ptr("ServerCert2"),
		// 						Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B0490D5886EBFE"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListCertificatesByResourceGroup.json
func ExampleCertificatesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCertificatesClient().NewListByResourceGroupPager("testrg123", nil)
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
		// page.AppCertificateCollection = armappservice.AppCertificateCollection{
		// 	Value: []*armappservice.AppCertificate{
		// 		{
		// 			Name: to.Ptr("testc6282"),
		// 			Type: to.Ptr("Microsoft.Web/certificates"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc6282"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.AppCertificateProperties{
		// 				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
		// 				FriendlyName: to.Ptr(""),
		// 				HostNames: []*string{
		// 					to.Ptr("ServerCert")},
		// 					IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
		// 					Issuer: to.Ptr("CACert"),
		// 					SubjectName: to.Ptr("ServerCert"),
		// 					Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testc4912"),
		// 				Type: to.Ptr("Microsoft.Web/certificates"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc4912"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armappservice.AppCertificateProperties{
		// 					ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2040-12-31T23:59:59.000Z"); return t}()),
		// 					FriendlyName: to.Ptr(""),
		// 					HostNames: []*string{
		// 						to.Ptr("ServerCert2")},
		// 						IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-12-12T23:40:25.000Z"); return t}()),
		// 						Issuer: to.Ptr("CACert"),
		// 						SubjectName: to.Ptr("ServerCert2"),
		// 						Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B0490D5886EBFE"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetCertificate.json
func ExampleCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().Get(ctx, "testrg123", "testc6282", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppCertificate = armappservice.AppCertificate{
	// 	Name: to.Ptr("testc6282"),
	// 	Type: to.Ptr("Microsoft.Web/certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc6282"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.AppCertificateProperties{
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		FriendlyName: to.Ptr(""),
	// 		HostNames: []*string{
	// 			to.Ptr("ServerCert")},
	// 			IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
	// 			Issuer: to.Ptr("CACert"),
	// 			SubjectName: to.Ptr("ServerCert"),
	// 			Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/CreateOrUpdateCertificate.json
func ExampleCertificatesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().CreateOrUpdate(ctx, "testrg123", "testc6282", armappservice.AppCertificate{
		Location: to.Ptr("East US"),
		Properties: &armappservice.AppCertificateProperties{
			HostNames: []*string{
				to.Ptr("ServerCert")},
			Password: to.Ptr("<password>"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppCertificate = armappservice.AppCertificate{
	// 	Name: to.Ptr("testc6282"),
	// 	Type: to.Ptr("Microsoft.Web/certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc6282"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.AppCertificateProperties{
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		FriendlyName: to.Ptr(""),
	// 		HostNames: []*string{
	// 			to.Ptr("ServerCert")},
	// 			IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
	// 			Issuer: to.Ptr("CACert"),
	// 			SubjectName: to.Ptr("ServerCert"),
	// 			Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/DeleteCertificate.json
func ExampleCertificatesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCertificatesClient().Delete(ctx, "testrg123", "testc6282", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/PatchCertificate.json
func ExampleCertificatesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificatesClient().Update(ctx, "testrg123", "testc6282", armappservice.AppCertificatePatchResource{
		Properties: &armappservice.AppCertificatePatchResourceProperties{
			Password: to.Ptr("<password>"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AppCertificate = armappservice.AppCertificate{
	// 	Name: to.Ptr("testc6282"),
	// 	Type: to.Ptr("Microsoft.Web/certificates"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/certificates/testc6282"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappservice.AppCertificateProperties{
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2039-12-31T23:59:59.000Z"); return t}()),
	// 		FriendlyName: to.Ptr(""),
	// 		HostNames: []*string{
	// 			to.Ptr("ServerCert")},
	// 			IssueDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-11-12T23:40:25.000Z"); return t}()),
	// 			Issuer: to.Ptr("CACert"),
	// 			SubjectName: to.Ptr("ServerCert"),
	// 			Thumbprint: to.Ptr("FE703D7411A44163B6D32B3AD9B03E175886EBFE"),
	// 		},
	// 	}
}
