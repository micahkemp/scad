package internal

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"a", false},
		{"0a", false},
		{"a-a", false},
		{"_a", false},
		{"a_", false},
		{"a a", false},
		{"a0a", true},
		{"aa", true},
		{"a_a", true},
		{"a0", true},
	}

	for _, test := range tests {
		if got := valid(test.input); got != test.want {
			t.Errorf("Valid(%q) = %v", test.input, got)
		}
	}
}

func TestFirstNonEmptyName(t *testing.T) {
	tests := []struct {
		choices  []string
		wantName string
		wantOk   bool
	}{
		{[]string{"", "a"}, "a", false},
		{[]string{"", "a", "aa", "bb", "cc"}, "a", false},
		{[]string{"", "aa", "bb", "cc"}, "aa", true},
	}

	for _, test := range tests {
		if got, ok := FirstNonEmptyName(test.choices...); got != test.wantName || ok != test.wantOk {
			t.Errorf("FirstNonEmptyName(%q) = (%q, %v), want (%q, %v)",
				test.choices,
				got,
				ok,
				test.wantName,
				test.wantOk)
		}
	}
}

func TestFilenameFilePath(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		wantFilename string
		wantFilePath string
	}{
		{"aa", ".", "aa.scad", "aa.scad"},
		{"bb", "relative_output_path", "bb.scad", "relative_output_path/bb.scad"},
		{"cc", "/absolute_output_path", "cc.scad", "/absolute_output_path/cc.scad"},
	}

	for _, test := range tests {
		n, _ := NewName(test.name)
		gotFilename := n.Filename()
		gotFilePath := n.FilePath(test.path)

		if gotFilename != test.wantFilename {
			t.Errorf("NewName(%q).Filename() = %q, want %q",
				test.name,
				gotFilename,
				test.wantFilename)
		}

		if gotFilePath != test.wantFilePath {
			t.Errorf("NewName(%q).FilePath(%q) = %q, want %q",
				test.name,
				test.path,
				gotFilePath,
				test.wantFilePath)
		}
	}
}
