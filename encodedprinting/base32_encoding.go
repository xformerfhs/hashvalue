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
// Version: 1.0.0
//
// Change history:
//    2025-03-02: V1.0.0: Created.
//

package encodedprinting

import (
	"encoding/base32"
	"os"
)

// Base32Encoder is used to encode bytes in base32 encoding.
type Base32Encoder struct {
	encoder *base32.Encoding
}

// NewBase32Encoder creates a new base32 encoder.
func NewBase32Encoder() *Base32Encoder {
	return &Base32Encoder{base32.StdEncoding.WithPadding(base32.NoPadding)}
}

// PrintEncoded prints bytes slices in base32 encoding.
func (e *Base32Encoder) PrintEncoded(value []byte) {
	writeStringln(os.Stdout, e.encoder.EncodeToString(value))
}
