//
// SPDX-FileCopyrightText: Copyright 2025 Frank Schwab
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
// Version: 3.2.0
//
// Change history:
//    2025-01-31: V1.0.0: Created.
//    2025-02-01: V2.0.0: Return an error on invalid padding for DecodeAndUnpad.
//    2025-02-02: V3.0.0: Structured errors.
//    2025-02-05: V3.1.0: Streamlined result size calculation.
//    2025-02-05: V3.2.0: Use unsigned types where possible.
//

// Package z85 implements Z85 encoding as specified in https://rfc.zeromq.org/spec/32.
// There are two additional functions that implement a padded variant for sources
// of lengths that are not a multiple of 4. Theses are _not_ standardized.
package z85

import (
	"encoding/binary"
	"math"
)

// ******** Private constants ********

// codeSize is the size of the encoding (i.e. the number of encoding characters).
const codeSize = 85

// byteChunkSize is the size of a byte chunk to be processed.
const byteChunkSize = 4

// byteChunkMask is the mask to check a byte chunk index.
const byteChunkMask = byteChunkSize - 1

// byteChunkShift is the shift value used for division by shifting.
const byteChunkShift = 2

// encodedChunkSize is the size of an encoded chunk.
const encodedChunkSize = 5

// encodeTable is the table used for encoding.
var encodeTable = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-:+=^!/*?&<>()[]{}@%$#`

// ******** Public functions ********

// Encode encodes a byte slice into a Z85 encoded string.
// The length of the slice must be a multiple of 4.
func Encode(source []byte) (string, error) {
	sourceLen := uint(len(source))

	if sourceLen > math.MaxInt/encodedChunkSize {
		return ``, ErrTooLong
	}

	if (sourceLen & byteChunkMask) != 0 {
		return ``, ErrInvalidLength(byteChunkSize)
	}

	chunkCount := sourceLen >> byteChunkShift
	result := make([]byte, sourceLen+chunkCount)
	destination := result
	for chunkIndex := uint(0); chunkIndex < chunkCount; chunkIndex++ {
		value := binary.BigEndian.Uint32(source[:byteChunkSize])

		// Generate 5 characters
		for i := byteChunkSize; i >= 0; i-- {
			valueDiv := value / codeSize
			destination[i] = encodeTable[value-(valueDiv*codeSize)]
			value = valueDiv
		}

		destination = destination[encodedChunkSize:]
		source = source[byteChunkSize:]
	}

	return string(result), nil
}
