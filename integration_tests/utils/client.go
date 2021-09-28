// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BaseURL = "http://127.0.0.1:1317"
)

func BuildURL(uri string) string {
	return BaseURL + "/" + uri
}

func SendGetRequest(uri string) ([]byte, int) {
	return sendRequest(uri, "GET", []byte{}, "", "")
}

func SendPostRequest(uri string, body []byte, account string, passphrase string) ([]byte, int) {
	return sendRequest(uri, "POST", body, account, passphrase)
}

func SendPutRequest(uri string, body []byte, account string, passphrase string) ([]byte, int) {
	return sendRequest(uri, "PUT", body, account, passphrase)
}

func SendPatchRequest(uri string, body []byte, account string, passphrase string) ([]byte, int) {
	return sendRequest(uri, "PATCH", body, account, passphrase)
}

func SendDeleteRequest(uri string, body []byte, account string, passphrase string) ([]byte, int) {
	return sendRequest(uri, "DELETE", body, account, passphrase)
}

func sendRequest(uri string, method string, body []byte, account string, passphrase string) ([]byte, int) {
	if len(account) == 0 {
		passphrase = ""
	}

	client := &http.Client{}
	//nolint:noctx
	req, err := http.NewRequest(method, BuildURL(uri), bytes.NewBuffer(body))
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	if len(account) != 0 && len(passphrase) != 0 {
		req.SetBasicAuth(account, passphrase)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error received from server: %v", err)
		return nil, http.StatusInternalServerError
	}

	response := ReadResponseBody(resp)
	println(string(response))

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode
	}

	return response, http.StatusOK
}

func ReadResponseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
