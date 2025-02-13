//
// SPDX-FileCopyrightText: Copyright 2025 Frank Schwab
//
// SPDX-License-Identifier: Apache-2.0
//
// SPDX-FileType: SOURCE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: Frank Schwab
//
// Version: 2.3.0
//
// Change history:
//    2025-01-31: V1.0.0: Created.
//    2025-02-01: V2.0.0: Use separate package name.
//    2025-02-01: V2.1.0: Better test cases for padded encoding.
//    2025-02-02: V2.2.0: Test structured errors.
//    2025-02-10: V2.3.0: Correct for new type of invalid byte error, fill buffers with random data.
//

package z85_test

import (
	"errors"
	"hashvalue/z85"
	"testing"
)

// ******** Private constants ********

// clearTheOne contains the clear bytes of the one test case on the https://rfc.zeromq.org/spec/32 website.
var clearTheOne = []byte{0x86, 0x4f, 0xd2, 0x6f, 0xb5, 0x59, 0xf7, 0x5B}

// encodedTheOne contains the encoded string of the one test case on the https://rfc.zeromq.org/spec/32 website.
var encodedTheOne = `HelloWorld`

// ******** Test functions ********

// TestEncodeTheOne implements the one test case documented on the https://rfc.zeromq.org/spec/32 website.
func TestEncodeTheOne(t *testing.T) {
	encoded, err := z85.Encode(clearTheOne)
	if err != nil {
		t.Fatalf(`Encoding failed: %v`, err)
	}

	if encoded != encodedTheOne {
		t.Fatalf(`Encoding did not result in '%s', but '%s'`, encodedTheOne, encoded)
	}
}

// TestNilEncode tests the encoding of an empty byte slice.
func TestNilEncode(t *testing.T) {
	encoded, err := z85.Encode(nil)
	if err != nil {
		t.Fatalf(`Encoding failed: %v`, err)
	}

	if len(encoded) != 0 {
		t.Fatalf(`Encoding nil did not result in an empty string, but '%s'`, encoded)
	}
}

// TestEncodeWithInvalidLength tests if an error occurs encoding with an invalid length.
func TestEncodeWithInvalidLength(t *testing.T) {
	_, err := z85.Encode(clearTheOne[2:5])
	if err == nil {
		t.Fatal(`Invalid length did not result in an error`)
	} else {
		var expectedErr z85.ErrInvalidLength
		ok := errors.As(err, &expectedErr)
		if !ok {
			t.Fatalf(`Wrong error when encoding invalid length string: '%v'`, err)
		}
	}
}
