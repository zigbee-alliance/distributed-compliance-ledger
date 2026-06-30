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

package cli

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const unreachableURL = "http://192.0.2.1:1"

func TestIsLiveURL(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{"200 OK", http.StatusOK, true},
		{"201 created", http.StatusCreated, true},
		{"204 no content", http.StatusNoContent, true},
		{"299 highest 2xx", 299, true},
		{"301 redirect", http.StatusMovedPermanently, true},
		{"302 found", http.StatusFound, true},
		{"308 permanent redirect", http.StatusPermanentRedirect, true},
		{"399 highest 3xx", 399, true},
		{"401 unauthorized", http.StatusUnauthorized, true},
		{"403 forbidden", http.StatusForbidden, true},
		{"405 method not allowed", http.StatusMethodNotAllowed, true},
		{"451 unavailable for legal reasons", http.StatusUnavailableForLegalReasons, true},
		{"400 bad request", http.StatusBadRequest, false},
		{"404 not found", http.StatusNotFound, false},
		{"429 too many requests", http.StatusTooManyRequests, false},
		{"500 internal server error", http.StatusInternalServerError, false},
		{"502 bad gateway", http.StatusBadGateway, false},
		{"503 service unavailable", http.StatusServiceUnavailable, false},
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

			require.Equal(t, tt.want, IsLiveURL(u.String()))
		})
	}
}

func TestIsLiveURLUnreachable(t *testing.T) {
	u, err := url.ParseRequestURI(unreachableURL)
	require.NoError(t, err)

	require.False(t, IsLiveURL(u.String()))
}

// TestIsLiveURLRequestCreationError covers the path where the HTTP request
// cannot even be constructed (e.g. the URL is malformed), so no network call
// is attempted and the URL is treated as not live.
func TestIsLiveURLRequestCreationError(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"invalid control character", "http://example.com/\x7f"},
		{"missing protocol scheme", ":no-scheme"},
		{"control character in host", "http://exa\x00mple.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.False(t, IsLiveURL(tt.url))
		})
	}
}

// TestIsLiveURLAllowed4XXNotMutated guards against the allow-list of 4XX status
// codes being changed unexpectedly, since IsLiveURL relies on it for the
// allowed-status path.
func TestIsLiveURLAllowed4XXNotMutated(t *testing.T) {
	require.ElementsMatch(t, []int{
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusUnavailableForLegalReasons,
		http.StatusMethodNotAllowed,
	}, allowed4XXStatusCodes)
}

func TestCheckURLsForLiveness(t *testing.T) {
	okSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer okSrv.Close()

	notFoundSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer notFoundSrv.Close()

	tests := []struct {
		name string
		urls []string
		want []string
	}{
		{"no URLs", nil, nil},
		{"all empty strings", []string{"", "", ""}, nil},
		{"all reachable", []string{okSrv.URL, okSrv.URL}, nil},
		{"single unreachable", []string{okSrv.URL, notFoundSrv.URL}, []string{notFoundSrv.URL}},
		{"empties skipped", []string{"", okSrv.URL, ""}, nil},
		{
			"multiple unreachable preserve input order",
			[]string{okSrv.URL, notFoundSrv.URL, unreachableURL},
			[]string{notFoundSrv.URL, unreachableURL},
		},
		{
			"unreachable later in list",
			[]string{okSrv.URL, "", notFoundSrv.URL},
			[]string{notFoundSrv.URL},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, CheckURLsForLiveness(tt.urls...))
		})
	}
}

func TestCheckURLsForLivenessRunsConcurrently(t *testing.T) {
	const handlerDelay = 200 * time.Millisecond
	const concurrentURLs = 5

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(handlerDelay)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	urls := make([]string, concurrentURLs)
	for i := range urls {
		urls[i] = srv.URL
	}

	start := time.Now()
	require.Empty(t, CheckURLsForLiveness(urls...))
	elapsed := time.Since(start)

	// Sequential calls would take approximately concurrentURLs*handlerDelay time
	// Concurrent execution should finish in roughly handlerDelay.
	require.Less(t, elapsed, time.Duration(concurrentURLs)*handlerDelay/2,
		"URL checks did not run concurrently (elapsed %s)", elapsed)
}
