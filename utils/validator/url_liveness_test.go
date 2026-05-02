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
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLLivenessCheck(t *testing.T) {
	negativeTests := []string{
		"https://dcl-test.org",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
	}
	positiveTests := []string{
		"http://github.com/",             // Redirects to https://github.com/
		"https://httpbin.org/status/401", // Private repo
		"https://httpbin.org/status/403", // Unavailable for some reason
	}

	for _, testUrl := range negativeTests {
		u, err := url.ParseRequestURI(testUrl)
		require.NoError(t, err)

		require.False(t, _isLiveURL(u))
	}

	for _, testUrl := range positiveTests {
		u, err := url.ParseRequestURI(testUrl)
		require.NoError(t, err)

		require.True(t, _isLiveURL(u))
	}
}
