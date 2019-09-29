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

import (
	"sync"

	"golang.org/x/net/context"
)

type (
	// Geography ...
	Geography interface {
		// GetLocationDistance ...
		GetLocationDistance(ctx context.Context, lon, lat float64) (float64, error)
	}

	geography struct {
		// repository GeographyRepository
	}
)

var (
	geographyInstance Geography
	geographyOnce     sync.Once
)

func (f *engineFactory) NewGeography() Geography {
	geographyOnce.Do(func() {
		geographyInstance = &geography{
			//repository: f.NewGeographyRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return geographyInstance
}
