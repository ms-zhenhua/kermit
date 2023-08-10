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

// EnrollmentGroupsClient is the azure IoT Central is a service that makes it easy to connect, monitor, and manage your
// IoT devices at scale.
type EnrollmentGroupsClient struct {
	BaseClient
}

// NewEnrollmentGroupsClient creates an instance of the EnrollmentGroupsClient client.
func NewEnrollmentGroupsClient(subdomain string) EnrollmentGroupsClient {
	return EnrollmentGroupsClient{New(subdomain)}
}

// Create create an enrollment group.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// body - enrollment group body.
func (client EnrollmentGroupsClient) Create(ctx context.Context, enrollmentGroupID string, body EnrollmentGroup) (result EnrollmentGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.DisplayName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, enrollmentGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client EnrollmentGroupsClient) CreatePreparer(ctx context.Context, enrollmentGroupID string, body EnrollmentGroup) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.IDScope = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) CreateResponder(resp *http.Response) (result EnrollmentGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateX509 sets the primary or secondary x509 certificate of an enrollment group.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// entry - entry of certificate only support primary and secondary.
// body - certificate definition.
func (client EnrollmentGroupsClient) CreateX509(ctx context.Context, enrollmentGroupID string, entry CertificateEntry, body SigningX509Certificate) (result SigningX509Certificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.CreateX509")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Info", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Info.Sha1Thumbprint", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "CreateX509", err.Error())
	}

	req, err := client.CreateX509Preparer(ctx, enrollmentGroupID, entry, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "CreateX509", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateX509Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "CreateX509", resp, "Failure sending request")
		return
	}

	result, err = client.CreateX509Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "CreateX509", resp, "Failure responding to request")
		return
	}

	return
}

// CreateX509Preparer prepares the CreateX509 request.
func (client EnrollmentGroupsClient) CreateX509Preparer(ctx context.Context, enrollmentGroupID string, entry CertificateEntry, body SigningX509Certificate) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
		"entry":             autorest.Encode("path", entry),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.Info = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}/certificates/{entry}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateX509Sender sends the CreateX509 request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) CreateX509Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateX509Responder handles the response to the CreateX509 request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) CreateX509Responder(resp *http.Response) (result SigningX509Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GenerateVerificationCodeX509 generate a verification code for the primary or secondary x509 certificate of an
// enrollment group.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// entry - entry of certificate only support primary and secondary.
func (client EnrollmentGroupsClient) GenerateVerificationCodeX509(ctx context.Context, enrollmentGroupID string, entry CertificateEntry) (result X509VerificationCode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.GenerateVerificationCodeX509")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "GenerateVerificationCodeX509", err.Error())
	}

	req, err := client.GenerateVerificationCodeX509Preparer(ctx, enrollmentGroupID, entry)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "GenerateVerificationCodeX509", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateVerificationCodeX509Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "GenerateVerificationCodeX509", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateVerificationCodeX509Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "GenerateVerificationCodeX509", resp, "Failure responding to request")
		return
	}

	return
}

// GenerateVerificationCodeX509Preparer prepares the GenerateVerificationCodeX509 request.
func (client EnrollmentGroupsClient) GenerateVerificationCodeX509Preparer(ctx context.Context, enrollmentGroupID string, entry CertificateEntry) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
		"entry":             autorest.Encode("path", entry),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}/certificates/{entry}/generateVerificationCode", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateVerificationCodeX509Sender sends the GenerateVerificationCodeX509 request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) GenerateVerificationCodeX509Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GenerateVerificationCodeX509Responder handles the response to the GenerateVerificationCodeX509 request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) GenerateVerificationCodeX509Responder(resp *http.Response) (result X509VerificationCode, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get details about an enrollment group by ID.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
func (client EnrollmentGroupsClient) Get(ctx context.Context, enrollmentGroupID string) (result EnrollmentGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, enrollmentGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client EnrollmentGroupsClient) GetPreparer(ctx context.Context, enrollmentGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) GetResponder(resp *http.Response) (result EnrollmentGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetX509 get details about the primary or secondary x509 certificate of an enrollment group, if it exists.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// entry - entry of certificate only support primary and secondary.
func (client EnrollmentGroupsClient) GetX509(ctx context.Context, enrollmentGroupID string, entry CertificateEntry) (result SigningX509Certificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.GetX509")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "GetX509", err.Error())
	}

	req, err := client.GetX509Preparer(ctx, enrollmentGroupID, entry)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "GetX509", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetX509Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "GetX509", resp, "Failure sending request")
		return
	}

	result, err = client.GetX509Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "GetX509", resp, "Failure responding to request")
		return
	}

	return
}

