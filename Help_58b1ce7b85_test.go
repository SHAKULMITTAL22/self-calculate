package main

import (
	"testing"
)

// Testing help function
func TestHelp(t *testing.T) {
	a := 5
	b := 10
	expected := 15
	actual := sum(a, b)
	if actual != expected {
		t.Errorf("Test failed for %v + %v, expected: %v, got: %v", a, b, expected, actual)
	}

	a = -5
	b = 10
	expected = 5
	actual = sum(a, b)
	if actual != expected {
		t.Errorf("Test failed for %v + %v, expected: %v, got: %v", a, b, expected, actual)
	}

	b = -10
	expected = -15
	actual = sum(a, b)
	if actual != expected {
		t.Errorf("Test failed for %v + %v, expected: %v, got: %v", a, b, expected, actual)
	} else {
		t.Logf("Test success for %v + %v, result: %v", a, b, actual)
	}
}

func sum(a int, b int) int {
	return a + b
}
