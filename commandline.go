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
	"hashvalue/hashimplementation"
)

// ******** Private constants ********

// parameterTooLongErrorFormat is the format to use for "parameter too long" errors.
const parameterTooLongErrorFormat = `%s is too long`

// maxHexParameterLen is the maximum length for a hex formatting parameter.
const maxHexParameterLen = 8

// ******** Private variables ********

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

// defineCommandLineFlags defines the command line flags.
func defineCommandLineFlags() {
	flag.StringVar(&hashTypeName, `hash`, `sha3-256`, "name of hash `algorithm`")
	flag.StringVar(&source, `source`, ``, "Source `text` (mutually exclusive with 'file')")
	flag.StringVar(&fileName, `file`, ``, "Source file `path` (mutually exclusive with 'source')")
	flag.StringVar(&separator, `separator`, ``, "Separator `text` between hex bytes")
	flag.StringVar(&prefix, `prefix`, ``, "Prefix `text` in front of hex bytes")
	flag.BoolVar(&useLower, `lower`, false, `Use lower case for hex output`)
	flag.BoolVar(&useUpper, `upper`, false, `Use upper case for hex output (default)`)
	flag.BoolVar(&useBase16, `base16`, false, `Encode hash in base16 (hex) format (alias for 'hex')'`)
	flag.BoolVar(&useBase32, `base32`, false, `Encode hash in base32 format (combinable with 'hex' and 'base64')'`)
	flag.BoolVar(&useBase64, `base64`, false, `Encode hash in base64 format (combinable with 'hex' and 'base32')`)
	flag.BoolVar(&useHex, `hex`, false, `Encode hash in hex (base16) format (default, modifiable with 'separator', 'prefix' and either 'lower' or 'upper', combinable with 'base32' and 'base64')`)
	flag.BoolVar(&noHeaders, `noheaders`, false, `Do not print the type of the output in front of it`)
	flag.BoolVar(&showVersion, `version`, false, `Show program version and exit`)

	flag.Usage = myUsage

	flag.Parse()
}

// myUsage is the function that is called by flag.Usage. It prints the usage information.
func myUsage() {
	errWriter := flag.CommandLine.Output()
	_, _ = fmt.Fprintf(errWriter, "\nUse '%s' with the following options:\n\n", myName)
	flag.PrintDefaults()
	_, _ = fmt.Fprintf(errWriter, "\nValid hash type names: %s\n", hashimplementation.KnownHashNames())
}

// checkCommandLineFlags checks the command line flags.
func checkCommandLineFlags() int {
	if flag.NArg() > 0 {
		return printUsageErrorf(`Arguments without flags present: %s`, flag.Args())
	}

	flag.Visit(visitOptions)

	if haveSource && haveFile {
		return printUsageError(`Do not specify 'source' and 'file'`)
	}

	if !(haveSource || haveFile) {
		return printUsageError(`Specify either 'source' or 'file'`)
	}

	if haveSource && len(source) == 0 {
		return printUsageError(`Source is empty`)
	}

	if haveFile && len(fileName) == 0 {
		return printUsageError(`File name is empty`)
	}

	if len(separator) > maxHexParameterLen {
		return printUsageErrorf(parameterTooLongErrorFormat, `separator`)
	}

	if len(prefix) > maxHexParameterLen {
		return printUsageErrorf(parameterTooLongErrorFormat, `prefix`)
	}

	if useLower && useUpper {
		return printUsageError(`Specify either 'lower' or 'upper'`)
	}

	return rcOK
}

// visitOptions is the visitor function that checks which options have been set.
func visitOptions(f *flag.Flag) {
	switch f.Name {
	case `source`:
		haveSource = true

	case `file`:
		haveFile = true
	}
}

// normalizeCommandLineFlags normalizes the command line flags.
func normalizeCommandLineFlags() {
	if !(useLower || useUpper) {
		useUpper = true
	}

	if useBase16 {
		useHex = true
	}

	if !(useBase32 || useBase64 || useHex) {
		useHex = true
	}
}