// GetX509Preparer prepares the GetX509 request.
func (client EnrollmentGroupsClient) GetX509Preparer(ctx context.Context, enrollmentGroupID string, entry CertificateEntry) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
		"entry":             autorest.Encode("path", entry),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}/certificates/{entry}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetX509Sender sends the GetX509 request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) GetX509Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetX509Responder handles the response to the GetX509 request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) GetX509Responder(resp *http.Response) (result SigningX509Certificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
func (client EnrollmentGroupsClient) List(ctx context.Context) (result EnrollmentGroupCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.List")
		defer func() {
			sc := -1
			if result.egc.Response.Response != nil {
				sc = result.egc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.egc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.egc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.egc.hasNextLink() && result.egc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EnrollmentGroupsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
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
		autorest.WithPath("/enrollmentGroups"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) ListResponder(resp *http.Response) (result EnrollmentGroupCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EnrollmentGroupsClient) listNextResults(ctx context.Context, lastResults EnrollmentGroupCollection) (result EnrollmentGroupCollection, err error) {
	req, err := lastResults.enrollmentGroupCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EnrollmentGroupsClient) ListComplete(ctx context.Context) (result EnrollmentGroupCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.List")
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

// Remove delete an enrollment group by ID.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
func (client EnrollmentGroupsClient) Remove(ctx context.Context, enrollmentGroupID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.Remove")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "Remove", err.Error())
	}

	req, err := client.RemovePreparer(ctx, enrollmentGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Remove", resp, "Failure responding to request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client EnrollmentGroupsClient) RemovePreparer(ctx context.Context, enrollmentGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RemoveX509 removes the primary or secondary x509 certificate of an enrollment group.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// entry - entry of certificate only support primary and secondary.
func (client EnrollmentGroupsClient) RemoveX509(ctx context.Context, enrollmentGroupID string, entry CertificateEntry) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.RemoveX509")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "RemoveX509", err.Error())
	}

	req, err := client.RemoveX509Preparer(ctx, enrollmentGroupID, entry)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "RemoveX509", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveX509Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "RemoveX509", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveX509Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "RemoveX509", resp, "Failure responding to request")
		return
	}

	return
}

// RemoveX509Preparer prepares the RemoveX509 request.
func (client EnrollmentGroupsClient) RemoveX509Preparer(ctx context.Context, enrollmentGroupID string, entry CertificateEntry) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
		"entry":             autorest.Encode("path", entry),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}/certificates/{entry}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveX509Sender sends the RemoveX509 request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) RemoveX509Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveX509Responder handles the response to the RemoveX509 request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) RemoveX509Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update an enrollment group.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// body - enrollment group patch body.
// ifMatch - only perform the operation if the entity's etag matches one of the etags provided or * is
// provided.
func (client EnrollmentGroupsClient) Update(ctx context.Context, enrollmentGroupID string, body interface{}, ifMatch string) (result EnrollmentGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, enrollmentGroupID, body, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EnrollmentGroupsClient) UpdatePreparer(ctx context.Context, enrollmentGroupID string, body interface{}, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/merge-patch+json"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}", pathParameters),
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
func (client EnrollmentGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) UpdateResponder(resp *http.Response) (result EnrollmentGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// VerifyX509 verify the primary or secondary x509 certificate of an enrollment group by providing a certificate with
// the signed verification code.
// Parameters:
// enrollmentGroupID - unique ID of the enrollment group.
// entry - entry of certificate only support primary and secondary.
// body - certificate for the signed verification code.
func (client EnrollmentGroupsClient) VerifyX509(ctx context.Context, enrollmentGroupID string, entry CertificateEntry, body X509VerificationCertificate) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentGroupsClient.VerifyX509")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: enrollmentGroupID,
			Constraints: []validation.Constraint{{Target: "enrollmentGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "enrollmentGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Certificate", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.EnrollmentGroupsClient", "VerifyX509", err.Error())
	}

	req, err := client.VerifyX509Preparer(ctx, enrollmentGroupID, entry, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "VerifyX509", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifyX509Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "VerifyX509", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyX509Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.EnrollmentGroupsClient", "VerifyX509", resp, "Failure responding to request")
		return
	}

	return
}

// VerifyX509Preparer prepares the VerifyX509 request.
func (client EnrollmentGroupsClient) VerifyX509Preparer(ctx context.Context, enrollmentGroupID string, entry CertificateEntry, body X509VerificationCertificate) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"enrollmentGroupId": autorest.Encode("path", enrollmentGroupID),
		"entry":             autorest.Encode("path", entry),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/enrollmentGroups/{enrollmentGroupId}/certificates/{entry}/verify", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// VerifyX509Sender sends the VerifyX509 request. The method will close the
// http.Response Body if it receives an error.
func (client EnrollmentGroupsClient) VerifyX509Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// VerifyX509Responder handles the response to the VerifyX509 request. The method always
// closes the http.Response Body.
func (client EnrollmentGroupsClient) VerifyX509Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}