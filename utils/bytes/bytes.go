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

package bytes

func cp(bz []byte) (ret []byte) {
	ret = make([]byte, len(bz))
	copy(ret, bz)

	return ret
}

// Returns a slice of the same length (big endian)
// except incremented by one.
// Appends 0x00 if bz is all 0xFF.
// CONTRACT: len(bz) > 0
//
// Taken from: "github.com/tendermint/iavl@v0.12.4/util.go".
func CpIncr(bz []byte) (ret []byte) {
	ret = cp(bz)
	for i := len(bz) - 1; i >= 0; i-- {
		if ret[i] < byte(0xFF) {
			ret[i]++

			return
		}

		ret[i] = byte(0x00)

		if i == 0 {
			return append(ret, 0x00)
		}
	}

	return []byte{0x00}
}
