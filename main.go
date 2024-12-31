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
// Version: 1.1.2
//
// Change history:
//    2024-12-20: V1.0.0: Created.
//    2024-12-29: V1.1.0: Make flags global and modularize code.
//    2024-12-30: V1.1.1: Show version.
//    2024-12-30: V1.1.2: Simplify hex byte output.
//

package main

import (
	"hashvalue/filehelper"
	"hashvalue/hashimplementation"
	"os"
)

// main is entry point of the program.
// It is a stub that calls the real main function which has an exit code.
func main() {
	os.Exit(realMain())
}

// ******** Private constants ********

// myVersion contains the current version of this program.
const myVersion = `1.1.2`

// myCopyright contains the copyright of this program.
const myCopyright = `Copyright (c) 2024 Frank Schwab`

// ******** Private variables ********

// myName contains the name of the current executable.
var myName string

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
