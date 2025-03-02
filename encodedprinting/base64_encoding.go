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
	"encoding/base64"
	"os"
)

// Base64Encoder is used to encode bytes in base64 encoding.
type Base64Encoder struct {
	encoder *base64.Encoding
}

// NewBase64Encoder creates a new base64 encoder.
func NewBase64Encoder() *Base64Encoder {
	return &Base64Encoder{base64.StdEncoding.WithPadding(base64.NoPadding)}
}

// PrintEncoded prints bytes slices in base64 encoding.
func (e *Base64Encoder) PrintEncoded(value []byte) {
	writeStringln(os.Stdout, e.encoder.EncodeToString(value))
}
