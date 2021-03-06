package drds

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

// RemoveDataExportTask invokes the drds.RemoveDataExportTask API synchronously
func (client *Client) RemoveDataExportTask(request *RemoveDataExportTaskRequest) (response *RemoveDataExportTaskResponse, err error) {
	response = CreateRemoveDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDataExportTaskWithChan invokes the drds.RemoveDataExportTask API asynchronously
func (client *Client) RemoveDataExportTaskWithChan(request *RemoveDataExportTaskRequest) (<-chan *RemoveDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *RemoveDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDataExportTask(request)
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

// RemoveDataExportTaskWithCallback invokes the drds.RemoveDataExportTask API asynchronously
func (client *Client) RemoveDataExportTaskWithCallback(request *RemoveDataExportTaskRequest, callback func(response *RemoveDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.RemoveDataExportTask(request)
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

// RemoveDataExportTaskRequest is the request struct for api RemoveDataExportTask
type RemoveDataExportTaskRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// RemoveDataExportTaskResponse is the response struct for api RemoveDataExportTask
type RemoveDataExportTaskResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	TaskManageResult TaskManageResult `json:"TaskManageResult" xml:"TaskManageResult"`
}

// CreateRemoveDataExportTaskRequest creates a request to invoke RemoveDataExportTask API
func CreateRemoveDataExportTaskRequest() (request *RemoveDataExportTaskRequest) {
	request = &RemoveDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "RemoveDataExportTask", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveDataExportTaskResponse creates a response to parse from RemoveDataExportTask response
func CreateRemoveDataExportTaskResponse() (response *RemoveDataExportTaskResponse) {
	response = &RemoveDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
