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

// ApplicationGatewayWafDynamicManifestsDefaultClient is the network Client
type ApplicationGatewayWafDynamicManifestsDefaultClient struct {
    BaseClient
}
// NewApplicationGatewayWafDynamicManifestsDefaultClient creates an instance of the
// ApplicationGatewayWafDynamicManifestsDefaultClient client.
func NewApplicationGatewayWafDynamicManifestsDefaultClient(subscriptionID string) ApplicationGatewayWafDynamicManifestsDefaultClient {
    return NewApplicationGatewayWafDynamicManifestsDefaultClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewayWafDynamicManifestsDefaultClientWithBaseURI creates an instance of the
// ApplicationGatewayWafDynamicManifestsDefaultClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewApplicationGatewayWafDynamicManifestsDefaultClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewayWafDynamicManifestsDefaultClient {
        return ApplicationGatewayWafDynamicManifestsDefaultClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get gets the regional application gateway waf manifest.
    // Parameters:
        // location - the region where the nrp are located at.
func (client ApplicationGatewayWafDynamicManifestsDefaultClient) Get(ctx context.Context, location string) (result ApplicationGatewayWafDynamicManifestResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ApplicationGatewayWafDynamicManifestsDefaultClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, location)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsDefaultClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsDefaultClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsDefaultClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ApplicationGatewayWafDynamicManifestsDefaultClient) GetPreparer(ctx context.Context, location string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "location": autorest.Encode("path",location),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-07-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/applicationGatewayWafDynamicManifests/dafault",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ApplicationGatewayWafDynamicManifestsDefaultClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ApplicationGatewayWafDynamicManifestsDefaultClient) GetResponder(resp *http.Response) (result ApplicationGatewayWafDynamicManifestResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

