// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure_test

import (
	"armdevopsinfrastructure"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
)

// Generated from example definition: 2024-04-04-preview/SubscriptionUsages_Usages.json
func ExampleSubscriptionUsagesClient_NewUsagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevopsinfrastructure.NewClientFactory("a2e95d27-c161-4b61-bda4-11512c14c2c2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionUsagesClient().NewUsagesPager("eastus", nil)
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
		// page = armdevopsinfrastructure.SubscriptionUsagesClientUsagesResponse{
		// 	PagedQuota: armdevopsinfrastructure.PagedQuota{
		// 		Value: []*armdevopsinfrastructure.Quota{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.DevOpsInfrastructure/Usages/standardDADSv5Family"),
		// 				Unit: to.Ptr("Count"),
		// 				CurrentValue: to.Ptr[int64](0),
		// 				Limit: to.Ptr[int64](212),
		// 				Name: &armdevopsinfrastructure.QuotaName{
		// 					Value: to.Ptr("standardDADSv5Family"),
		// 					LocalizedValue: to.Ptr("Standard DADSv5 Family vCPUs (PME VMSS)"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus/providers/Microsoft.DevOpsInfrastructure/Usages/standardDPLDSv5Family"),
		// 				Unit: to.Ptr("Count"),
		// 				CurrentValue: to.Ptr[int64](0),
		// 				Limit: to.Ptr[int64](100),
		// 				Name: &armdevopsinfrastructure.QuotaName{
		// 					Value: to.Ptr("standardDPLDSv5Family"),
		// 					LocalizedValue: to.Ptr("Standard DPLDSv5 Family vCPUs (PME VMSS)"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
