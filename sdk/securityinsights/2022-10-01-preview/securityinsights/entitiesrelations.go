package securityinsight

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

// EntitiesRelationsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type EntitiesRelationsClient struct {
	BaseClient
}

// NewEntitiesRelationsClient creates an instance of the EntitiesRelationsClient client.
func NewEntitiesRelationsClient(subscriptionID string) EntitiesRelationsClient {
	return NewEntitiesRelationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEntitiesRelationsClientWithBaseURI creates an instance of the EntitiesRelationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewEntitiesRelationsClientWithBaseURI(baseURI string, subscriptionID string) EntitiesRelationsClient {
	return EntitiesRelationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets all relations of an entity.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// entityID - entity ID
// filter - filters the results, based on a Boolean condition. Optional.
// orderby - sorts the results. Optional.
// top - returns only the first n results. Optional.
// skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls. Optional.
func (client EntitiesRelationsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (result RelationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntitiesRelationsClient.List")
		defer func() {
			sc := -1
			if result.rl.Response.Response != nil {
				sc = result.rl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.EntitiesRelationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, entityID, filter, orderby, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntitiesRelationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.EntitiesRelationsClient", "List", resp, "Failure sending request")
		return
	}

	result.rl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntitiesRelationsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rl.hasNextLink() && result.rl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client EntitiesRelationsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"entityId":          autorest.Encode("path", entityID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2022-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/relations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EntitiesRelationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EntitiesRelationsClient) ListResponder(resp *http.Response) (result RelationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EntitiesRelationsClient) listNextResults(ctx context.Context, lastResults RelationList) (result RelationList, err error) {
	req, err := lastResults.relationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.EntitiesRelationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.EntitiesRelationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.EntitiesRelationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EntitiesRelationsClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, filter string, orderby string, top *int32, skipToken string) (result RelationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntitiesRelationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName, entityID, filter, orderby, top, skipToken)
	return
}
