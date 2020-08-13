package scad_test

import (
	"github.com/micahkemp/scad/scad"
	"testing"
)

func TestTranslate(t *testing.T) {
	child1 := scad.NewCube(1, 2, 3, "")

	test := struct {
		x, y, z float64
		n       string
		want    string
	}{
		1, 2, 3, "", `use <cube_component/cube_component.scad>

module translate_component() {
	translate(v=[1, 2, 3]) {
		cube_component();
	}
}

// call module when loaded directly
translate_component();
`,
	}

	newTranslate := scad.NewTranslate(test.x, test.y, test.z, test.n, child1)

	if newTranslate.String() != test.want {
		t.Errorf("NewTranslate(%f, %f, %f, %s, %v).String() = %s, want %s",
			test.x,
			test.y,
			test.z,
			test.n,
			child1,
			newTranslate.String(),
			test.want,
		)
	}

	// this should give us the same results as creating the Translate object directly
	// but we want to give a good failure message, so it's broken out separately
	childTranslate := child1.Translate(test.x, test.y, test.z, test.n)
	if childTranslate.String() != test.want {
		t.Errorf("child1.Translate(%f, %f, %f, %s).String() = %s, want %s",
			test.x,
			test.y,
			test.z,
			test.n,
			childTranslate.String(),
			test.want,
		)
	}
}
