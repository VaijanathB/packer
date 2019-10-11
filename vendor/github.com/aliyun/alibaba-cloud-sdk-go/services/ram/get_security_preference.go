package ram

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetSecurityPreference invokes the ram.GetSecurityPreference API synchronously
// api document: https://help.aliyun.com/api/ram/getsecuritypreference.html
func (client *Client) GetSecurityPreference(request *GetSecurityPreferenceRequest) (response *GetSecurityPreferenceResponse, err error) {
	response = CreateGetSecurityPreferenceResponse()
	err = client.DoAction(request, response)
	return
}

// GetSecurityPreferenceWithChan invokes the ram.GetSecurityPreference API asynchronously
// api document: https://help.aliyun.com/api/ram/getsecuritypreference.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSecurityPreferenceWithChan(request *GetSecurityPreferenceRequest) (<-chan *GetSecurityPreferenceResponse, <-chan error) {
	responseChan := make(chan *GetSecurityPreferenceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSecurityPreference(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetSecurityPreferenceWithCallback invokes the ram.GetSecurityPreference API asynchronously
// api document: https://help.aliyun.com/api/ram/getsecuritypreference.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSecurityPreferenceWithCallback(request *GetSecurityPreferenceRequest, callback func(response *GetSecurityPreferenceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSecurityPreferenceResponse
		var err error
		defer close(result)
		response, err = client.GetSecurityPreference(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetSecurityPreferenceRequest is the request struct for api GetSecurityPreference
type GetSecurityPreferenceRequest struct {
	*requests.RpcRequest
}

// GetSecurityPreferenceResponse is the response struct for api GetSecurityPreference
type GetSecurityPreferenceResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	SecurityPreference SecurityPreference `json:"SecurityPreference" xml:"SecurityPreference"`
}

// CreateGetSecurityPreferenceRequest creates a request to invoke GetSecurityPreference API
func CreateGetSecurityPreferenceRequest() (request *GetSecurityPreferenceRequest) {
	request = &GetSecurityPreferenceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "GetSecurityPreference", "ram", "openAPI")
	return
}

// CreateGetSecurityPreferenceResponse creates a response to parse from GetSecurityPreference response
func CreateGetSecurityPreferenceResponse() (response *GetSecurityPreferenceResponse) {
	response = &GetSecurityPreferenceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
