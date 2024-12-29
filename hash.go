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
	"fmt"
	"hash"
	"hashvalue/filehelper"
	"io"
	"os"
)

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
