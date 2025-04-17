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
// Version: 3.1.0
//
// Change history:
//    2024-12-29: V1.0.0: Created.
//    2025-01-31: V1.1.0: Add Z85 encoding of output.
//    2025-02-26: V2.0.0: No more headers. Allow only one encoding.
//    2025-03-02: V3.0.0: New command line structure. Ability to process hex bytes.
//    2025-04-17: V3.1.0: Change "hash type" to "hash algorithm". No default hash algorithm.
//

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"hashvalue/encodedprinting"
	"hashvalue/hashfactory"
	"hashvalue/stringhelper"
	"strings"
)

// ******** Private constants ********

// parameterTooLongErrorFormat is the format to use for "parameter too long" errors.
const parameterTooLongErrorFormat = `%s is too long`

// maxHexParameterLen is the maximum length for a hex formatting parameter.
const maxHexParameterLen = 8

// errFmtIsEmpty is the error string for an empty variable.
const errFmtIsEmpty = `%s is empty`

// ******** Private variables ********

// Option presence flags.

// haveSource is true if the 'source' option has been set.
var haveSource = false

// haveHexSource is true if the 'hexsource' option has been set.
var haveHexSource = false

// haveFile is true if the 'file' option has been set.
var haveFile = false

// Option values.

// They have to be global in order to modularize the main program.
// Otherwise, there would have been an awful lot of parameters to pass to functions.

// hashAlgorithm is the name of the hash.
var hashAlgorithm string

// source is the source text to hash.
var source string

// hexSource is the source text to hash in hex encoding.
var hexSource string

// fileName is the name of the file whose contents are to be hashed.
var fileName string

// encodingType specifies the output encoding to use.
var encodingType string

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

// showVersion indicates that the version information should be printed.
var showVersion bool

// sourceBytes contains the bytes of the source.
var sourceBytes []byte

// ******** Private functions ********

// parseCommandLineWithFlags defines the command line flags and parses the command line.
func parseCommandLineWithFlags() {
	// 1. Define flags.
	flag.StringVar(&hashAlgorithm, `hash`, ``, "name of hash `algorithm`")
	flag.StringVar(&source, `source`, ``, "Source `text` (mutually exclusive with 'hexsource' and 'file')")
	flag.StringVar(&hexSource, `hexsource`, ``, "Hexadecimal source `text` (mutually exclusive with 'source' and 'file')")
	flag.StringVar(&fileName, `file`, ``, "Source file `path` (mutually exclusive with 'source' and 'hexsource')")
	flag.StringVar(&encodingType, `encoding`, `hex`, "Encoding `type` of hash value (one of 'hex', 'base16', 'base32', 'base64', or 'z85')")
	flag.StringVar(&separator, `separator`, ``, "Separator `text` between hex bytes")
	flag.StringVar(&prefix, `prefix`, ``, "Prefix `text` in front of hex bytes")
	flag.BoolVar(&showVersion, `version`, false, `Show program version and exit`)
	flag.BoolVar(&useLower, `lower`, false, `Use lower case for hex output`)
	flag.BoolVar(&useUpper, `upper`, false, `Use upper case for hex output (default)`)

	// 2. Set usage function.
	flag.Usage = myUsage

	// 3. Parse command line.
	flag.Parse()
}

// myUsage is the function called by flag.Usage. It prints the usage information.
func myUsage() {
	errWriter := flag.CommandLine.Output()
	_, _ = fmt.Fprintf(errWriter, "\nUse '%s' with the following options:\n\n", myName)
	flag.PrintDefaults()
	_, _ = fmt.Fprintln(errWriter, "\nSpecify only one encoding.")
	_, _ = fmt.Fprintf(errWriter, "\nValid hash algorithm names: %s\n", hashfactory.KnownHashNames())
}

// normalizeCommandLineFlags normalizes the command line flags.
func normalizeCommandLineFlags() {
	// Normalize encoding type.
	if len(encodingType) > 0 {
		encodingType = strings.ToLower(strings.TrimSpace(encodingType))
	}

	if len(encodingType) == 0 || encodingType == `base16` {
		encodingType = `hex`
	}

	// Normalize hex source.
	if len(hexSource) > 0 {
		hexSource = stringhelper.RemoveAllWhitespace(hexSource)
	}

	// Normalize hash algorithm name.
	if len(hashAlgorithm) > 0 {
		hashAlgorithm = strings.ToLower(strings.TrimSpace(hashAlgorithm))
	}

	// File name is *not* normalized as a file name may end or start with blanks.

	// Separator and prefix are not normalized as they are always processed as they are.
}

// checkCommandLineFlags checks the command line flags.
func checkCommandLineFlags() (encodedprinting.EncodedPrinter, int) {
	if flag.NArg() > 0 {
		return nil, printUsageErrorf(`Arguments without flags present: %s`, flag.Args())
	}

	if len(hashAlgorithm) == 0 {
		return nil, printUsageError(`No hash algorithm specified`)
	}

	flag.Visit(visitOptions)

	numSources := countTrues(haveSource, haveHexSource, haveFile)

	if numSources == 0 {
		return nil, printUsageError(`Specify either 'source', 'hexsource' or 'file'`)
	}

	if numSources > 1 {
		return nil, printUsageError(`Specify only one of 'source', 'hexsource' or 'file'`)
	}

	if haveSource {
		if len(source) != 0 {
			sourceBytes = stringhelper.UnsafeStringBytes(source)
		} else {
			return nil, printUsageErrorf(errFmtIsEmpty, `Source`)
		}
	}

	if haveHexSource {
		if len(hexSource) != 0 {
			var err error
			sourceBytes, err = hex.DecodeString(hexSource)

			if err != nil {
				return nil, printUsageErrorf(`Invalid hex string: %v`, err)
			}
		} else {
			return nil, printUsageErrorf(errFmtIsEmpty, `Hex source`)
		}
	}

	if haveFile && len(fileName) == 0 {
		return nil, printUsageErrorf(errFmtIsEmpty, `File name`)
	}

	encodedPrinter, isValid := encodingTypeToPrinter(encodingType)
	if !isValid {
		return nil, printUsageErrorf(`Invalid encoding type '%s'`, encodingType)
	}

	if len(separator) > maxHexParameterLen {
		return nil, printUsageErrorf(parameterTooLongErrorFormat, `separator`)
	}

	if len(prefix) > maxHexParameterLen {
		return nil, printUsageErrorf(parameterTooLongErrorFormat, `prefix`)
	}

	if useLower && useUpper {
		return nil, printUsageError(`Specify either 'lower' or 'upper'`)
	}

	return encodedPrinter, rcOK
}

// visitOptions is the visitor function that checks which options have been set.
func visitOptions(f *flag.Flag) {
	switch f.Name {
	case `source`:
		haveSource = true

	case `hexsource`:
		haveHexSource = true

	case `file`:
		haveFile = true
	}
}

// countTrues counts the number of arguments that have a value of "true".
func countTrues(v ...bool) int {
	result := 0
	for _, b := range v {
		if b {
			result++
		}
	}

	return result
}

// encodingTypeToPrinter converts the encoding type to a printer.
func encodingTypeToPrinter(encodingType string) (encodedprinting.EncodedPrinter, bool) {
	switch encodingType {
	case `hex`, `base16`:
		return encodedprinting.NewHexEncoder(separator, prefix, useLower), true

	case `base32`:
		return encodedprinting.NewBase32Encoder(), true

	case `base64`:
		return encodedprinting.NewBase64Encoder(), true

	case `z85`:
		return encodedprinting.NewZ85Encoder(), true

	default:
		return nil, false
	}
}
