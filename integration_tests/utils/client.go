package utils

import (
	"bytes"
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

func sendRequest(uri string, method string, body []byte, account string, passphrase string) ([]byte, int) {
	if len(account) == 0 {
		passphrase = ""
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, BuildURL(uri), bytes.NewBuffer(body))

	if err != nil {
		return nil, http.StatusInternalServerError
	}

	if len(account) != 0 && len(passphrase) != 0 {
		req.SetBasicAuth(account, passphrase)
	}

	resp, err := client.Do(req)

	if err != nil {
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
