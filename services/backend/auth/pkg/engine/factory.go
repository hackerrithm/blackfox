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

package engine

type (
	// Factory interface allows us to provide
	// other parts of the system with a way to make
	// instances of our use-case / interactors when
	// they need to
	Factory interface {
		// NewUser creates a new User interactor
		NewUser() Auth
	}

	// engine factory stores the state of our engine
	// which only involves a storage factory instance
	engineFactory struct {
		AuthenticationFactory
	}
)

// NewEngine creates a new engine factory that will
// make use of the passed in StorageFactory for any
// data persistence needs.
func NewEngine(jwt AuthenticationFactory) Factory {
	return &engineFactory{
		jwt,
	}
}
