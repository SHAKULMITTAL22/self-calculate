package main

import (
	"fmt"

	"github.com/SHAKULMITTAL22/self-calculate/helper"
	"testing"
)

func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Positive integers
	a := helper.Dontknow(1, 2)
	expected := 2
	if a != expected {
		t.Errorf("Dontknow(1, 2) = %d, expected %d", a, expected)
	} else {
		t.Log("Test case 1 passed")
	}

	// Test case 2: Negative integers
	a = helper.Dontknow(-1, -2)
	expected = 2
	if a != expected {
		t.Errorf("Dontknow(-1, -2) = %d, expected %d", a, expected)
	} else {
		t.Log("Test case 2 passed")
	}

	// Test case 3: Zero values
	a = helper.Dontknow(0, 0)
	expected = 0
	if a != expected {
		t.Errorf("Dontknow(0, 0) = %d, expected %d", a, expected)
	} else {
		t.Log("Test case 3 passed")
	}

	// Test case 4: Large integers
	a = helper.Dontknow(1000000000, 1000000000)
	expected = 1000000000000000000
	if a != expected {
		t.Errorf("Dontknow(1000000000, 1000000000) = %d, expected %d", a, expected)
	} else {
		t.Log("Test case 4 passed")
	}

	// Test case 5: Floating point numbers
	a = helper.Dontknow(1.5, 2.5)
	expected = 3
	if a != expected {
		t.Errorf("Dontknow(1.5, 2.5) = %d, expected %d", a, expected)
	} else {
		t.Log("Test case 5 passed")
	}

	// Test case 6: Invalid input
	a = helper.Dontknow("a", "b")
	expected = 0
	if a != expected {
		t.Errorf("Dontknow(\"a\", \"b\") = %d, expected %d", a, expected)
	} else {
		t.Log("Test case 6 passed")
	}
}
