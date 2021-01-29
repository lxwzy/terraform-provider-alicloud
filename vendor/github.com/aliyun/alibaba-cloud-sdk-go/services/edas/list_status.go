package edas

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

// ListStatus invokes the edas.ListStatus API synchronously
func (client *Client) ListStatus(request *ListStatusRequest) (response *ListStatusResponse, err error) {
	response = CreateListStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListStatusWithChan invokes the edas.ListStatus API asynchronously
func (client *Client) ListStatusWithChan(request *ListStatusRequest) (<-chan *ListStatusResponse, <-chan error) {
	responseChan := make(chan *ListStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStatus(request)
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

// ListStatusWithCallback invokes the edas.ListStatus API asynchronously
func (client *Client) ListStatusWithCallback(request *ListStatusRequest, callback func(response *ListStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStatusResponse
		var err error
		defer close(result)
		response, err = client.ListStatus(request)
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

// ListStatusRequest is the request struct for api ListStatus
type ListStatusRequest struct {
	*requests.RoaRequest
	Ip          string `position:"Query" name:"Ip"`
	PodName     string `position:"Query" name:"PodName"`
	Source      string `position:"Query" name:"Source"`
	AccountId   string `position:"Query" name:"AccountId"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
	AppId       string `position:"Query" name:"AppId"`
	TenantId    string `position:"Query" name:"TenantId"`
	Region      string `position:"Query" name:"Region"`
	Status      string `position:"Query" name:"Status"`
}

// ListStatusResponse is the response struct for api ListStatus
type ListStatusResponse struct {
	*responses.BaseResponse
	Code      int       `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	Success   bool      `json:"Success" xml:"Success"`
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Data      []Results `json:"Data" xml:"Data"`
}

// CreateListStatusRequest creates a request to invoke ListStatus API
func CreateListStatusRequest() (request *ListStatusRequest) {
	request = &ListStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListStatus", "/pop/sp/api/mse/status/list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListStatusResponse creates a response to parse from ListStatus response
func CreateListStatusResponse() (response *ListStatusResponse) {
	response = &ListStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}