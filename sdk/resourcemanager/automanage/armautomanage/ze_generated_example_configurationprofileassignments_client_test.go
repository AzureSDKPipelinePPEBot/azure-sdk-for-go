//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/createOrUpdateConfigurationProfileAssignment.json
func ExampleConfigurationProfileAssignmentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfileAssignmentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<configuration-profile-assignment-name>",
		"<resource-group-name>",
		"<vm-name>",
		armautomanage.ConfigurationProfileAssignment{
			Properties: &armautomanage.ConfigurationProfileAssignmentProperties{
				AccountID:                        to.StringPtr("<account-id>"),
				ConfigurationProfile:             armautomanage.ConfigurationProfileAzureVirtualMachineBestPracticesProduction.ToPtr(),
				ConfigurationProfilePreferenceID: to.StringPtr("<configuration-profile-preference-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationProfileAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/getConfigurationProfileAssignment.json
func ExampleConfigurationProfileAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfileAssignmentsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<configuration-profile-assignment-name>",
		"<vm-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationProfileAssignment.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/deleteConfigurationProfileAssignment.json
func ExampleConfigurationProfileAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfileAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<configuration-profile-assignment-name>",
		"<vm-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/listConfigurationProfileAssignments.json
func ExampleConfigurationProfileAssignmentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfileAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.List(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/listConfigurationProfileAssignmentsBySubscription.json
func ExampleConfigurationProfileAssignmentsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfileAssignmentsClient("<subscription-id>", cred, nil)
	_, err = client.ListBySubscription(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
