//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchangeanalysis

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ChangesListChangesByResourceGroupPager provides operations for iterating over paged responses.
type ChangesListChangesByResourceGroupPager struct {
	client    *ChangesClient
	current   ChangesListChangesByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ChangesListChangesByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ChangesListChangesByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ChangesListChangesByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ChangeList.NextLink == nil || len(*p.current.ChangeList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listChangesByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listChangesByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ChangesListChangesByResourceGroupResponse page.
func (p *ChangesListChangesByResourceGroupPager) PageResponse() ChangesListChangesByResourceGroupResponse {
	return p.current
}

// ChangesListChangesBySubscriptionPager provides operations for iterating over paged responses.
type ChangesListChangesBySubscriptionPager struct {
	client    *ChangesClient
	current   ChangesListChangesBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ChangesListChangesBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ChangesListChangesBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ChangesListChangesBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ChangeList.NextLink == nil || len(*p.current.ChangeList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listChangesBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listChangesBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ChangesListChangesBySubscriptionResponse page.
func (p *ChangesListChangesBySubscriptionPager) PageResponse() ChangesListChangesBySubscriptionResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceProviderOperationList.NextLink == nil || len(*p.current.ResourceProviderOperationList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// ResourceChangesListPager provides operations for iterating over paged responses.
type ResourceChangesListPager struct {
	client    *ResourceChangesClient
	current   ResourceChangesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ResourceChangesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ResourceChangesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ResourceChangesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ChangeList.NextLink == nil || len(*p.current.ChangeList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ResourceChangesListResponse page.
func (p *ResourceChangesListPager) PageResponse() ResourceChangesListResponse {
	return p.current
}
