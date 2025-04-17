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
// Version: 4.0.0
//
// Change history:
//    2024-12-20: V1.0.0: Created.
//    2024-12-29: V1.1.0: Make flags global and modularize code.
//    2024-12-30: V1.1.1: Show version.
//    2024-12-30: V1.1.2: Simplify hex byte output.
//    2025-01-28: V1.2.0: Write output directly without the "fmt" package.
//    2025-01-28: V1.3.0: Remove hash "blake2s-128" as it needs a key.
//    2025-01-31: V1.4.0: Add Z85 encoding of output.
//    2025-02-05: V1.4.1: Simpler Z85 size calculation.
//    2025-02-15: V1.4.2: Use z85 package from GitHub.
//    2025-02-26: V2.0.0: Just print the value in one encoding. No headers. No multiple encodings.
//    2025-03-02: V3.0.0: New command line structure. Ability to specify hex bytes.
//    2025-04-17: V4.0.0: No default hash algorithm.
//

package main

import (
	"hashvalue/filehelper"
	"hashvalue/hashfactory"
	"os"
)

// main is entry point of the program.
// It is a stub that calls the real main function which has an exit code.
func main() {
	os.Exit(realMain())
}

// ******** Private constants ********

// myVersion contains the current version of this program.
const myVersion = `4.0.0`

// myCopyright contains the copyright of this program.
const myCopyright = `Copyright (c) 2024-2025 Frank Schwab`

// ******** Private variables ********

// myName contains the name of the current executable.
var myName string

// ******** Private functions ********

// realMain is the real main function that returns a return code.
func realMain() int {
	myName = filehelper.GetRealBaseName(os.Args[0])

	// 1. Define command line flags.
	parseCommandLineWithFlags()

	// Show version and exit if version is requested.
	if showVersion {
		printVersion()
		return rcOK
	}

	// 2. Normalize command line flags.
	normalizeCommandLineFlags()

	// 3. Check for command line errors.
	encodedPrinter, rc := checkCommandLineFlags()
	if rc != rcOK {
		return rc
	}

	// 4. Get hash function.
	hashFunc, ok := hashfactory.New(hashAlgorithm)
	if !ok {
		return printUsageErrorf(`Invalid hash algorithm: '%s'`, hashAlgorithm)
	}

	// 3. Hash data.
	hashValue, err := hashData(hashFunc, sourceBytes, fileName)
	if err != nil {
		return printErrorf(`Error hashing data: %s`, err)
	}

	// 4. Print result.
	encodedPrinter.PrintEncoded(hashValue)

	return rcOK
}
