package math

import "testing"

func Test1(t *testing.T) {
	x := 1
	result := Compute(x)
	if result != 2 {
		t.Error(`Compute("1") = false"`)
	}
}

func Test2(t *testing.T) {
	x := 2
	result := Compute(x)
	if result != 4 {
		t.Error(`Compute("2") = false"`)
	}
}

func Test3(t *testing.T) {
	x := 3
	result := Compute(x)
	if result != 7 {
		t.Error(`Compute("3") = false"`)
	}
}

func Test4(t *testing.T) {
	x := 4
	result := Compute(x)
	if result != 28 {
		t.Error(`Compute("4") = false"`)
	}
}

func Test5(t *testing.T) {
	x := 5
	result := Compute(x)
	if result != 33 {
		t.Error(`Compute("5") = false"`)
	}
}

func Test6(t *testing.T) {
	x := 6
	result := Compute(x)
	if result != 198 {
		t.Error(`Compute("6") = false"`)
	}
}

func Test7(t *testing.T) {
	x := 7
	result := Compute(x)
	if result != 205 {
		t.Error(`Compute("7") = false"`)
	}
}
