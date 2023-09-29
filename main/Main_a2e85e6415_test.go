package main

import (
	"testing"
)

func Dontknow(a int, b int) int {
	return a + b
}

func TestDontknow(t *testing.T) {
	total := Dontknow(1, 2)
	if total != 3 {
		t.Errorf("Dontknow was incorrect, got: %d, want: %d.", total, 3)
	}

	total = Dontknow(1, -1)
	if total != 0 {
		t.Errorf("Dontknow was incorrect, got: %d, want: %d.", total, 0)
	}
}
