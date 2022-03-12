// Copyright 2022 Micah Kemp
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extrude

import (
	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/values"
)

// LinearExtrude is a linear extrude.
type LinearExtrude struct {
	linearExtrude scad.AutoFunctionName `scad:"linear_extrude"` //nolint:golint,structcheck,unused

	Height values.Float `scad:"height"`
	Twist  values.Float `scad:"twist"`
	Center values.Bool  `scad:"center"`
	Slices values.Int   `scad:"slices"`

	// Only one of Scale or ScaleXY should be set.
	Scale   values.Float   `scad:"scale"`
	ScaleXY values.FloatXY `scad:"scale"`

	Children []interface{}
}
