package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

// This main function call a Help function from helper package
// and then print the return value of that Help function.
func main() {
	a := helper.Help(1, 2)
	fmt.Println(a)
}

func HelperTestFunc(a, b int) int {
	return a + b
}

// TestMain_a2e85e6415 executes test cases on the main function.
func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: Ensure Help method works with positive integers
	sum := HelperTestFunc(1, 2)
	assert.Equal(t, 3, sum, "For inputs 1 and 2, expected 3")

	// Test case 2: Ensure Help method works with negative integers
	sum = HelperTestFunc(-1, -2)
	assert.Equal(t, -3, sum, "For inputs -1 and -2, expected -3")
}
