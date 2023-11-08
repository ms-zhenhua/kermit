package iotcentral

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

// OrganizationsClient is the azure IoT Central is a service that makes it easy to connect, monitor, and manage your
// IoT devices at scale.
type OrganizationsClient struct {
	BaseClient
}

// NewOrganizationsClient creates an instance of the OrganizationsClient client.
func NewOrganizationsClient(subdomain string) OrganizationsClient {
	return OrganizationsClient{New(subdomain)}
}

// Create sends the create request.
// Parameters:
// organizationID - unique ID for the organization.
// body - organization body.
func (client OrganizationsClient) Create(ctx context.Context, organizationID string, body Organization) (result Organization, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: organizationID,
			Constraints: []validation.Constraint{{Target: "organizationID", Name: validation.MaxLength, Rule: 48, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.OrganizationsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, organizationID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client OrganizationsClient) CreatePreparer(ctx context.Context, organizationID string, body Organization) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"organizationId": autorest.Encode("path", organizationID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/organizations/{organizationId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client OrganizationsClient) CreateResponder(resp *http.Response) (result Organization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// organizationID - unique ID for the organization.
func (client OrganizationsClient) Get(ctx context.Context, organizationID string) (result Organization, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: organizationID,
			Constraints: []validation.Constraint{{Target: "organizationID", Name: validation.MaxLength, Rule: 48, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.OrganizationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, organizationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OrganizationsClient) GetPreparer(ctx context.Context, organizationID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"organizationId": autorest.Encode("path", organizationID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/organizations/{organizationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OrganizationsClient) GetResponder(resp *http.Response) (result Organization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client OrganizationsClient) List(ctx context.Context) (result OrganizationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationsClient.List")
		defer func() {
			sc := -1
			if result.oc.Response.Response != nil {
				sc = result.oc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.oc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "List", resp, "Failure sending request")
		return
	}

	result.oc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.oc.hasNextLink() && result.oc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client OrganizationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPath("/organizations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OrganizationsClient) ListResponder(resp *http.Response) (result OrganizationCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client OrganizationsClient) listNextResults(ctx context.Context, lastResults OrganizationCollection) (result OrganizationCollection, err error) {
	req, err := lastResults.organizationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client OrganizationsClient) ListComplete(ctx context.Context) (result OrganizationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// Remove sends the remove request.
// Parameters:
// organizationID - unique ID for the organization.
func (client OrganizationsClient) Remove(ctx context.Context, organizationID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationsClient.Remove")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: organizationID,
			Constraints: []validation.Constraint{{Target: "organizationID", Name: validation.MaxLength, Rule: 48, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.OrganizationsClient", "Remove", err.Error())
	}

	req, err := client.RemovePreparer(ctx, organizationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Remove", resp, "Failure responding to request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client OrganizationsClient) RemovePreparer(ctx context.Context, organizationID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"organizationId": autorest.Encode("path", organizationID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/organizations/{organizationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client OrganizationsClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update sends the update request.
// Parameters:
// organizationID - unique ID for the organization.
// body - organization patch body.
// ifMatch - only perform the operation if the entity's etag matches one of the etags provided or * is
// provided.
func (client OrganizationsClient) Update(ctx context.Context, organizationID string, body interface{}, ifMatch string) (result Organization, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrganizationsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: organizationID,
			Constraints: []validation.Constraint{{Target: "organizationID", Name: validation.MaxLength, Rule: 48, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.OrganizationsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, organizationID, body, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.OrganizationsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client OrganizationsClient) UpdatePreparer(ctx context.Context, organizationID string, body interface{}, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"organizationId": autorest.Encode("path", organizationID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/organizations/{organizationId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OrganizationsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OrganizationsClient) UpdateResponder(resp *http.Response) (result Organization, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
