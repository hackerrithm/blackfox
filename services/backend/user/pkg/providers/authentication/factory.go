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

package authentication

import (
	"sync"

	"github.com/hackerrithm/blackfox/services/backend/user/pkg/engine"
)

type (
	authenticationFactory struct {
		// token jwtgo.Token
	}
)

var (
	jwtInstance engine.JWTSignParser
	jwtOnce     sync.Once
)

// NewAuthentication ...
func NewAuthentication(token string) engine.AuthenticationFactory {
	return &authenticationFactory{}
}

// NewUserRepository creates a new datastore User repository
func (f *authenticationFactory) NewAuthentication() engine.JWTSignParser {
	jwtOnce.Do(func() {
		jwtInstance = newJWT()
	})
	return jwtInstance
}
