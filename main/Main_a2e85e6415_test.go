package main

import (
	"testing"
	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

func TestMain_a2e85e6415(t *testing.T) {
    {
        a := helper.Help(1, 2)
        expected := 3
        if a != expected {
            t.Errorf("Test Case 1 Failed: {1, 2} was expected to return %v but got %v", expected, a)
        } else {
            t.Logf("Test Case 1 Successful: {1, 2} returned expected result %v", a)
        }
    }
    {
        a := helper.Help(-1, -2)
        expected := -3
        if a != expected {
            t.Errorf("Test Case 2 Failed: {-1, -2} was expected to return %v but got %v", expected, a)
        } else {
            t.Logf("Test Case 2 Successful: {-1, -2} returned expected result %v", a)
        }
    }
}
