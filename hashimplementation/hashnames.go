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

// Package hashimplementation implements the interface to the hash functions.
package hashimplementation

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
	"hash"
	"sort"
	"strings"
)

// ******** Private variables ********

// hashTypeMap maps the hash type name to the hash function.
var hashTypeMap = make(map[string]func() hash.Hash)

// ******** Public functions ********

// NewHashFunctionOfType creates a hash function from the hash type name.
func NewHashFunctionOfType(hashTypeName string) hash.Hash {
	hashFunc, ok := hashTypeMap[strings.ToLower(strings.TrimSpace(hashTypeName))]
	if ok {
		return hashFunc()
	} else {
		return nil
	}
}

// KnownHashNames returns an array of valid known names.
func KnownHashNames() []string {
	result := make([]string, 0, len(hashTypeMap))
	for name, _ := range hashTypeMap {
		result = append(result, name)
	}

	sort.Strings(result)
	
	return result
}

// -------- Hash helper functions --------
// These helpers encapsulate the strange interface of the Blake2x functions
// to look like the interface of the other hash functions.

// NewBlake2b_256 creates a Blake2b-256 hash function.
func NewBlake2b_256() hash.Hash {
	hashFunc, _ := blake2b.New256(nil)
	return hashFunc
}

// NewBlake2b_384 creates a Blake2b-384 hash function.
func NewBlake2b_384() hash.Hash {
	hashFunc, _ := blake2b.New384(nil)
	return hashFunc
}

// NewBlake2b_512 creates a Blake2b-512 hash function.
func NewBlake2b_512() hash.Hash {
	hashFunc, _ := blake2b.New512(nil)
	return hashFunc
}

// NewBlake2s_128 creates a Blake2s-128 hash function.
func NewBlake2s_128() hash.Hash {
	hashFunc, _ := blake2s.New128(nil)
	return hashFunc
}

// NewBlake2s_256 creates a Blake2s-256 hash function.
func NewBlake2s_256() hash.Hash {
	hashFunc, _ := blake2s.New256(nil)
	return hashFunc
}

// ******** Private functions ********

// init is the package initialization function.
func init() {
	hashTypeMap[`md5`] = md5.New
	hashTypeMap[`sha1`] = sha1.New
	hashTypeMap[`sha2-224`] = sha256.New224
	hashTypeMap[`sha2-256`] = sha256.New
	hashTypeMap[`sha2-384`] = sha512.New384
	hashTypeMap[`sha2-512`] = sha512.New
	hashTypeMap[`sha2-512_224`] = sha512.New512_224
	hashTypeMap[`sha2-512_256`] = sha512.New512_256
	hashTypeMap[`sha3-256`] = sha3.New256
	hashTypeMap[`sha3-384`] = sha3.New384
	hashTypeMap[`sha3-512`] = sha3.New512
	hashTypeMap[`blake2b-256`] = NewBlake2b_256
	hashTypeMap[`blake2b-384`] = NewBlake2b_384
	hashTypeMap[`blake2b-512`] = NewBlake2b_512
	hashTypeMap[`blake2s-128`] = NewBlake2s_128
	hashTypeMap[`blake2s-256`] = NewBlake2s_256
}
