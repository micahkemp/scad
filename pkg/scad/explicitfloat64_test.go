package scad

import "testing"

func TestExplicitFloat64_scadString(t *testing.T) {
	tests := fieldValueFormatterTests{
		{
			ExplicitFloat64(1),
			"1",
		},
		{
			ExplicitFloat64(1.001),
			"1.001",
		},
	}

	tests.run(t)
}
