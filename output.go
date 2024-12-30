//
// SPDX-FileCopyrightText: Copyright 2024 Frank Schwab
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
// Version: 1.1.0
//
// Change history:
//    2024-12-29: V1.0.0: Created.
//    2024-12-30: V1.1.0: Print hex bytes directly and not via fmt.Printf.
//

package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"os"
)

// ******** Private constants ********

// lowerOffset is the offset between lower and upper case characters.
const lowerOffset byte = 'a' - 'A'

// characterOffset is the offset between a digit character and an alphabetical character.
const characterOffset byte = 'A' - '0' - 10

// ******** Private variables ********

// hexCharBuffer is the one byte slice that holds the byte to print as a hex character.
var hexCharBuffer = make([]byte, 1)

// ******** Private functions ********

// printResult prints the hash value.
func printResult(normalizedHashTypeName string, hashValue []byte) {
	if !noHeaders {
		fmt.Print(`Hash  : `)
	}
	fmt.Println(normalizedHashTypeName)

	if useHex {
		if !noHeaders {
			fmt.Print(`Hex   : `)
		}
		printHex(hashValue, separator, prefix, useLower)
	}

	if useBase32 {
		if !noHeaders {
			fmt.Print(`Base32: `)
		}
		fmt.Println(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashValue))
	}

	if useBase64 {
		if !noHeaders {
			fmt.Print(`Base64: `)
		}
		fmt.Println(base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(hashValue))
	}
}

// printHex prints a byte array in hex format where the bytes are separated
// by [separator] and prefixed by [prefix]. The byte values are printed
// either with lower or upper case characters.
func printHex(hashValue []byte, separator string, prefix string, useLower bool) {
	caseOffset := characterOffset
	if useLower {
		caseOffset += lowerOffset
	}

	separatorBytes := []byte(separator)
	prefixBytes := []byte(prefix)

	useSeparator := false
	usePrefix := len(prefix) != 0
	for _, b := range hashValue {
		if useSeparator {
			_, _ = os.Stdout.Write(separatorBytes)
		} else {
			useSeparator = true
		}

		if usePrefix {
			_, _ = os.Stdout.Write(prefixBytes)
		}

		printHexByte(b, caseOffset)
	}

	fmt.Println()
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
