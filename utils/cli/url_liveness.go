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
}
var httpClient = &http.Client{Timeout: livenessCheckTimeout}

func IsLiveURL(u string) bool {
	if config.DisableURLLivenessCheck {
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
	defer resp.Body.Close()

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

// FirstUnreachableURL checks the liveness of the given URLs concurrently and
// returns the first one that is not reachable.
// Empty strings are skipped.
//
// Returns an empty string if all non-empty URLs are reachable.
func FirstUnreachableURL(urls ...string) string {
	unreachable := make([]string, len(urls))

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
				unreachable[i] = u
			}
		}(i, u)
	}
	wg.Wait()

	for _, u := range unreachable {
		if u != "" {
			return u
		}
	}

	return ""
}
