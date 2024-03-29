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

package primitive2d_test

import (
	"fmt"

	"github.com/micahkemp/scad/pkg/scad"
	"github.com/micahkemp/scad/pkg/scad/primitive2d"
	"github.com/micahkemp/scad/pkg/scad/value"
)

func ExampleText() {
	text := primitive2d.Text{
		Text: value.NewString("Hello World!"),
		Size: value.NewFloat(10),
	}

	content, _ := scad.FunctionContent(text)
	fmt.Println(content)
	// Output: text(size=10, text="Hello World!");
}
