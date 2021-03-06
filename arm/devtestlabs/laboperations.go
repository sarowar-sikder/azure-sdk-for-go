package devtestlabs

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// LabOperationsClient is the the DevTest Labs Client.
type LabOperationsClient struct {
	ManagementClient
}

// NewLabOperationsClient creates an instance of the LabOperationsClient
// client.
func NewLabOperationsClient(subscriptionID string) LabOperationsClient {
	return NewLabOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLabOperationsClientWithBaseURI creates an instance of the
// LabOperationsClient client.
func NewLabOperationsClientWithBaseURI(baseURI string, subscriptionID string) LabOperationsClient {
	return LabOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateEnvironment create virtual machines in a Lab. This operation can take
// a while to complete. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used
// to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) CreateEnvironment(resourceGroupName string, name string, labVirtualMachine LabVirtualMachine, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateEnvironmentPreparer(resourceGroupName, name, labVirtualMachine, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "CreateEnvironment", nil, "Failure preparing request")
	}

	resp, err := client.CreateEnvironmentSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "CreateEnvironment", resp, "Failure sending request")
	}

	result, err = client.CreateEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "CreateEnvironment", resp, "Failure responding to request")
	}

	return
}

// CreateEnvironmentPreparer prepares the CreateEnvironment request.
func (client LabOperationsClient) CreateEnvironmentPreparer(resourceGroupName string, name string, labVirtualMachine LabVirtualMachine, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/createEnvironment", pathParameters),
		autorest.WithJSON(labVirtualMachine),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateEnvironmentSender sends the CreateEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) CreateEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateEnvironmentResponder handles the response to the CreateEnvironment request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) CreateEnvironmentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdateResource create or replace an existing Lab. This operation
// can take a while to complete. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will
// be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) CreateOrUpdateResource(resourceGroupName string, name string, lab Lab, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdateResourcePreparer(resourceGroupName, name, lab, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "CreateOrUpdateResource", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateResourceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "CreateOrUpdateResource", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "CreateOrUpdateResource", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client LabOperationsClient) CreateOrUpdateResourcePreparer(resourceGroupName string, name string, lab Lab, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithJSON(lab),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) CreateOrUpdateResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) CreateOrUpdateResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteResource delete lab. This operation can take a while to complete.
// This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling
// and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) DeleteResource(resourceGroupName string, name string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeleteResourcePreparer(resourceGroupName, name, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "DeleteResource", nil, "Failure preparing request")
	}

	resp, err := client.DeleteResourceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "DeleteResource", resp, "Failure sending request")
	}

	result, err = client.DeleteResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "DeleteResource", resp, "Failure responding to request")
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client LabOperationsClient) DeleteResourcePreparer(resourceGroupName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) DeleteResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateUploadURI generate a URI for uploading custom disk images to a Lab.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) GenerateUploadURI(resourceGroupName string, name string, generateUploadURIParameter GenerateUploadURIParameter) (result GenerateUploadURIResponse, err error) {
	req, err := client.GenerateUploadURIPreparer(resourceGroupName, name, generateUploadURIParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "GenerateUploadURI", nil, "Failure preparing request")
	}

	resp, err := client.GenerateUploadURISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "GenerateUploadURI", resp, "Failure sending request")
	}

	result, err = client.GenerateUploadURIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "GenerateUploadURI", resp, "Failure responding to request")
	}

	return
}

// GenerateUploadURIPreparer prepares the GenerateUploadURI request.
func (client LabOperationsClient) GenerateUploadURIPreparer(resourceGroupName string, name string, generateUploadURIParameter GenerateUploadURIParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/generateUploadUri", pathParameters),
		autorest.WithJSON(generateUploadURIParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GenerateUploadURISender sends the GenerateUploadURI request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) GenerateUploadURISender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GenerateUploadURIResponder handles the response to the GenerateUploadURI request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) GenerateUploadURIResponder(resp *http.Response) (result GenerateUploadURIResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetResource get lab.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) GetResource(resourceGroupName string, name string) (result Lab, err error) {
	req, err := client.GetResourcePreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "GetResource", nil, "Failure preparing request")
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "GetResource", resp, "Failure sending request")
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client LabOperationsClient) GetResourcePreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) GetResourceResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list labs in a resource group.
//
// resourceGroupName is the name of the resource group. filter is the filter
// to apply on the operation. top is the maximum number of resources to
// return from the operation. orderBy is the ordering expression for the
// results, using OData notation.
func (client LabOperationsClient) ListByResourceGroup(resourceGroupName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationLab, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, filter, top, orderBy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client LabOperationsClient) ListByResourceGroupPreparer(resourceGroupName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) ListByResourceGroupResponder(resp *http.Response) (result ResponseWithContinuationLab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client LabOperationsClient) ListByResourceGroupNextResults(lastResults ResponseWithContinuationLab) (result ResponseWithContinuationLab, err error) {
	req, err := lastResults.ResponseWithContinuationLabPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListByResourceGroup", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListByResourceGroup", resp, "Failure sending next results request request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListByResourceGroup", resp, "Failure responding to next results request request")
	}

	return
}

