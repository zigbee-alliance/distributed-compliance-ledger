package utils

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"io/ioutil"
	"net/http"
)

const (
	BaseUrl = "http://127.0.0.1:1317"
)

func BuildUrl(uri string) string {
	return BaseUrl + "/" + uri
}

func SendGetRequest(uri string) ([]byte, error) {
	return sendRequest(uri, "GET", []byte{}, "", "")
}

func SendPostRequest(uri string, body []byte, account string, passphrase string) ([]byte, error) {
	return sendRequest(uri, "POST", body, account, passphrase)
}

func SendPutRequest(uri string, body []byte, account string, passphrase string) ([]byte, error) {
	return sendRequest(uri, "PUT", body, account, passphrase)
}

func SendPatchRequest(uri string, body []byte, account string, passphrase string) ([]byte, error) {
	return sendRequest(uri, "PATCH", body, account, passphrase)
}

func sendRequest(uri string, method string, body []byte, account string, passphrase string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, BuildUrl(uri), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if len(account) != 0 && len(passphrase) != 0 {
		req.SetBasicAuth(account, passphrase)
	}

	resp, err := client.Do(req)

	response := ReadResponseBody(resp)
	println(string(response))

	if resp.StatusCode != 200 {
		return nil, sdk.NewError("test", sdk.CodeType(resp.StatusCode), "Error occurred")
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

func ReadResponseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
