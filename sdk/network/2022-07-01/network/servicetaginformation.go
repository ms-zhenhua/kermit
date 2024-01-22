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

// ServiceTagInformationClient is the network Client
type ServiceTagInformationClient struct {
	BaseClient
}

// NewServiceTagInformationClient creates an instance of the ServiceTagInformationClient client.
func NewServiceTagInformationClient(subscriptionID string) ServiceTagInformationClient {
	return NewServiceTagInformationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceTagInformationClientWithBaseURI creates an instance of the ServiceTagInformationClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewServiceTagInformationClientWithBaseURI(baseURI string, subscriptionID string) ServiceTagInformationClient {
	return ServiceTagInformationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of service tag information resources with pagination.
// Parameters:
// location - the location that will be used as a reference for cloud (not as a filter based on location, you
// will get the list of service tags with prefix details across all regions but limited to the cloud that your
// subscription belongs to).
// noAddressPrefixes - do not return address prefixes for the tag(s).
// tagName - return tag information for a particular tag.
func (client ServiceTagInformationClient) List(ctx context.Context, location string, noAddressPrefixes *bool, tagName string) (result ServiceTagInformationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceTagInformationClient.List")
		defer func() {
			sc := -1
			if result.stilr.Response.Response != nil {
				sc = result.stilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, noAddressPrefixes, tagName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceTagInformationClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.stilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ServiceTagInformationClient", "List", resp, "Failure sending request")
		return
	}

	result.stilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceTagInformationClient", "List", resp, "Failure responding to request")
		return
	}
	if result.stilr.hasNextLink() && result.stilr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceTagInformationClient) ListPreparer(ctx context.Context, location string, noAddressPrefixes *bool, tagName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if noAddressPrefixes != nil {
		queryParameters["noAddressPrefixes"] = autorest.Encode("query", *noAddressPrefixes)
	}
	if len(tagName) > 0 {
		queryParameters["tagName"] = autorest.Encode("query", tagName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTagDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceTagInformationClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceTagInformationClient) ListResponder(resp *http.Response) (result ServiceTagInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ServiceTagInformationClient) listNextResults(ctx context.Context, lastResults ServiceTagInformationListResult) (result ServiceTagInformationListResult, err error) {
	req, err := lastResults.serviceTagInformationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ServiceTagInformationClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ServiceTagInformationClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceTagInformationClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceTagInformationClient) ListComplete(ctx context.Context, location string, noAddressPrefixes *bool, tagName string) (result ServiceTagInformationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceTagInformationClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location, noAddressPrefixes, tagName)
	return
}
