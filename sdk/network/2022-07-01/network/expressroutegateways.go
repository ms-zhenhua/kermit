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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// ExpressRouteGatewaysClient is the network Client
type ExpressRouteGatewaysClient struct {
	BaseClient
}

// NewExpressRouteGatewaysClient creates an instance of the ExpressRouteGatewaysClient client.
func NewExpressRouteGatewaysClient(subscriptionID string) ExpressRouteGatewaysClient {
	return NewExpressRouteGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteGatewaysClientWithBaseURI creates an instance of the ExpressRouteGatewaysClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewExpressRouteGatewaysClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteGatewaysClient {
	return ExpressRouteGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a ExpressRoute gateway in a specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
// putExpressRouteGatewayParameters - parameters required in an ExpressRoute gateway PUT operation.
func (client ExpressRouteGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway) (result ExpressRouteGatewaysCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteGatewaysClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: putExpressRouteGatewayParameters,
			Constraints: []validation.Constraint{{Target: "putExpressRouteGatewayParameters.ExpressRouteGatewayProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "putExpressRouteGatewayParameters.ExpressRouteGatewayProperties.VirtualHub", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("network.ExpressRouteGatewaysClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, expressRouteGatewayName, putExpressRouteGatewayParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteGatewaysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	putExpressRouteGatewayParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", pathParameters),
		autorest.WithJSON(putExpressRouteGatewayParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteGatewaysClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteGatewaysCreateOrUpdateFuture, err error) {
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
func (client ExpressRouteGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteGateway, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be
// deleted when there are no connection subresources.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
func (client ExpressRouteGatewaysClient) Delete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (result ExpressRouteGatewaysDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteGatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, expressRouteGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteGatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteGatewaysClient) DeleteSender(req *http.Request) (future ExpressRouteGatewaysDeleteFuture, err error) {
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
func (client ExpressRouteGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get fetches the details of a ExpressRoute gateway in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// expressRouteGatewayName - the name of the ExpressRoute gateway.
func (client ExpressRouteGatewaysClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (result ExpressRouteGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteGatewaysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, expressRouteGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteGatewaysClient) GetPreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteGatewaysClient) GetResponder(resp *http.Response) (result ExpressRouteGateway, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists ExpressRoute gateways in a given resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client ExpressRouteGatewaysClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ExpressRouteGatewayList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteGatewaysClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ExpressRouteGatewaysClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteGatewaysClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ExpressRouteGatewaysClient) ListByResourceGroupResponder(resp *http.Response) (result ExpressRouteGatewayList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription lists ExpressRoute gateways under a given subscription.
func (client ExpressRouteGatewaysClient) ListBySubscription(ctx context.Context) (result ExpressRouteGatewayList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteGatewaysClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ExpressRouteGatewaysClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteGatewaysClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ExpressRouteGatewaysClient) ListBySubscriptionResponder(resp *http.Response) (result ExpressRouteGatewayList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateTags updates express route gateway tags.
// Parameters:
// resourceGroupName - the resource group name of the ExpressRouteGateway.
// expressRouteGatewayName - the name of the gateway.
// expressRouteGatewayParameters - parameters supplied to update a virtual wan express route gateway tags.
func (client ExpressRouteGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject) (result ExpressRouteGatewaysUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteGatewaysClient.UpdateTags")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, expressRouteGatewayName, expressRouteGatewayParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteGatewaysClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client ExpressRouteGatewaysClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"expressRouteGatewayName": autorest.Encode("path", expressRouteGatewayName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", pathParameters),
		autorest.WithJSON(expressRouteGatewayParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteGatewaysClient) UpdateTagsSender(req *http.Request) (future ExpressRouteGatewaysUpdateTagsFuture, err error) {
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

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client ExpressRouteGatewaysClient) UpdateTagsResponder(resp *http.Response) (result ExpressRouteGateway, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
