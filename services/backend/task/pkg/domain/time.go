// Copyright 2019 kemar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domain

import "time"

var (
	// now returns the current UTC time
	// It is a replaceable function to allow for easy unit testing
	now = defaultNow
)

// set it back to this function to restore normal functionality
func defaultNow() time.Time {
	return time.Now().UTC()
}
