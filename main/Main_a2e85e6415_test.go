package main

import (
	"fmt"
	"testing"

	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Valid input
	a := helper.Dontknow(1, 2)
	expected := 3
	if a != expected {
		t.Errorf("Expected %d, got %d", expected, a)
	} else {
		t.Logf("Success: Dontknow(%d, %d) = %d", 1, 2, a)
	}

	// Test case 2: Invalid input
	a = helper.Dontknow(-1, 2)
	expected = -1
	if a != expected {
		t.Errorf("Expected %d, got %d", expected, a)
	} else {
		t.Logf("Success: Dontknow(%d, %d) = %d", -1, 2, a)
	}
}
