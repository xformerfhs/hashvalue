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
// Version: 1.0.0
//
// Change history:
//    2024-12-29: V1.0.0: Created.
//

package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

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
	var hexFormat string
	if useLower {
		hexFormat = `%02x`
	} else {
		hexFormat = `%02X`
	}

	useSeparator := false
	usePrefix := len(prefix) != 0
	for _, b := range hashValue {
		if useSeparator {
			fmt.Print(separator)
		} else {
			useSeparator = true
		}

		if usePrefix {
			fmt.Print(prefix)
		}
		fmt.Printf(hexFormat, b)
	}

	fmt.Println()
}
