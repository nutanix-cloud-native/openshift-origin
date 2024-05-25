package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualMachineRunCommandsClient is the compute Client
type VirtualMachineRunCommandsClient struct {
	BaseClient
}

// NewVirtualMachineRunCommandsClient creates an instance of the VirtualMachineRunCommandsClient client.
func NewVirtualMachineRunCommandsClient(subscriptionID string) VirtualMachineRunCommandsClient {
	return NewVirtualMachineRunCommandsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineRunCommandsClientWithBaseURI creates an instance of the VirtualMachineRunCommandsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewVirtualMachineRunCommandsClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineRunCommandsClient {
	return VirtualMachineRunCommandsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update the run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMName - the name of the virtual machine where the run command should be created or updated.
// runCommandName - the name of the virtual machine run command.
// runCommand - parameters supplied to the Create Virtual Machine RunCommand operation.
func (client VirtualMachineRunCommandsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, runCommandName string, runCommand VirtualMachineRunCommand) (result VirtualMachineRunCommandsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, VMName, runCommandName, runCommand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineRunCommandsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, VMName string, runCommandName string, runCommand VirtualMachineRunCommand) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", VMName),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}", pathParameters),
		autorest.WithJSON(runCommand),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) CreateOrUpdateSender(req *http.Request) (future VirtualMachineRunCommandsCreateOrUpdateFuture, err error) {
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
func (client VirtualMachineRunCommandsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualMachineRunCommand, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete the run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMName - the name of the virtual machine where the run command should be deleted.
// runCommandName - the name of the virtual machine run command.
func (client VirtualMachineRunCommandsClient) Delete(ctx context.Context, resourceGroupName string, VMName string, runCommandName string) (result VirtualMachineRunCommandsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, VMName, runCommandName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineRunCommandsClient) DeletePreparer(ctx context.Context, resourceGroupName string, VMName string, runCommandName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", VMName),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) DeleteSender(req *http.Request) (future VirtualMachineRunCommandsDeleteFuture, err error) {
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
func (client VirtualMachineRunCommandsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets specific run command for a subscription in a location.
// Parameters:
// location - the location upon which run commands is queried.
// commandID - the command id.
func (client VirtualMachineRunCommandsClient) Get(ctx context.Context, location string, commandID string) (result RunCommandDocument, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.VirtualMachineRunCommandsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, location, commandID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineRunCommandsClient) GetPreparer(ctx context.Context, location string, commandID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"commandId":      autorest.Encode("path", commandID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) GetResponder(resp *http.Response) (result RunCommandDocument, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByVirtualMachine the operation to get the run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMName - the name of the virtual machine containing the run command.
// runCommandName - the name of the virtual machine run command.
// expand - the expand expression to apply on the operation.
func (client VirtualMachineRunCommandsClient) GetByVirtualMachine(ctx context.Context, resourceGroupName string, VMName string, runCommandName string, expand string) (result VirtualMachineRunCommand, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.GetByVirtualMachine")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByVirtualMachinePreparer(ctx, resourceGroupName, VMName, runCommandName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "GetByVirtualMachine", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByVirtualMachineSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "GetByVirtualMachine", resp, "Failure sending request")
		return
	}

	result, err = client.GetByVirtualMachineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "GetByVirtualMachine", resp, "Failure responding to request")
		return
	}

	return
}

// GetByVirtualMachinePreparer prepares the GetByVirtualMachine request.
func (client VirtualMachineRunCommandsClient) GetByVirtualMachinePreparer(ctx context.Context, resourceGroupName string, VMName string, runCommandName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", VMName),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByVirtualMachineSender sends the GetByVirtualMachine request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) GetByVirtualMachineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByVirtualMachineResponder handles the response to the GetByVirtualMachine request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) GetByVirtualMachineResponder(resp *http.Response) (result VirtualMachineRunCommand, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all available run commands for a subscription in a location.
// Parameters:
// location - the location upon which run commands is queried.
func (client VirtualMachineRunCommandsClient) List(ctx context.Context, location string) (result RunCommandListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.List")
		defer func() {
			sc := -1
			if result.rclr.Response.Response != nil {
				sc = result.rclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.VirtualMachineRunCommandsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", resp, "Failure sending request")
		return
	}

	result.rclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rclr.hasNextLink() && result.rclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineRunCommandsClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) ListResponder(resp *http.Response) (result RunCommandListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualMachineRunCommandsClient) listNextResults(ctx context.Context, lastResults RunCommandListResult) (result RunCommandListResult, err error) {
	req, err := lastResults.runCommandListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineRunCommandsClient) ListComplete(ctx context.Context, location string) (result RunCommandListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}

// ListByVirtualMachine the operation to get all run commands of a Virtual Machine.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMName - the name of the virtual machine containing the run command.
// expand - the expand expression to apply on the operation.
func (client VirtualMachineRunCommandsClient) ListByVirtualMachine(ctx context.Context, resourceGroupName string, VMName string, expand string) (result VirtualMachineRunCommandsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.ListByVirtualMachine")
		defer func() {
			sc := -1
			if result.vmrclr.Response.Response != nil {
				sc = result.vmrclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVirtualMachineNextResults
	req, err := client.ListByVirtualMachinePreparer(ctx, resourceGroupName, VMName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "ListByVirtualMachine", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVirtualMachineSender(req)
	if err != nil {
		result.vmrclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "ListByVirtualMachine", resp, "Failure sending request")
		return
	}

	result.vmrclr, err = client.ListByVirtualMachineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "ListByVirtualMachine", resp, "Failure responding to request")
		return
	}
	if result.vmrclr.hasNextLink() && result.vmrclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByVirtualMachinePreparer prepares the ListByVirtualMachine request.
func (client VirtualMachineRunCommandsClient) ListByVirtualMachinePreparer(ctx context.Context, resourceGroupName string, VMName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", VMName),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVirtualMachineSender sends the ListByVirtualMachine request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) ListByVirtualMachineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVirtualMachineResponder handles the response to the ListByVirtualMachine request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) ListByVirtualMachineResponder(resp *http.Response) (result VirtualMachineRunCommandsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVirtualMachineNextResults retrieves the next set of results, if any.
func (client VirtualMachineRunCommandsClient) listByVirtualMachineNextResults(ctx context.Context, lastResults VirtualMachineRunCommandsListResult) (result VirtualMachineRunCommandsListResult, err error) {
	req, err := lastResults.virtualMachineRunCommandsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "listByVirtualMachineNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVirtualMachineSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "listByVirtualMachineNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVirtualMachineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "listByVirtualMachineNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVirtualMachineComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineRunCommandsClient) ListByVirtualMachineComplete(ctx context.Context, resourceGroupName string, VMName string, expand string) (result VirtualMachineRunCommandsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.ListByVirtualMachine")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVirtualMachine(ctx, resourceGroupName, VMName, expand)
	return
}

// Update the operation to update the run command.
// Parameters:
// resourceGroupName - the name of the resource group.
// VMName - the name of the virtual machine where the run command should be updated.
// runCommandName - the name of the virtual machine run command.
// runCommand - parameters supplied to the Update Virtual Machine RunCommand operation.
func (client VirtualMachineRunCommandsClient) Update(ctx context.Context, resourceGroupName string, VMName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate) (result VirtualMachineRunCommandsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineRunCommandsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, VMName, runCommandName, runCommand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineRunCommandsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualMachineRunCommandsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, VMName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runCommandName":    autorest.Encode("path", runCommandName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vmName":            autorest.Encode("path", VMName),
	}

	const APIVersion = "2022-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}", pathParameters),
		autorest.WithJSON(runCommand),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineRunCommandsClient) UpdateSender(req *http.Request) (future VirtualMachineRunCommandsUpdateFuture, err error) {
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

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualMachineRunCommandsClient) UpdateResponder(resp *http.Response) (result VirtualMachineRunCommand, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
