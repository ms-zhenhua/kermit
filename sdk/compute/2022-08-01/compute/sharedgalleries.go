package compute

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

// SharedGalleriesClient is the compute Client
type SharedGalleriesClient struct {
    BaseClient
}
// NewSharedGalleriesClient creates an instance of the SharedGalleriesClient client.
func NewSharedGalleriesClient(subscriptionID string) SharedGalleriesClient {
    return NewSharedGalleriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharedGalleriesClientWithBaseURI creates an instance of the SharedGalleriesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewSharedGalleriesClientWithBaseURI(baseURI string, subscriptionID string) SharedGalleriesClient {
        return SharedGalleriesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get a shared gallery by subscription id or tenant id.
    // Parameters:
        // location - resource location.
        // galleryUniqueName - the unique name of the Shared Gallery.
func (client SharedGalleriesClient) Get(ctx context.Context, location string, galleryUniqueName string) (result SharedGallery, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SharedGalleriesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, location, galleryUniqueName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client SharedGalleriesClient) GetPreparer(ctx context.Context, location string, galleryUniqueName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "galleryUniqueName": autorest.Encode("path",galleryUniqueName),
        "location": autorest.Encode("path",location),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-01-03"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SharedGalleriesClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client SharedGalleriesClient) GetResponder(resp *http.Response) (result SharedGallery, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list shared galleries by subscription id or tenant id.
    // Parameters:
        // location - resource location.
        // sharedTo - the query parameter to decide what shared galleries to fetch when doing listing operations.
func (client SharedGalleriesClient) List(ctx context.Context, location string, sharedTo SharedToValues) (result SharedGalleryListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SharedGalleriesClient.List")
        defer func() {
            sc := -1
        if result.sgl.Response.Response != nil {
        sc = result.sgl.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, location, sharedTo)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.sgl.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "List", resp, "Failure sending request")
        return
        }

        result.sgl, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "List", resp, "Failure responding to request")
        return
        }
            if result.sgl.hasNextLink() && result.sgl.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client SharedGalleriesClient) ListPreparer(ctx context.Context, location string, sharedTo SharedToValues) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "location": autorest.Encode("path",location),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-01-03"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(string(sharedTo)) > 0 {
        queryParameters["sharedTo"] = autorest.Encode("query",sharedTo)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client SharedGalleriesClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client SharedGalleriesClient) ListResponder(resp *http.Response) (result SharedGalleryList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client SharedGalleriesClient) listNextResults(ctx context.Context, lastResults SharedGalleryList) (result SharedGalleryList, err error) {
            req, err := lastResults.sharedGalleryListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.SharedGalleriesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client SharedGalleriesClient) ListComplete(ctx context.Context, location string, sharedTo SharedToValues) (result SharedGalleryListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SharedGalleriesClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, location, sharedTo)
                            return
            }

