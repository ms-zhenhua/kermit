package appplatform

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

// StoragesClient is the REST API for Azure Spring Apps
type StoragesClient struct {
    BaseClient
}
// NewStoragesClient creates an instance of the StoragesClient client.
func NewStoragesClient(subscriptionID string) StoragesClient {
    return NewStoragesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStoragesClientWithBaseURI creates an instance of the StoragesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewStoragesClientWithBaseURI(baseURI string, subscriptionID string) StoragesClient {
        return StoragesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update storage resource.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // storageName - the name of the storage resource.
        // storageResource - parameters for the create or update operation
func (client StoragesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, storageName string, storageResource StorageResource) (result StoragesCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StoragesClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, storageName, storageResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client StoragesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, storageName string, storageResource StorageResource) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "storageName": autorest.Encode("path",storageName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/storages/{storageName}",pathParameters),
autorest.WithJSON(storageResource),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client StoragesClient) CreateOrUpdateSender(req *http.Request) (future StoragesCreateOrUpdateFuture, err error) {
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
    func (client StoragesClient) CreateOrUpdateResponder(resp *http.Response) (result StorageResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete the storage resource.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // storageName - the name of the storage resource.
func (client StoragesClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, storageName string) (result StoragesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StoragesClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, storageName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client StoragesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, storageName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "storageName": autorest.Encode("path",storageName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/storages/{storageName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client StoragesClient) DeleteSender(req *http.Request) (future StoragesDeleteFuture, err error) {
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
    func (client StoragesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get the storage resource.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // storageName - the name of the storage resource.
func (client StoragesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, storageName string) (result StorageResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StoragesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, storageName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client StoragesClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, storageName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "storageName": autorest.Encode("path",storageName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/storages/{storageName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client StoragesClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client StoragesClient) GetResponder(resp *http.Response) (result StorageResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list all the storages of one Azure Spring Apps resource.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
func (client StoragesClient) List(ctx context.Context, resourceGroupName string, serviceName string) (result StorageResourceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StoragesClient.List")
        defer func() {
            sc := -1
        if result.src.Response.Response != nil {
        sc = result.src.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.src.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "List", resp, "Failure sending request")
        return
        }

        result.src, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "List", resp, "Failure responding to request")
        return
        }
            if result.src.hasNextLink() && result.src.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client StoragesClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/storages",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client StoragesClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client StoragesClient) ListResponder(resp *http.Response) (result StorageResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client StoragesClient) listNextResults(ctx context.Context, lastResults StorageResourceCollection) (result StorageResourceCollection, err error) {
            req, err := lastResults.storageResourceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "appplatform.StoragesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "appplatform.StoragesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "appplatform.StoragesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client StoragesClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result StorageResourceCollectionIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/StoragesClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, serviceName)
                            return
            }