// ListBySubscription list labs in a subscription.
//
// filter is the filter to apply on the operation. top is the maximum number
// of resources to return from the operation. orderBy is the ordering
// expression for the results, using OData notation.
func (client LabOperationsClient) ListBySubscription(filter string, top *int32, orderBy string) (result ResponseWithContinuationLab, err error) {
	req, err := client.ListBySubscriptionPreparer(filter, top, orderBy)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListBySubscription", nil, "Failure preparing request")
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListBySubscription", resp, "Failure sending request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client LabOperationsClient) ListBySubscriptionPreparer(filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) ListBySubscriptionResponder(resp *http.Response) (result ResponseWithContinuationLab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionNextResults retrieves the next set of results, if any.
func (client LabOperationsClient) ListBySubscriptionNextResults(lastResults ResponseWithContinuationLab) (result ResponseWithContinuationLab, err error) {
	req, err := lastResults.ResponseWithContinuationLabPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListBySubscription", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListBySubscription", resp, "Failure sending next results request request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListBySubscription", resp, "Failure responding to next results request request")
	}

	return
}

// ListVhds list disk images available for custom image creation.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) ListVhds(resourceGroupName string, name string) (result ResponseWithContinuationLabVhd, err error) {
	req, err := client.ListVhdsPreparer(resourceGroupName, name)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListVhds", nil, "Failure preparing request")
	}

	resp, err := client.ListVhdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListVhds", resp, "Failure sending request")
	}

	result, err = client.ListVhdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListVhds", resp, "Failure responding to request")
	}

	return
}

// ListVhdsPreparer prepares the ListVhds request.
func (client LabOperationsClient) ListVhdsPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/listVhds", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListVhdsSender sends the ListVhds request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) ListVhdsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListVhdsResponder handles the response to the ListVhds request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) ListVhdsResponder(resp *http.Response) (result ResponseWithContinuationLabVhd, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListVhdsNextResults retrieves the next set of results, if any.
func (client LabOperationsClient) ListVhdsNextResults(lastResults ResponseWithContinuationLabVhd) (result ResponseWithContinuationLabVhd, err error) {
	req, err := lastResults.ResponseWithContinuationLabVhdPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListVhds", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListVhdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListVhds", resp, "Failure sending next results request request")
	}

	result, err = client.ListVhdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "ListVhds", resp, "Failure responding to next results request request")
	}

	return
}

// PatchResource modify properties of labs.
//
// resourceGroupName is the name of the resource group. name is the name of
// the lab.
func (client LabOperationsClient) PatchResource(resourceGroupName string, name string, lab Lab) (result Lab, err error) {
	req, err := client.PatchResourcePreparer(resourceGroupName, name, lab)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "PatchResource", nil, "Failure preparing request")
	}

	resp, err := client.PatchResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "PatchResource", resp, "Failure sending request")
	}

	result, err = client.PatchResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.LabOperationsClient", "PatchResource", resp, "Failure responding to request")
	}

	return
}

// PatchResourcePreparer prepares the PatchResource request.
func (client LabOperationsClient) PatchResourcePreparer(resourceGroupName string, name string, lab Lab) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithJSON(lab),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchResourceSender sends the PatchResource request. The method will close the
// http.Response Body if it receives an error.
func (client LabOperationsClient) PatchResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResourceResponder handles the response to the PatchResource request. The method always
// closes the http.Response Body.
func (client LabOperationsClient) PatchResourceResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
