package appplatform

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

// ApplicationLiveViewsClient is the REST API for Azure Spring Apps
type ApplicationLiveViewsClient struct {
	BaseClient
}

// NewApplicationLiveViewsClient creates an instance of the ApplicationLiveViewsClient client.
func NewApplicationLiveViewsClient(subscriptionID string) ApplicationLiveViewsClient {
	return NewApplicationLiveViewsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationLiveViewsClientWithBaseURI creates an instance of the ApplicationLiveViewsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewApplicationLiveViewsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationLiveViewsClient {
	return ApplicationLiveViewsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create the default Application Live View or update the existing Application Live View.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationLiveViewName - the name of Application Live View.
// applicationLiveViewResource - parameters for the update operation
func (client ApplicationLiveViewsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, applicationLiveViewResource ApplicationLiveViewResource) (result ApplicationLiveViewsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationLiveViewsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appplatform.ApplicationLiveViewsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, applicationLiveViewName, applicationLiveViewResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationLiveViewsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string, applicationLiveViewResource ApplicationLiveViewResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationLiveViewName": autorest.Encode("path", applicationLiveViewName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serviceName":             autorest.Encode("path", serviceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews/{applicationLiveViewName}", pathParameters),
		autorest.WithJSON(applicationLiveViewResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationLiveViewsClient) CreateOrUpdateSender(req *http.Request) (future ApplicationLiveViewsCreateOrUpdateFuture, err error) {
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
func (client ApplicationLiveViewsClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationLiveViewResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete disable the default Application Live View.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationLiveViewName - the name of Application Live View.
func (client ApplicationLiveViewsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string) (result ApplicationLiveViewsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationLiveViewsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appplatform.ApplicationLiveViewsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, applicationLiveViewName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationLiveViewsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationLiveViewName": autorest.Encode("path", applicationLiveViewName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serviceName":             autorest.Encode("path", serviceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews/{applicationLiveViewName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationLiveViewsClient) DeleteSender(req *http.Request) (future ApplicationLiveViewsDeleteFuture, err error) {
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
func (client ApplicationLiveViewsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the Application Live  and its properties.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// applicationLiveViewName - the name of Application Live View.
func (client ApplicationLiveViewsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string) (result ApplicationLiveViewResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationLiveViewsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appplatform.ApplicationLiveViewsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, applicationLiveViewName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationLiveViewsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationLiveViewName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationLiveViewName": autorest.Encode("path", applicationLiveViewName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serviceName":             autorest.Encode("path", serviceName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews/{applicationLiveViewName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationLiveViewsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationLiveViewsClient) GetResponder(resp *http.Response) (result ApplicationLiveViewResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List handles requests to list all resources in a Service.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
func (client ApplicationLiveViewsClient) List(ctx context.Context, resourceGroupName string, serviceName string) (result ApplicationLiveViewResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationLiveViewsClient.List")
		defer func() {
			sc := -1
			if result.alvrc.Response.Response != nil {
				sc = result.alvrc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("appplatform.ApplicationLiveViewsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.alvrc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "List", resp, "Failure sending request")
		return
	}

	result.alvrc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.alvrc.hasNextLink() && result.alvrc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationLiveViewsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2023-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationLiveViews", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationLiveViewsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationLiveViewsClient) ListResponder(resp *http.Response) (result ApplicationLiveViewResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ApplicationLiveViewsClient) listNextResults(ctx context.Context, lastResults ApplicationLiveViewResourceCollection) (result ApplicationLiveViewResourceCollection, err error) {
	req, err := lastResults.applicationLiveViewResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.ApplicationLiveViewsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationLiveViewsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result ApplicationLiveViewResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationLiveViewsClient.List")
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
