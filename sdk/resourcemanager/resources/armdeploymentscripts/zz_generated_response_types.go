//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentscripts

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// DeploymentScriptsCreatePollerResponse contains the response from method DeploymentScripts.Create.
type DeploymentScriptsCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DeploymentScriptsCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DeploymentScriptsCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DeploymentScriptsCreateResponse, error) {
	respType := DeploymentScriptsCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DeploymentScriptClassification)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DeploymentScriptsCreatePollerResponse from the provided client and resume token.
func (l *DeploymentScriptsCreatePollerResponse) Resume(ctx context.Context, client *DeploymentScriptsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DeploymentScriptsClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &DeploymentScriptsCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DeploymentScriptsCreateResponse contains the response from method DeploymentScripts.Create.
type DeploymentScriptsCreateResponse struct {
	DeploymentScriptsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsCreateResult contains the result from method DeploymentScripts.Create.
type DeploymentScriptsCreateResult struct {
	DeploymentScriptClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentScriptsCreateResult.
func (d *DeploymentScriptsCreateResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDeploymentScriptClassification(data)
	if err != nil {
		return err
	}
	d.DeploymentScriptClassification = res
	return nil
}

// DeploymentScriptsDeleteResponse contains the response from method DeploymentScripts.Delete.
type DeploymentScriptsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsGetLogsDefaultResponse contains the response from method DeploymentScripts.GetLogsDefault.
type DeploymentScriptsGetLogsDefaultResponse struct {
	DeploymentScriptsGetLogsDefaultResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsGetLogsDefaultResult contains the result from method DeploymentScripts.GetLogsDefault.
type DeploymentScriptsGetLogsDefaultResult struct {
	ScriptLog
}

// DeploymentScriptsGetLogsResponse contains the response from method DeploymentScripts.GetLogs.
type DeploymentScriptsGetLogsResponse struct {
	DeploymentScriptsGetLogsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsGetLogsResult contains the result from method DeploymentScripts.GetLogs.
type DeploymentScriptsGetLogsResult struct {
	ScriptLogsList
}

// DeploymentScriptsGetResponse contains the response from method DeploymentScripts.Get.
type DeploymentScriptsGetResponse struct {
	DeploymentScriptsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsGetResult contains the result from method DeploymentScripts.Get.
type DeploymentScriptsGetResult struct {
	DeploymentScriptClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentScriptsGetResult.
func (d *DeploymentScriptsGetResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDeploymentScriptClassification(data)
	if err != nil {
		return err
	}
	d.DeploymentScriptClassification = res
	return nil
}

// DeploymentScriptsListByResourceGroupResponse contains the response from method DeploymentScripts.ListByResourceGroup.
type DeploymentScriptsListByResourceGroupResponse struct {
	DeploymentScriptsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsListByResourceGroupResult contains the result from method DeploymentScripts.ListByResourceGroup.
type DeploymentScriptsListByResourceGroupResult struct {
	DeploymentScriptListResult
}

// DeploymentScriptsListBySubscriptionResponse contains the response from method DeploymentScripts.ListBySubscription.
type DeploymentScriptsListBySubscriptionResponse struct {
	DeploymentScriptsListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsListBySubscriptionResult contains the result from method DeploymentScripts.ListBySubscription.
type DeploymentScriptsListBySubscriptionResult struct {
	DeploymentScriptListResult
}

// DeploymentScriptsUpdateResponse contains the response from method DeploymentScripts.Update.
type DeploymentScriptsUpdateResponse struct {
	DeploymentScriptsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DeploymentScriptsUpdateResult contains the result from method DeploymentScripts.Update.
type DeploymentScriptsUpdateResult struct {
	DeploymentScriptClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeploymentScriptsUpdateResult.
func (d *DeploymentScriptsUpdateResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDeploymentScriptClassification(data)
	if err != nil {
		return err
	}
	d.DeploymentScriptClassification = res
	return nil
}
