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
	"github.com/xformerfhs/z85"
	"os"
)

// Z85Encoder is used to encode bytes in Z85 encoding.
type Z85Encoder struct {
	// There are no fields in this structure.
}

// NewZ85Encoder creates a new Z85 encoder.
func NewZ85Encoder() *Z85Encoder {
	return &Z85Encoder{}
}

// PrintEncoded prints bytes slices in Z85 encoding.
func (e *Z85Encoder) PrintEncoded(value []byte) {
	encoded, _ := z85.Encode(value)
	writeStringln(os.Stdout, encoded)
}
