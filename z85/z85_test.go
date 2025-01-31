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
// Version: 1.0.0
//
// Change history:
//    2025-01-31: V1.0.0: Created.
//

package z85

import "testing"

// TestEncodeTheOne implements the one test case documented on the https://rfc.zeromq.org/spec/32 website.
func TestEncodeTheOne(t *testing.T) {
	source := []byte{0x86, 0x4f, 0xd2, 0x6f, 0xb5, 0x59, 0xf7, 0x5B}

	encoded, err := Encode(source)
	if err != nil {
		t.Fatalf(`Encoding failed: %v`, err)
	}

	if encoded != `HelloWorld` {
		t.Fatalf(`Encoding did not result in 'HelloWorld', but '%s'`, encoded)
	}
}

func TestNilEncode(t *testing.T) {
	encoded, err := Encode(nil)
	if err != nil {
		t.Fatalf(`Encoding failed: %v`, err)
	}

	if len(encoded) != 0 {
		t.Fatalf(`Encoding nil did not result in an empty string, but '%s'`, encoded)
	}
}

func TestInvalidLength(t *testing.T) {
	source := []byte{0x11, 0x22, 0x33}
	_, err := Encode(source)
	if err == nil {
		t.Fatal(`Invalid length did not result in an error`)
	}
}
