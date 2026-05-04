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

//go:build !dev

package validator

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLiveURL(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{"200 OK", http.StatusOK, true},
		{"301 redirect", http.StatusMovedPermanently, true},
		{"401 unauthorized", http.StatusUnauthorized, true},
		{"403 forbidden", http.StatusForbidden, true},
		{"451 unavailable for legal reasons", http.StatusUnavailableForLegalReasons, true},
		{"404 not found", http.StatusNotFound, false},
		{"500 internal server error", http.StatusInternalServerError, false},
		{"502 bad gateway", http.StatusBadGateway, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodHead, r.Method)
				w.WriteHeader(tt.statusCode)
			}))
			defer srv.Close()

			u, err := url.ParseRequestURI(srv.URL)
			require.NoError(t, err)

			require.Equal(t, tt.want, isLiveURL(u))
		})
	}
}

func TestIsLiveURLUnreachable(t *testing.T) {
	u, err := url.ParseRequestURI("http://192.0.2.1:1")
	require.NoError(t, err)

	require.False(t, isLiveURL(u))
}
