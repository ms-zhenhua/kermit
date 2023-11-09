package synapse

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// PrivateEndpointConnectionsPrivateLinkHubClient is the azure Synapse Analytics Management Client
type PrivateEndpointConnectionsPrivateLinkHubClient struct {
    BaseClient
}
// NewPrivateEndpointConnectionsPrivateLinkHubClient creates an instance of the
// PrivateEndpointConnectionsPrivateLinkHubClient client.
func NewPrivateEndpointConnectionsPrivateLinkHubClient(subscriptionID string) PrivateEndpointConnectionsPrivateLinkHubClient {
    return NewPrivateEndpointConnectionsPrivateLinkHubClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionsPrivateLinkHubClientWithBaseURI creates an instance of the
// PrivateEndpointConnectionsPrivateLinkHubClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewPrivateEndpointConnectionsPrivateLinkHubClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsPrivateLinkHubClient {
        return PrivateEndpointConnectionsPrivateLinkHubClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get all PrivateEndpointConnection in the PrivateLinkHub by name
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // privateLinkHubName - name of the privateLinkHub
        // privateEndpointConnectionName - name of the privateEndpointConnection
func (client PrivateEndpointConnectionsPrivateLinkHubClient) Get(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string) (result PrivateEndpointConnectionForPrivateLinkHub, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateEndpointConnectionsPrivateLinkHubClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, privateLinkHubName, privateEndpointConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client PrivateEndpointConnectionsPrivateLinkHubClient) GetPreparer(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "privateEndpointConnectionName": autorest.Encode("path",privateEndpointConnectionName),
        "privateLinkHubName": autorest.Encode("path",privateLinkHubName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2021-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections/{privateEndpointConnectionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateEndpointConnectionsPrivateLinkHubClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client PrivateEndpointConnectionsPrivateLinkHubClient) GetResponder(resp *http.Response) (result PrivateEndpointConnectionForPrivateLinkHub, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List get all PrivateEndpointConnections in the PrivateLinkHub
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // privateLinkHubName - name of the privateLinkHub
func (client PrivateEndpointConnectionsPrivateLinkHubClient) List(ctx context.Context, resourceGroupName string, privateLinkHubName string) (result PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateEndpointConnectionsPrivateLinkHubClient.List")
        defer func() {
            sc := -1
        if result.pecfplhrcr.Response.Response != nil {
        sc = result.pecfplhrcr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, privateLinkHubName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.pecfplhrcr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "List", resp, "Failure sending request")
        return
        }

        result.pecfplhrcr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "List", resp, "Failure responding to request")
        return
        }
            if result.pecfplhrcr.hasNextLink() && result.pecfplhrcr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client PrivateEndpointConnectionsPrivateLinkHubClient) ListPreparer(ctx context.Context, resourceGroupName string, privateLinkHubName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "privateLinkHubName": autorest.Encode("path",privateLinkHubName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2021-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateEndpointConnectionsPrivateLinkHubClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client PrivateEndpointConnectionsPrivateLinkHubClient) ListResponder(resp *http.Response) (result PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client PrivateEndpointConnectionsPrivateLinkHubClient) listNextResults(ctx context.Context, lastResults PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponse) (result PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponse, err error) {
            req, err := lastResults.privateEndpointConnectionForPrivateLinkHubResourceCollectionResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.PrivateEndpointConnectionsPrivateLinkHubClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client PrivateEndpointConnectionsPrivateLinkHubClient) ListComplete(ctx context.Context, resourceGroupName string, privateLinkHubName string) (result PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponseIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PrivateEndpointConnectionsPrivateLinkHubClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, privateLinkHubName)
                            return
            }

