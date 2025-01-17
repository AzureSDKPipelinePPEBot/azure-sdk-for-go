//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// DataFlowDebugSessionClient contains the methods for the DataFlowDebugSession group.
// Don't use this type directly, use NewDataFlowDebugSessionClient() instead.
type DataFlowDebugSessionClient struct {
	con *Connection
}

// NewDataFlowDebugSessionClient creates a new instance of DataFlowDebugSessionClient with the specified values.
func NewDataFlowDebugSessionClient(con *Connection) *DataFlowDebugSessionClient {
	return &DataFlowDebugSessionClient{con: con}
}

// AddDataFlow - Add a data flow into debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) AddDataFlow(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (AddDataFlowToDebugSessionResponseResponse, error) {
	req, err := client.addDataFlowCreateRequest(ctx, request, options)
	if err != nil {
		return AddDataFlowToDebugSessionResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AddDataFlowToDebugSessionResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AddDataFlowToDebugSessionResponseResponse{}, client.addDataFlowHandleError(resp)
	}
	return client.addDataFlowHandleResponse(resp)
}

// addDataFlowCreateRequest creates the AddDataFlow request.
func (client *DataFlowDebugSessionClient) addDataFlowCreateRequest(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/addDataFlowToDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// addDataFlowHandleResponse handles the AddDataFlow response.
func (client *DataFlowDebugSessionClient) addDataFlowHandleResponse(resp *azcore.Response) (AddDataFlowToDebugSessionResponseResponse, error) {
	var val *AddDataFlowToDebugSessionResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AddDataFlowToDebugSessionResponseResponse{}, err
	}
	return AddDataFlowToDebugSessionResponseResponse{RawResponse: resp.Response, AddDataFlowToDebugSessionResponse: val}, nil
}

// addDataFlowHandleError handles the AddDataFlow error response.
func (client *DataFlowDebugSessionClient) addDataFlowHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginCreateDataFlowDebugSession - Creates a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) BeginCreateDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (CreateDataFlowDebugSessionResponsePollerResponse, error) {
	resp, err := client.createDataFlowDebugSession(ctx, request, options)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	result := CreateDataFlowDebugSessionResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("DataFlowDebugSessionClient.CreateDataFlowDebugSession", resp, client.con.Pipeline(), client.createDataFlowDebugSessionHandleError)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	poller := &createDataFlowDebugSessionResponsePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (CreateDataFlowDebugSessionResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateDataFlowDebugSession creates a new CreateDataFlowDebugSessionResponsePoller from the specified resume token.
// token - The value must come from a previous call to CreateDataFlowDebugSessionResponsePoller.ResumeToken().
func (client *DataFlowDebugSessionClient) ResumeCreateDataFlowDebugSession(ctx context.Context, token string) (CreateDataFlowDebugSessionResponsePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("DataFlowDebugSessionClient.CreateDataFlowDebugSession", token, client.con.Pipeline(), client.createDataFlowDebugSessionHandleError)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	poller := &createDataFlowDebugSessionResponsePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return CreateDataFlowDebugSessionResponsePollerResponse{}, err
	}
	result := CreateDataFlowDebugSessionResponsePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (CreateDataFlowDebugSessionResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateDataFlowDebugSession - Creates a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) createDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (*azcore.Response, error) {
	req, err := client.createDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createDataFlowDebugSessionHandleError(resp)
	}
	return resp, nil
}

// createDataFlowDebugSessionCreateRequest creates the CreateDataFlowDebugSession request.
func (client *DataFlowDebugSessionClient) createDataFlowDebugSessionCreateRequest(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (*azcore.Request, error) {
	urlPath := "/createDataFlowDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// createDataFlowDebugSessionHandleError handles the CreateDataFlowDebugSession error response.
func (client *DataFlowDebugSessionClient) createDataFlowDebugSessionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// DeleteDataFlowDebugSession - Deletes a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) DeleteDataFlowDebugSession(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*http.Response, error) {
	req, err := client.deleteDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteDataFlowDebugSessionHandleError(resp)
	}
	return resp.Response, nil
}

// deleteDataFlowDebugSessionCreateRequest creates the DeleteDataFlowDebugSession request.
func (client *DataFlowDebugSessionClient) deleteDataFlowDebugSessionCreateRequest(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*azcore.Request, error) {
	urlPath := "/deleteDataFlowDebugSession"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// deleteDataFlowDebugSessionHandleError handles the DeleteDataFlowDebugSession error response.
func (client *DataFlowDebugSessionClient) deleteDataFlowDebugSessionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) BeginExecuteCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (DataFlowDebugCommandResponsePollerResponse, error) {
	resp, err := client.executeCommand(ctx, request, options)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	result := DataFlowDebugCommandResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("DataFlowDebugSessionClient.ExecuteCommand", resp, client.con.Pipeline(), client.executeCommandHandleError)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	poller := &dataFlowDebugCommandResponsePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowDebugCommandResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeExecuteCommand creates a new DataFlowDebugCommandResponsePoller from the specified resume token.
// token - The value must come from a previous call to DataFlowDebugCommandResponsePoller.ResumeToken().
func (client *DataFlowDebugSessionClient) ResumeExecuteCommand(ctx context.Context, token string) (DataFlowDebugCommandResponsePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("DataFlowDebugSessionClient.ExecuteCommand", token, client.con.Pipeline(), client.executeCommandHandleError)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	poller := &dataFlowDebugCommandResponsePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return DataFlowDebugCommandResponsePollerResponse{}, err
	}
	result := DataFlowDebugCommandResponsePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (DataFlowDebugCommandResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) executeCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*azcore.Response, error) {
	req, err := client.executeCommandCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.executeCommandHandleError(resp)
	}
	return resp, nil
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *DataFlowDebugSessionClient) executeCommandCreateRequest(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*azcore.Request, error) {
	urlPath := "/executeDataFlowDebugCommand"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// executeCommandHandleError handles the ExecuteCommand error response.
func (client *DataFlowDebugSessionClient) executeCommandHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// QueryDataFlowDebugSessionsByWorkspace - Query all active data flow debug sessions.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspace(options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) QueryDataFlowDebugSessionsResponsePager {
	return &queryDataFlowDebugSessionsResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.queryDataFlowDebugSessionsByWorkspaceHandleResponse,
		errorer:   client.queryDataFlowDebugSessionsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp QueryDataFlowDebugSessionsResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.QueryDataFlowDebugSessionsResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// queryDataFlowDebugSessionsByWorkspaceCreateRequest creates the QueryDataFlowDebugSessionsByWorkspace request.
func (client *DataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/queryDataFlowDebugSessions"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleResponse handles the QueryDataFlowDebugSessionsByWorkspace response.
func (client *DataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp *azcore.Response) (QueryDataFlowDebugSessionsResponseResponse, error) {
	var val *QueryDataFlowDebugSessionsResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return QueryDataFlowDebugSessionsResponseResponse{}, err
	}
	return QueryDataFlowDebugSessionsResponseResponse{RawResponse: resp.Response, QueryDataFlowDebugSessionsResponse: val}, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleError handles the QueryDataFlowDebugSessionsByWorkspace error response.
func (client *DataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
