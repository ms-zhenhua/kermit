package synapse

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

// NotebookOperationResultClient is the client for the NotebookOperationResult methods of the Synapse service.
type NotebookOperationResultClient struct {
	BaseClient
}

// NewNotebookOperationResultClient creates an instance of the NotebookOperationResultClient client.
func NewNotebookOperationResultClient(endpoint string) NotebookOperationResultClient {
	return NotebookOperationResultClient{New(endpoint)}
}

// Get get notebook operation result
// Parameters:
// operationID - operation ID.
func (client NotebookOperationResultClient) Get(ctx context.Context, operationID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NotebookOperationResultClient.Get")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.NotebookOperationResultClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.NotebookOperationResultClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.NotebookOperationResultClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client NotebookOperationResultClient) GetPreparer(ctx context.Context, operationID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"operationId": autorest.Encode("path", operationID),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/notebookOperationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NotebookOperationResultClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NotebookOperationResultClient) GetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
