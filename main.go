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
// Version: 1.1.1
//
// Change history:
//    2024-12-20: V1.0.0: Created.
//    2024-12-29: V1.1.0: Make flags global and modularize code.
//    2024-12-30: V1.1.1: Show version.
//

package main

import (
	"hashvalue/filehelper"
	"hashvalue/hashimplementation"
	"os"
)

func main() {
	os.Exit(realMain())
}

// ******** Private constants ********

// These are the possible return codes of [realMain].
const (
	rcOK               = 0
	rcParameterError   = 1
	rcpProcessingError = 2
)

// myVersion contains the current version of this program.
const myVersion = `1.1.1`

// myCopyright contains the copyright of this program.
const myCopyright = `Copyright (c) 2024 Frank Schwab`

// ******** Private variables ********

// myName contains the name of the current executable.
var myName string

// Option presence flags.

// haveSource is true if the 'source' option has been set.
var haveSource = false

// haveFile is true if the 'file' option has been set.
var haveFile = false

// Option values.

// They have to be global in order to modularize the main program.
// Otherwise, there would have been an awful lot of parameters to pass to functions.

// hashTypeName is the name of the hash.
var hashTypeName string

// source is the source text to hash.
var source string

// fileName is the name of the file whose contents are to be hashed.
var fileName string

// separator is the separator text for hex output.
var separator string

// prefix is prefix text for hex output.
var prefix string

// useLower indicates that lower-case characters should be used for hex output.
var useLower bool

// useUpper indicates that upper-case characters should be used for hex output.
// This is a helper flag and is not used for execution control.
// It is mapped to useLower.
var useUpper bool

// useBase16 indicates that base16 (hex) encoding should be used for hash output.
// This is a helper flag and is not used for execution control.
// It is mapped to useHex.
var useBase16 bool

// useBase32 indicates that base32 encoding should be used for hash output.
var useBase32 bool

// useBase64 indicates that base64 encoding should be used for hash output.
var useBase64 bool

// useHex indicates that hex encoding should be used for hash output.
var useHex bool

// noHeaders indicates that output should not be prefixed by a header.
var noHeaders bool

// showVersion indicates that the version information should be printed.
var showVersion bool

// ******** Private functions ********

// realMain is the real main function that returns a return code.
func realMain() int {
	myName = filehelper.GetRealBaseName(os.Args[0])

	// 1. Define command line flags.
	defineCommandLineFlags()

	// Show version and exit if version is requested.
	if showVersion {
		printVersion()
		return rcOK
	}

	// 2. Check for command line errors.
	rc := checkCommandLineFlags()
	if rc != rcOK {
		return rc
	}

	// 3. Normalize command line flags.
	normalizeCommandLineFlags()

	// 4. Get hash function.
	normalizedHashTypeName, hashFunc, ok := hashimplementation.NewHashFunctionOfType(hashTypeName)
	if !ok {
		return printUsageErrorf(`Invalid hash type: '%s'`, hashTypeName)
	}

	// 3. Hash data.
	hashValue, err := hashData(hashFunc, source, fileName)
	if err != nil {
		return printErrorf(`Error hashing data: %s`, err)
	}

	// 4. Print result.
	printResult(normalizedHashTypeName, hashValue)

	return rcOK
}
