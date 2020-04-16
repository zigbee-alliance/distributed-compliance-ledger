package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	BaseUrl = "http://127.0.0.1:1317"
)

func BuildUrl(uri string) string {
	return BaseUrl + "/" + uri
}

func SendGetRequest(uri string) []byte {
	return sendRequest(uri, "GET", []byte{}, "", "")
}

func SendPostRequest(uri string, body []byte, account string, passphrase string) []byte {
	return sendRequest(uri, "POST", body, account, passphrase)
}

func SendPutRequest(uri string, body []byte, account string, passphrase string) []byte {
	return sendRequest(uri, "PUT", body, account, passphrase)
}

func sendRequest(uri string, method string, body []byte, account string, passphrase string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest(method, BuildUrl(uri), bytes.NewBuffer(body))
	if len(account) != 0 && len(passphrase) != 0 {
		req.SetBasicAuth(account, passphrase)
	}
	resp, _ := client.Do(req)
	response := ReadResponseBody(resp)
	println(string(response))
	return response
}

func ReadResponseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
