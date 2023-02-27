package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
)

// ApplicationGatewayPrivateLinkResourcesClient is the network Client
type ApplicationGatewayPrivateLinkResourcesClient struct {
	BaseClient
}

// NewApplicationGatewayPrivateLinkResourcesClient creates an instance of the
// ApplicationGatewayPrivateLinkResourcesClient client.
func NewApplicationGatewayPrivateLinkResourcesClient(subscriptionID string) ApplicationGatewayPrivateLinkResourcesClient {
	return NewApplicationGatewayPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewayPrivateLinkResourcesClientWithBaseURI creates an instance of the
// ApplicationGatewayPrivateLinkResourcesClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewApplicationGatewayPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewayPrivateLinkResourcesClient {
	return ApplicationGatewayPrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all private link resources on an application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
func (client ApplicationGatewayPrivateLinkResourcesClient) List(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewayPrivateLinkResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateLinkResourcesClient.List")
		defer func() {
			sc := -1
			if result.agplrlr.Response.Response != nil {
				sc = result.agplrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateLinkResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.agplrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateLinkResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result.agplrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateLinkResourcesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.agplrlr.hasNextLink() && result.agplrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationGatewayPrivateLinkResourcesClient) ListPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewayPrivateLinkResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationGatewayPrivateLinkResourcesClient) ListResponder(resp *http.Response) (result ApplicationGatewayPrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ApplicationGatewayPrivateLinkResourcesClient) listNextResults(ctx context.Context, lastResults ApplicationGatewayPrivateLinkResourceListResult) (result ApplicationGatewayPrivateLinkResourceListResult, err error) {
	req, err := lastResults.applicationGatewayPrivateLinkResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateLinkResourcesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateLinkResourcesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayPrivateLinkResourcesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGatewayPrivateLinkResourcesClient) ListComplete(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewayPrivateLinkResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayPrivateLinkResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, applicationGatewayName)
	return
}
