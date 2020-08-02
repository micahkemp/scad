package name

import "testing"

func TestFilename(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"aa", "aa.scad"},
		{"bb", "bb.scad"},
	}

	for _, test := range tests {
		n, _ := NewName(test.input)
		got := n.Filename()

		if got != test.want {
			t.Errorf("NewName(%q).Filename() = %q, want %q",
				test.input,
				got,
				test.want)
		}
	}
}