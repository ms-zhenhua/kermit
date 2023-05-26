package network

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
)

// DscpConfigurationClient is the network Client
type DscpConfigurationClient struct {
    BaseClient
}
// NewDscpConfigurationClient creates an instance of the DscpConfigurationClient client.
func NewDscpConfigurationClient(subscriptionID string) DscpConfigurationClient {
    return NewDscpConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDscpConfigurationClientWithBaseURI creates an instance of the DscpConfigurationClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewDscpConfigurationClientWithBaseURI(baseURI string, subscriptionID string) DscpConfigurationClient {
        return DscpConfigurationClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a DSCP Configuration.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // dscpConfigurationName - the name of the resource.
        // parameters - parameters supplied to the create or update dscp configuration operation.
func (client DscpConfigurationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration) (result DscpConfigurationCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, dscpConfigurationName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client DscpConfigurationClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, dscpConfigurationName string, parameters DscpConfiguration) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "dscpConfigurationName": autorest.Encode("path",dscpConfigurationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            parameters.Etag = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscpConfigurationClient) CreateOrUpdateSender(req *http.Request) (future DscpConfigurationCreateOrUpdateFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client DscpConfigurationClient) CreateOrUpdateResponder(resp *http.Response) (result DscpConfiguration, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes a DSCP Configuration.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // dscpConfigurationName - the name of the resource.
func (client DscpConfigurationClient) Delete(ctx context.Context, resourceGroupName string, dscpConfigurationName string) (result DscpConfigurationDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, dscpConfigurationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client DscpConfigurationClient) DeletePreparer(ctx context.Context, resourceGroupName string, dscpConfigurationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "dscpConfigurationName": autorest.Encode("path",dscpConfigurationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscpConfigurationClient) DeleteSender(req *http.Request) (future DscpConfigurationDeleteFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client DscpConfigurationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets a DSCP Configuration.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // dscpConfigurationName - the name of the resource.
func (client DscpConfigurationClient) Get(ctx context.Context, resourceGroupName string, dscpConfigurationName string) (result DscpConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, dscpConfigurationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client DscpConfigurationClient) GetPreparer(ctx context.Context, resourceGroupName string, dscpConfigurationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "dscpConfigurationName": autorest.Encode("path",dscpConfigurationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations/{dscpConfigurationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscpConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client DscpConfigurationClient) GetResponder(resp *http.Response) (result DscpConfiguration, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List gets a DSCP Configuration.
    // Parameters:
        // resourceGroupName - the name of the resource group.
func (client DscpConfigurationClient) List(ctx context.Context, resourceGroupName string) (result DscpConfigurationListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.List")
        defer func() {
            sc := -1
        if result.dclr.Response.Response != nil {
        sc = result.dclr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.dclr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "List", resp, "Failure sending request")
        return
        }

        result.dclr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "List", resp, "Failure responding to request")
        return
        }
            if result.dclr.hasNextLink() && result.dclr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client DscpConfigurationClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dscpConfigurations",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscpConfigurationClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client DscpConfigurationClient) ListResponder(resp *http.Response) (result DscpConfigurationListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client DscpConfigurationClient) listNextResults(ctx context.Context, lastResults DscpConfigurationListResult) (result DscpConfigurationListResult, err error) {
            req, err := lastResults.dscpConfigurationListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client DscpConfigurationClient) ListComplete(ctx context.Context, resourceGroupName string) (result DscpConfigurationListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName)
                            return
            }

// ListAll gets all dscp configurations in a subscription.
func (client DscpConfigurationClient) ListAll(ctx context.Context) (result DscpConfigurationListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.ListAll")
        defer func() {
            sc := -1
        if result.dclr.Response.Response != nil {
        sc = result.dclr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listAllNextResults
    req, err := client.ListAllPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "ListAll", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListAllSender(req)
        if err != nil {
        result.dclr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "ListAll", resp, "Failure sending request")
        return
        }

        result.dclr, err = client.ListAllResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "ListAll", resp, "Failure responding to request")
        return
        }
            if result.dclr.hasNextLink() && result.dclr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListAllPreparer prepares the ListAll request.
    func (client DscpConfigurationClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/dscpConfigurations",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListAllSender sends the ListAll request. The method will close the
    // http.Response Body if it receives an error.
    func (client DscpConfigurationClient) ListAllSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListAllResponder handles the response to the ListAll request. The method always
    // closes the http.Response Body.
    func (client DscpConfigurationClient) ListAllResponder(resp *http.Response) (result DscpConfigurationListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listAllNextResults retrieves the next set of results, if any.
            func (client DscpConfigurationClient) listAllNextResults(ctx context.Context, lastResults DscpConfigurationListResult) (result DscpConfigurationListResult, err error) {
            req, err := lastResults.dscpConfigurationListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "listAllNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListAllSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "listAllNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListAllResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.DscpConfigurationClient", "listAllNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListAllComplete enumerates all values, automatically crossing page boundaries as required.
            func (client DscpConfigurationClient) ListAllComplete(ctx context.Context) (result DscpConfigurationListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/DscpConfigurationClient.ListAll")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListAll(ctx)
                            return
            }

