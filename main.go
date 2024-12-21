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
//    2024-12-20: V1.0.0: Created.
//

package main

import (
	"encoding/base32"
	"encoding/base64"
	"flag"
	"fmt"
	"hash"
	"hashvalue/filehelper"
	"hashvalue/hashimplementation"
	"io"
	"os"
)

func main() {
	os.Exit(realMain())
}

// These are the possible return codes of [realMain].
const (
	rcOK               = 0
	rcParameterError   = 1
	rcpProcessingError = 2
)

// ******** Private variables ********

// haveSource is true if the 'source' option has been set.
var haveSource = false

// haveFile is true if the 'file' option has been set.
var haveFile = false

// ******** Private functions ********

// realMain is the real main function that returns a return code.
func realMain() int {
	// 1. Define command line flags.

	var hashTypeName string
	flag.StringVar(&hashTypeName, `hash`, `sha3-256`, "name of `hash type`")

	var source string
	flag.StringVar(&source, `source`, ``, "Source `text` (mutually exclusive with 'file')")

	var fileName string
	flag.StringVar(&fileName, `file`, ``, "Name of `source file` (mutually exclusive with 'source')")

	var separator string
	flag.StringVar(&separator, `separator`, ``, "Separator `text` between hex bytes")

	var prefix string
	flag.StringVar(&prefix, `prefix`, ``, "Prefix `text` in front of hex bytes")

	var useLower bool
	flag.BoolVar(&useLower, `lower`, false, `Use lower case for hex output (default)`)

	var useUpper bool
	flag.BoolVar(&useUpper, `upper`, false, `Use upper case for hex output`)

	var useBase32 bool
	flag.BoolVar(&useBase32, `base32`, false, `Encode hash in base32 format (combinable with 'hex' and 'base64')'`)

	var useBase64 bool
	flag.BoolVar(&useBase64, `base64`, false, `Encode hash in base64 format (combinable with 'hex' and 'base32')`)

	var useHex bool
	flag.BoolVar(&useHex, `hex`, false, `Encode hash in hex format (default, modifiable with 'separator', 'prefix' and either 'lower' or 'upper', combinable with 'base32' and 'base64')`)

	var noHeader bool
	flag.BoolVar(&noHeader, `noheader`, false, `Do not print the type of the encoding in front of it`)

	flag.Usage = MyUsage

	flag.Parse()

	// 2. Check for command line errors.

	if flag.NArg() > 0 {
		return printUsageErrorf("Arguments without flags present: %s\n", flag.Args())
	}

	flag.Visit(visitOptions)

	if haveSource && haveFile {
		return printUsageError("Do not specify 'source' and 'file'")
	}

	if !(haveSource || haveFile) {
		return printUsageError("Specify either 'source' or 'file'")
	}

	if haveSource && len(source) == 0 {
		return printUsageError(`Source is empty`)
	}

	if haveFile && len(fileName) == 0 {
		return printUsageError(`File name is empty`)
	}

	if useLower && useUpper {
		return printUsageError(`Specify either 'lower' or 'upper'`)
	}

	if !(useLower || useUpper) {
		useLower = true
	}

	if !(useBase32 || useBase64 || useHex) {
		useHex = true
	}

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

	if !noHeader {
		fmt.Print(`Hash  : `)
	}
	fmt.Println(normalizedHashTypeName)

	if useHex {
		if !noHeader {
			fmt.Print(`Hex   : `)
		}
		printHex(hashValue, separator, prefix, useLower)
	}

	if useBase32 {
		if !noHeader {
			fmt.Print(`Base32: `)
		}
		fmt.Println(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashValue))
	}

	if useBase64 {
		if !noHeader {
			fmt.Print(`Base64: `)
		}
		fmt.Println(base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(hashValue))
	}

	return rcOK
}

// MyUsage is the function that is called by flag.Usage. It prints the usage information.
func MyUsage() {
	errWriter := flag.CommandLine.Output()
	_, _ = fmt.Fprintf(errWriter, "\nUse '%s' with the following options:\n\n", filehelper.GetRealBaseName(os.Args[0]))
	flag.PrintDefaults()
	_, _ = fmt.Fprintf(errWriter, "\nValid hash type names: %s\n", hashimplementation.KnownHashNames())
}

// ******** Private function ********

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

// hashData hashes the data in [source] or from file [fileName].
func hashData(hashFunc hash.Hash, source string, fileName string) ([]byte, error) {
	var hashValue []byte
	if haveSource {
		hashFunc.Write([]byte(source))
		hashValue = hashFunc.Sum(nil)
	} else {
		var err error
		hashValue, err = fileHash(hashFunc, fileName)
		if err != nil {
			return nil, fmt.Errorf("error reading file '%s': %w", fileName, err)
		}
	}

	return hashValue, nil
}

// fileHash calculates the hash value of a file.
func fileHash(hashFunc hash.Hash, fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer filehelper.CloseFile(f)

	if _, err = io.Copy(hashFunc, f); err != nil {
		return nil, err
	}

	return hashFunc.Sum(nil), nil
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

// -------- Error functions --------

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

	return rcpProcessingError
}
