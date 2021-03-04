package math

import "testing"

func Test1(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{1, 2},
		{2, 4},
		{3, 7},
		{4, 28},
		{5, 33},
		{6, 198},
		{7, 205},
	}
	for _, test := range tests {
		if got := Compute(test.input); got != test.want {
			t.Errorf("Compute(%q) = %v", test.input, got)
		}
	}
}
