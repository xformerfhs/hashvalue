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
//    2025-03-02: V1.0.0: Created.
//

package encodedprinting

import (
	"hashvalue/stringhelper"
	"os"
)

// HexEncoder is used to encode bytes in hex encoding.
type HexEncoder struct {
	separator  []byte
	prefix     []byte
	caseOffset byte
}

// ******** Private constants ********

// lowerOffset is the offset between lower and upper case characters.
const lowerOffset byte = 'a' - 'A'

// characterOffset is the offset between a digit character and an alphabetical character.
const characterOffset byte = 'A' - '9' - 1

// hexCharBuffer is the one byte slice that holds the byte to print as a hex character.
var hexCharBuffer = make([]byte, 1)

// NewHexEncoder creates a new hexadecimal encoder.
func NewHexEncoder(separator string, prefix string, useLower bool) *HexEncoder {
	caseOffset := characterOffset
	if useLower {
		caseOffset += lowerOffset
	}

	return &HexEncoder{
		separator:  stringhelper.UnsafeStringBytes(separator),
		prefix:     stringhelper.UnsafeStringBytes(prefix),
		caseOffset: caseOffset,
	}
}

// PrintEncoded prints a byte array in hex format where the bytes are separated
// by separator and prefixed by prefix. The byte values are printed either with
// lower or upper case characters, depending on useLower.
func (e *HexEncoder) PrintEncoded(hashValue []byte) {
	out := os.Stdout

	useSeparator := false
	usePrefix := len(e.prefix) != 0
	for _, b := range hashValue {
		if useSeparator {
			_, _ = out.Write(e.separator)
		} else {
			useSeparator = true
		}

		if usePrefix {
			_, _ = out.Write(e.prefix)
		}

		printHexByte(b, e.caseOffset)
	}

	_, _ = out.Write(newLine)
}

// printHexByte prints one byte in hexadecimal (base16) encoding.
func printHexByte(b byte, caseOffset byte) {
	// Print upper nibble.
	i := b >> 4
	printHexChar(i, caseOffset)

	// Print lower nibble.
	i = b & 0x0f
	printHexChar(i, caseOffset)
}

// printHexChar prints one hex character.
func printHexChar(b byte, caseOffset byte) {
	// 1. Convert to byte starting at '0'.
	c := b + '0'

	// 2. If the value is above 9, add the correct offset that starts the byte at 'A' or 'a'.
	if b > 9 {
		c += caseOffset
	}

	// 3. The "Write" function needs a byte slice. So copy character byte to byte slice.
	hexCharBuffer[0] = c

	// 4. Write the byte to the Stdout writer.
	_, _ = os.Stdout.Write(hexCharBuffer)
}
