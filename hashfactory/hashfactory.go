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
// Version: 4.2.0
//
// Change history:
//    2024-12-20: V1.0.0: Created.
//    2024-12-21: V2.0.0: Make Blake2x creation functions private.
//    2025-01-28: V2.1.0: Remove "blake2s-128" as it needs a key.
//    2025-02-05: V3.0.0: New package name.
//    2025-02-26: V4.0.0: No longer return normalized hash type name.
//    2025-03-02: V4.1.0: Remove conversion no longer necessary.
//    2025-04-17: V4.2.0: Change names from "hash type" to "hash algorithm".
//

// Package hashfactory implements the hash factory functions.
package hashfactory

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
	"hash"
	"slices"
)

// ******** Private variables ********

// hashAlgorithmNameToFunction maps the hash algorithm name to the hash creation function.
var hashAlgorithmNameToFunction = make(map[string]func() hash.Hash)

// ******** Public functions ********

// New creates a hash function from the hash algorithm name.
func New(hashAlgorithm string) (hash.Hash, bool) {
	hashCreationFunction, ok := hashAlgorithmNameToFunction[hashAlgorithm]

	if ok {
		return hashCreationFunction(), ok
	} else {
		return nil, ok
	}
}

// KnownHashNames returns an array of valid known names.
func KnownHashNames() []string {
	result := make([]string, 0, len(hashAlgorithmNameToFunction))
	for name := range hashAlgorithmNameToFunction {
		result = append(result, name)
	}

	slices.Sort(result)

	return result
}

// ******** Private functions ********

// init is the package initialization function.
func init() {
	hashAlgorithmNameToFunction[`md5`] = md5.New
	hashAlgorithmNameToFunction[`sha1`] = sha1.New
	hashAlgorithmNameToFunction[`sha2-224`] = sha256.New224
	hashAlgorithmNameToFunction[`sha2-256`] = sha256.New
	hashAlgorithmNameToFunction[`sha2-384`] = sha512.New384
	hashAlgorithmNameToFunction[`sha2-512`] = sha512.New
	hashAlgorithmNameToFunction[`sha2-512_224`] = sha512.New512_224
	hashAlgorithmNameToFunction[`sha2-512_256`] = sha512.New512_256
	hashAlgorithmNameToFunction[`sha3-224`] = sha3.New224
	hashAlgorithmNameToFunction[`sha3-256`] = sha3.New256
	hashAlgorithmNameToFunction[`sha3-384`] = sha3.New384
	hashAlgorithmNameToFunction[`sha3-512`] = sha3.New512
	hashAlgorithmNameToFunction[`blake2b-256`] = newBlake2b_256
	hashAlgorithmNameToFunction[`blake2b-384`] = newBlake2b_384
	hashAlgorithmNameToFunction[`blake2b-512`] = newBlake2b_512
	hashAlgorithmNameToFunction[`blake2s-256`] = newBlake2s_256
}

// -------- Hash helper functions --------
// These helpers encapsulate the strange interface of the Blake2x functions
// to look like the interface of the other hash functions.
// The only error that can be returned is a "key too long" error, which can not
// happen here, as there is no key.
// A proper interface would have had two different "New" functions. One for the hash
// and one for the MAC.

// newBlake2b_256 creates a Blake2b-256 hash function.
func newBlake2b_256() hash.Hash {
	hashFunc, _ := blake2b.New256(nil)
	return hashFunc
}

// newBlake2b_384 creates a Blake2b-384 hash function.
func newBlake2b_384() hash.Hash {
	hashFunc, _ := blake2b.New384(nil)
	return hashFunc
}

// newBlake2b_512 creates a Blake2b-512 hash function.
func newBlake2b_512() hash.Hash {
	hashFunc, _ := blake2b.New512(nil)
	return hashFunc
}

// newBlake2s_256 creates a Blake2s-256 hash function.
func newBlake2s_256() hash.Hash {
	hashFunc, _ := blake2s.New256(nil)
	return hashFunc
}
