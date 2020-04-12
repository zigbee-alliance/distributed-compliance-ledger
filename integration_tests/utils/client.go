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
	resp, _ := http.Get(BuildUrl(uri))
	response := ReadResponseBody(resp)
	println(string(response))
	return response
}

func SendPostRequest(uri string, body []byte) []byte {
	resp, _ := http.Post(BuildUrl(uri), "text", bytes.NewBuffer(body))
	response := ReadResponseBody(resp)
	println(string(response))
	return response

}

func ReadResponseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
