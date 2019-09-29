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
	"context"
	"fmt"

	geo "github.com/kellydunn/golang-geo"
)

func (s *geography) GetLocationDistance(ctx context.Context, lon, lat float64) (float64, error) {
	// Make a few points
	p := geo.NewPoint(42.25, 120.2)
	p2 := geo.NewPoint(30.25, 112.2)

	// find the great circle distance between them
	dist := p.GreatCircleDistance(p2)
	fmt.Printf("great circle distance: %f\n", dist)

	return dist, nil
}
