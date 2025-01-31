//
// SPDX-FileCopyrightText: Copyright 2024-2025 Frank Schwab
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
// Version: 1.3.0
//
// Change history:
//    2024-12-29: V1.0.0: Created.
//    2024-12-30: V1.1.0: Print hex bytes directly and not via fmt.Printf.
//    2025-01-28: V1.2.0: Get rid of "fmt" package.
//    2025-01-31: V1.3.0: Add Z85 encoding of output.
//

package main

import (
	"encoding/base32"
	"encoding/base64"
	"hashvalue/z85"
	"os"
)

// ******** Private constants ********

// lowerOffset is the offset between lower and upper case characters.
const lowerOffset byte = 'a' - 'A'

// characterOffset is the offset between a digit character and an alphabetical character.
const characterOffset byte = 'A' - '9' - 1

// newLine contains a byte slice with the newline character.
var newLine = []byte{'\n'}

// ******** Private variables ********

// hexCharBuffer is the one byte slice that holds the byte to print as a hex character.
var hexCharBuffer = make([]byte, 1)

// ******** Private functions ********

// printResult prints the hash value.
func printResult(normalizedHashTypeName string, hashValue []byte) {
	out := os.Stdout
	if !noHeaders {
		_, _ = out.WriteString(`Hash  : `)
	}
	writeStringln(out, normalizedHashTypeName)

	if useHex {
		if !noHeaders {
			_, _ = out.WriteString(`Hex   : `)
		}
		printHex(hashValue, separator, prefix, useLower)
	}

	if useBase32 {
		if !noHeaders {
			_, _ = out.WriteString(`Base32: `)
		}
		writeStringln(out, base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashValue))
	}

	if useBase64 {
		if !noHeaders {
			_, _ = out.WriteString(`Base64: `)
		}
		writeStringln(out, base64.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashValue))
	}

	if useZ85 {
		if !noHeaders {
			_, _ = out.WriteString(`Z85   : `)
		}
		encoded, _ := z85.Encode(hashValue)
		writeStringln(out, encoded)
	}
}

// printHex prints a byte array in hex format where the bytes are separated
// by separator and prefixed by prefix. The byte values are printed either with
// lower or upper case characters, depending on useLower.
func printHex(hashValue []byte, separator string, prefix string, useLower bool) {
	caseOffset := characterOffset
	if useLower {
		caseOffset += lowerOffset
	}

	separatorBytes := []byte(separator)
	prefixBytes := []byte(prefix)
	out := os.Stdout

	useSeparator := false
	usePrefix := len(prefix) != 0
	for _, b := range hashValue {
		if useSeparator {
			_, _ = out.Write(separatorBytes)
		} else {
			useSeparator = true
		}

		if usePrefix {
			_, _ = out.Write(prefixBytes)
		}

		printHexByte(b, caseOffset)
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

// writeStringln writes a string followed by a newline character.
func writeStringln(out *os.File, s string) {
	_, _ = out.WriteString(s)
	_, _ = out.Write(newLine)
}
