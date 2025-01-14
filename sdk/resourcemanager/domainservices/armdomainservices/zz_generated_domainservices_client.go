//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DomainServicesClient contains the methods for the DomainServices group.
// Don't use this type directly, use NewDomainServicesClient() instead.
type DomainServicesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDomainServicesClient creates a new instance of DomainServicesClient with the specified values.
func NewDomainServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DomainServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DomainServicesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - The Create Domain Service operation creates a new domain service with the specified parameters. If the specific service already
// exists, then any patchable properties will be updated and any immutable
// properties will remain unchanged.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *DomainServicesBeginCreateOrUpdateOptions) (DomainServicesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return DomainServicesCreateOrUpdatePollerResponse{}, err
	}
	result := DomainServicesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainServicesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DomainServicesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DomainServicesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The Create Domain Service operation creates a new domain service with the specified parameters. If the specific service already exists,
// then any patchable properties will be updated and any immutable
// properties will remain unchanged.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *DomainServicesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *DomainServicesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainService)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DomainServicesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - The Delete Domain Service operation deletes an existing Domain Service.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, domainServiceName string, options *DomainServicesBeginDeleteOptions) (DomainServicesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, domainServiceName, options)
	if err != nil {
		return DomainServicesDeletePollerResponse{}, err
	}
	result := DomainServicesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainServicesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DomainServicesDeletePollerResponse{}, err
	}
	result.Poller = &DomainServicesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The Delete Domain Service operation deletes an existing Domain Service.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, domainServiceName string, options *DomainServicesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DomainServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, options *DomainServicesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DomainServicesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - The Get Domain Service operation retrieves a json representation of the Domain Service.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) Get(ctx context.Context, resourceGroupName string, domainServiceName string, options *DomainServicesGetOptions) (DomainServicesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainServiceName, options)
	if err != nil {
		return DomainServicesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainServicesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainServicesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DomainServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, options *DomainServicesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainServicesClient) getHandleResponse(resp *http.Response) (DomainServicesGetResponse, error) {
	result := DomainServicesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainService); err != nil {
		return DomainServicesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DomainServicesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - The List Domain Services in Subscription operation lists all the domain services available under the given subscription (and across all resource
// groups within that subscription).
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) List(options *DomainServicesListOptions) *DomainServicesListPager {
	return &DomainServicesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DomainServicesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainServiceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DomainServicesClient) listCreateRequest(ctx context.Context, options *DomainServicesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AAD/domainServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DomainServicesClient) listHandleResponse(resp *http.Response) (DomainServicesListResponse, error) {
	result := DomainServicesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainServiceListResult); err != nil {
		return DomainServicesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DomainServicesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - The List Domain Services in Resource Group operation lists all the domain services available under the given resource group.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) ListByResourceGroup(resourceGroupName string, options *DomainServicesListByResourceGroupOptions) *DomainServicesListByResourceGroupPager {
	return &DomainServicesListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DomainServicesListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainServiceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DomainServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DomainServicesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DomainServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (DomainServicesListByResourceGroupResponse, error) {
	result := DomainServicesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainServiceListResult); err != nil {
		return DomainServicesListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DomainServicesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - The Update Domain Service operation can be used to update the existing deployment. The update call only supports the properties listed
// in the PATCH body.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *DomainServicesBeginUpdateOptions) (DomainServicesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return DomainServicesUpdatePollerResponse{}, err
	}
	result := DomainServicesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainServicesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DomainServicesUpdatePollerResponse{}, err
	}
	result.Poller = &DomainServicesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The Update Domain Service operation can be used to update the existing deployment. The update call only supports the properties listed in the
// PATCH body.
// If the operation fails it returns the *CloudError error type.
func (client *DomainServicesClient) update(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *DomainServicesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DomainServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *DomainServicesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainService)
}

// updateHandleError handles the Update error response.
func (client *DomainServicesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
