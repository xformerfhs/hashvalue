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
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
)

// ******** Private functions ********

// printUsageError prints an error message and the usage information.
func printUsageError(msg string) int {
	errWriter := flag.CommandLine.Output()

	_, _ = fmt.Fprintln(errWriter)
	_, _ = fmt.Fprint(errWriter, msg)

	return printUsage(errWriter)
}

// printUsageErrorf print a formatted error message and the usage information.
func printUsageErrorf(format string, a ...interface{}) int {
	errWriter := flag.CommandLine.Output()

	_, _ = fmt.Fprintln(errWriter)
	_, _ = fmt.Fprintf(errWriter, format, a...)

	return printUsage(errWriter)
}

// printUsage prints the usage.
func printUsage(errWriter io.Writer) int {
	_, _ = fmt.Fprintln(errWriter)
	flag.Usage()

	return rcParameterError
}

// printErrorf prints a processing error message.
func printErrorf(format string, a ...interface{}) int {
	_, _ = fmt.Fprintln(os.Stderr)
	_, _ = fmt.Fprintf(os.Stderr, format, a...)

	return rcProcessingError
}

// printVersion prints the version information for this program.
func printVersion() {
	fmt.Printf("\n%s V%s (%s), %s\n", myName, myVersion, runtime.Version(), myCopyright)
}
