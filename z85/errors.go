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
// Version: 2.0.0
//
// Change history:
//    2025-02-03: V1.0.0: Created.
//    2025-02-13: V2.0.0: Change type of invalid byte error, correct type of length error.
//

package z85

import (
	"errors"
	"fmt"
)

// ******** Private constants ********

// invalidLengthMessage means that the input has a length that is not valid for the operation.
const invalidLengthMessage = `input length is not a multiple of %d`

// ******** Public types and functions ********

// ErrTooLong means that the input is too long for the operation.
var ErrTooLong = errors.New(`input is too long`)

// ErrInvalidLength means that the input has a length that is not valid for the operation.
type ErrInvalidLength byte

// Error returns the error message for an invalid length error.
func (e ErrInvalidLength) Error() string {
	return fmt.Sprintf(invalidLengthMessage, e)
}
