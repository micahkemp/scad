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

package primitive2d

import (
	"github.com/micahkemp/scad/pkg/scad/value"
)

// Square is a square.
type Square struct {
	// Only one of Size or SizeXY should be set.
	Size   value.Float `scad:"size"`
	SizeXY value.Float `scad:"size"`

	Center value.Bool
}
