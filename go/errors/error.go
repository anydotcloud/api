// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package errors

import "errors"

var (
	// NotFound indicates the resource could not be found
	NotFound = errors.New("resource not found")

	// Duplicate indicates there was an attempt to create a resource with an
	// identifier or unique field that already exists
	Duplicate = errors.New("resource already exists")

	// InvalidInput indicates there was a failure to validate some input data
	InvalidInput = errors.New("invalid input")
)
