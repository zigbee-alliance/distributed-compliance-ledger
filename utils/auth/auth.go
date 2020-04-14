package auth

import (
	"encoding/base64"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"net/http"
	"strings"
)

func GetCredentialsFromRequest(r *http.Request) (string, string, error) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return "", "", sdk.ErrUnauthorized("authorization failed")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		return "", "", sdk.ErrUnauthorized("authorization failed")
	}

	return pair[0], pair[1], nil
}
