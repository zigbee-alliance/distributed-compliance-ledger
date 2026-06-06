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

package cli

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/zigbee-alliance/distributed-compliance-ledger/internal/config"
)

const (
	livenessCheckTimeout = 10 * time.Second
)

var allowed4XXStatusCodes = []int{
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusUnavailableForLegalReasons,
	http.StatusMethodNotAllowed,
}
var httpClient = &http.Client{Timeout: livenessCheckTimeout}

func IsLiveURL(u string) bool {
	if !config.URLLivenessCheckEnabled {
		return true
	}

	ctx, cancel := context.WithTimeout(context.Background(), livenessCheckTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, u, nil)
	if err != nil {
		return false
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return false
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		return true
	}

	for _, code := range allowed4XXStatusCodes {
		if code == resp.StatusCode {
			return true
		}
	}

	return false
}

// CheckURLsForLiveness checks the liveness of the given URLs concurrently and
// returns unreachable URLs as a list.
// Empty strings are skipped.
//
// Returns an empty list if all non-empty URLs are reachable.
func CheckURLsForLiveness(urls ...string) []string {
	results := make([]string, len(urls))

	var wg sync.WaitGroup
	for i, u := range urls {
		if u == "" {
			continue
		}
		// Call each URL concurrently
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()
			if !IsLiveURL(u) {
				results[i] = u
			}
		}(i, u)
	}
	wg.Wait()

	var unreachable []string
	for _, u := range results {
		if u != "" {
			unreachable = append(unreachable, u)
		}
	}

	return unreachable
}
