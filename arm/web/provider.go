package web

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

// ProviderClient is the use these APIs to manage Azure Websites resources
// through the Azure Resource Manager. All task operations conform to the
// HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type ProviderClient struct {
	ManagementClient
}

// NewProviderClient creates an instance of the ProviderClient client.
func NewProviderClient(subscriptionID string) ProviderClient {
	return NewProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderClientWithBaseURI creates an instance of the ProviderClient
// client.
func NewProviderClientWithBaseURI(baseURI string, subscriptionID string) ProviderClient {
	return ProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetPublishingUser sends the get publishing user request.
func (client ProviderClient) GetPublishingUser() (result User, err error) {
	req, err := client.GetPublishingUserPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "GetPublishingUser", nil, "Failure preparing request")
	}

	resp, err := client.GetPublishingUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "GetPublishingUser", resp, "Failure sending request")
	}

	result, err = client.GetPublishingUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetPublishingUser", resp, "Failure responding to request")
	}

	return
}

// GetPublishingUserPreparer prepares the GetPublishingUser request.
func (client ProviderClient) GetPublishingUserPreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetPublishingUserSender sends the GetPublishingUser request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetPublishingUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetPublishingUserResponder handles the response to the GetPublishingUser request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetPublishingUserResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSourceControl sends the get source control request.
//
// sourceControlType is type of source control
func (client ProviderClient) GetSourceControl(sourceControlType string) (result SourceControl, err error) {
	req, err := client.GetSourceControlPreparer(sourceControlType)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControl", nil, "Failure preparing request")
	}

	resp, err := client.GetSourceControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControl", resp, "Failure sending request")
	}

	result, err = client.GetSourceControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControl", resp, "Failure responding to request")
	}

	return
}

// GetSourceControlPreparer prepares the GetSourceControl request.
func (client ProviderClient) GetSourceControlPreparer(sourceControlType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceControlType": autorest.Encode("path", sourceControlType),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSourceControlSender sends the GetSourceControl request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetSourceControlSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetSourceControlResponder handles the response to the GetSourceControl request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetSourceControlResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSourceControls sends the get source controls request.
func (client ProviderClient) GetSourceControls() (result SourceControlCollection, err error) {
	req, err := client.GetSourceControlsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControls", nil, "Failure preparing request")
	}

	resp, err := client.GetSourceControlsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControls", resp, "Failure sending request")
	}

	result, err = client.GetSourceControlsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControls", resp, "Failure responding to request")
	}

	return
}

// GetSourceControlsPreparer prepares the GetSourceControls request.
func (client ProviderClient) GetSourceControlsPreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/sourcecontrols"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSourceControlsSender sends the GetSourceControls request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetSourceControlsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetSourceControlsResponder handles the response to the GetSourceControls request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetSourceControlsResponder(resp *http.Response) (result SourceControlCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdatePublishingUser sends the update publishing user request.
//
// requestMessage is details of publishing user
func (client ProviderClient) UpdatePublishingUser(requestMessage User) (result User, err error) {
	req, err := client.UpdatePublishingUserPreparer(requestMessage)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "UpdatePublishingUser", nil, "Failure preparing request")
	}

	resp, err := client.UpdatePublishingUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "UpdatePublishingUser", resp, "Failure sending request")
	}

	result, err = client.UpdatePublishingUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdatePublishingUser", resp, "Failure responding to request")
	}

	return
}

// UpdatePublishingUserPreparer prepares the UpdatePublishingUser request.
func (client ProviderClient) UpdatePublishingUserPreparer(requestMessage User) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithJSON(requestMessage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdatePublishingUserSender sends the UpdatePublishingUser request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) UpdatePublishingUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdatePublishingUserResponder handles the response to the UpdatePublishingUser request. The method always
// closes the http.Response Body.
func (client ProviderClient) UpdatePublishingUserResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateSourceControl sends the update source control request.
//
// sourceControlType is type of source control requestMessage is source
// control token information
func (client ProviderClient) UpdateSourceControl(sourceControlType string, requestMessage SourceControl) (result SourceControl, err error) {
	req, err := client.UpdateSourceControlPreparer(sourceControlType, requestMessage)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "UpdateSourceControl", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSourceControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "UpdateSourceControl", resp, "Failure sending request")
	}

	result, err = client.UpdateSourceControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdateSourceControl", resp, "Failure responding to request")
	}

	return
}

// UpdateSourceControlPreparer prepares the UpdateSourceControl request.
func (client ProviderClient) UpdateSourceControlPreparer(sourceControlType string, requestMessage SourceControl) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceControlType": autorest.Encode("path", sourceControlType),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", pathParameters),
		autorest.WithJSON(requestMessage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSourceControlSender sends the UpdateSourceControl request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) UpdateSourceControlSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateSourceControlResponder handles the response to the UpdateSourceControl request. The method always
// closes the http.Response Body.
func (client ProviderClient) UpdateSourceControlResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
