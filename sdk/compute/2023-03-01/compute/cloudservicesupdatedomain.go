package compute

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

// CloudServicesUpdateDomainClient is the compute Client
type CloudServicesUpdateDomainClient struct {
	BaseClient
}

// NewCloudServicesUpdateDomainClient creates an instance of the CloudServicesUpdateDomainClient client.
func NewCloudServicesUpdateDomainClient(subscriptionID string) CloudServicesUpdateDomainClient {
	return NewCloudServicesUpdateDomainClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCloudServicesUpdateDomainClientWithBaseURI creates an instance of the CloudServicesUpdateDomainClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewCloudServicesUpdateDomainClientWithBaseURI(baseURI string, subscriptionID string) CloudServicesUpdateDomainClient {
	return CloudServicesUpdateDomainClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetUpdateDomain gets the specified update domain of a cloud service. Use nextLink property in the response to get
// the next page of update domains. Do this till nextLink is null to fetch all the update domains.
// Parameters:
// resourceGroupName - name of the resource group.
// cloudServiceName - name of the cloud service.
// updateDomain - specifies an integer value that identifies the update domain. Update domains are identified
// with a zero-based index: the first update domain has an ID of 0, the second has an ID of 1, and so on.
func (client CloudServicesUpdateDomainClient) GetUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32) (result UpdateDomain, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServicesUpdateDomainClient.GetUpdateDomain")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetUpdateDomainPreparer(ctx, resourceGroupName, cloudServiceName, updateDomain)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "GetUpdateDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUpdateDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "GetUpdateDomain", resp, "Failure sending request")
		return
	}

	result, err = client.GetUpdateDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "GetUpdateDomain", resp, "Failure responding to request")
		return
	}

	return
}

// GetUpdateDomainPreparer prepares the GetUpdateDomain request.
func (client CloudServicesUpdateDomainClient) GetUpdateDomainPreparer(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"updateDomain":      autorest.Encode("path", updateDomain),
	}

	const APIVersion = "2022-09-04"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains/{updateDomain}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUpdateDomainSender sends the GetUpdateDomain request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServicesUpdateDomainClient) GetUpdateDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetUpdateDomainResponder handles the response to the GetUpdateDomain request. The method always
// closes the http.Response Body.
func (client CloudServicesUpdateDomainClient) GetUpdateDomainResponder(resp *http.Response) (result UpdateDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUpdateDomains gets a list of all update domains in a cloud service.
// Parameters:
// resourceGroupName - name of the resource group.
// cloudServiceName - name of the cloud service.
func (client CloudServicesUpdateDomainClient) ListUpdateDomains(ctx context.Context, resourceGroupName string, cloudServiceName string) (result UpdateDomainListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServicesUpdateDomainClient.ListUpdateDomains")
		defer func() {
			sc := -1
			if result.udlr.Response.Response != nil {
				sc = result.udlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listUpdateDomainsNextResults
	req, err := client.ListUpdateDomainsPreparer(ctx, resourceGroupName, cloudServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "ListUpdateDomains", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUpdateDomainsSender(req)
	if err != nil {
		result.udlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "ListUpdateDomains", resp, "Failure sending request")
		return
	}

	result.udlr, err = client.ListUpdateDomainsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "ListUpdateDomains", resp, "Failure responding to request")
		return
	}
	if result.udlr.hasNextLink() && result.udlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListUpdateDomainsPreparer prepares the ListUpdateDomains request.
func (client CloudServicesUpdateDomainClient) ListUpdateDomainsPreparer(ctx context.Context, resourceGroupName string, cloudServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-04"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUpdateDomainsSender sends the ListUpdateDomains request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServicesUpdateDomainClient) ListUpdateDomainsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListUpdateDomainsResponder handles the response to the ListUpdateDomains request. The method always
// closes the http.Response Body.
func (client CloudServicesUpdateDomainClient) ListUpdateDomainsResponder(resp *http.Response) (result UpdateDomainListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listUpdateDomainsNextResults retrieves the next set of results, if any.
func (client CloudServicesUpdateDomainClient) listUpdateDomainsNextResults(ctx context.Context, lastResults UpdateDomainListResult) (result UpdateDomainListResult, err error) {
	req, err := lastResults.updateDomainListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "listUpdateDomainsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListUpdateDomainsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "listUpdateDomainsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListUpdateDomainsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "listUpdateDomainsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListUpdateDomainsComplete enumerates all values, automatically crossing page boundaries as required.
func (client CloudServicesUpdateDomainClient) ListUpdateDomainsComplete(ctx context.Context, resourceGroupName string, cloudServiceName string) (result UpdateDomainListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServicesUpdateDomainClient.ListUpdateDomains")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListUpdateDomains(ctx, resourceGroupName, cloudServiceName)
	return
}

// WalkUpdateDomain updates the role instances in the specified update domain.
// Parameters:
// resourceGroupName - name of the resource group.
// cloudServiceName - name of the cloud service.
// updateDomain - specifies an integer value that identifies the update domain. Update domains are identified
// with a zero-based index: the first update domain has an ID of 0, the second has an ID of 1, and so on.
// parameters - the update domain object.
func (client CloudServicesUpdateDomainClient) WalkUpdateDomain(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, parameters *UpdateDomain) (result CloudServicesUpdateDomainWalkUpdateDomainFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CloudServicesUpdateDomainClient.WalkUpdateDomain")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.WalkUpdateDomainPreparer(ctx, resourceGroupName, cloudServiceName, updateDomain, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "WalkUpdateDomain", nil, "Failure preparing request")
		return
	}

	result, err = client.WalkUpdateDomainSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.CloudServicesUpdateDomainClient", "WalkUpdateDomain", result.Response(), "Failure sending request")
		return
	}

	return
}

// WalkUpdateDomainPreparer prepares the WalkUpdateDomain request.
func (client CloudServicesUpdateDomainClient) WalkUpdateDomainPreparer(ctx context.Context, resourceGroupName string, cloudServiceName string, updateDomain int32, parameters *UpdateDomain) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cloudServiceName":  autorest.Encode("path", cloudServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"updateDomain":      autorest.Encode("path", updateDomain),
	}

	const APIVersion = "2022-09-04"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/updateDomains/{updateDomain}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// WalkUpdateDomainSender sends the WalkUpdateDomain request. The method will close the
// http.Response Body if it receives an error.
func (client CloudServicesUpdateDomainClient) WalkUpdateDomainSender(req *http.Request) (future CloudServicesUpdateDomainWalkUpdateDomainFuture, err error) {
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

// WalkUpdateDomainResponder handles the response to the WalkUpdateDomain request. The method always
// closes the http.Response Body.
func (client CloudServicesUpdateDomainClient) WalkUpdateDomainResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
