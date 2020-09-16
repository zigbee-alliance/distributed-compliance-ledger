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

package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandString() string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var b strings.Builder

	for i := 0; i < 5; i++ {
		// nolint:gosec
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func ParseUint(num string) uint64 {
	val, _ := strconv.ParseUint(num, 10, 64)

	return val
}
