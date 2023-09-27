package main

import (
	"fmt"
	"testing"

	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Positive integers
	a := helper.Dontknow(1, 2)
	expected := 2
	if a != expected {
		t.Errorf("Dontknow(1, 2) = %d, expected %d", a, expected)
	} else {
		fmt.Println("Test case 1 passed")
	}

	// Test case 2: Negative integers
	a = helper.Dontknow(-1, -2)
	expected = 2
	if a != expected {
		t.Errorf("Dontknow(-1, -2) = %d, expected %d", a, expected)
	} else {
		fmt.Println("Test case 2 passed")
	}

	// Test case 3: Zero
	a = helper.Dontknow(0, 2)
	expected = 0
	if a != expected {
		t.Errorf("Dontknow(0, 2) = %d, expected %d", a, expected)
	} else {
		fmt.Println("Test case 3 passed")
	}

	// Test case 4: Large numbers
	a = helper.Dontknow(1000000000, 2000000000)
	expected = 2000000000000000000
	if a != expected {
		t.Errorf("Dontknow(1000000000, 2000000000) = %d, expected %d", a, expected)
	} else {
		fmt.Println("Test case 4 passed")
	}

	// Test case 5: Error handling
	_, err := helper.Dontknow(1, 0)
	if err == nil {
		t.Errorf("Dontknow(1, 0) should return an error")
	} else {
		fmt.Println("Test case 5 passed")
	}
}
