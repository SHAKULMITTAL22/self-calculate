package main

import (
	"testing"
	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

func TestCanAddTwoNumbers(t *testing.T) {
	got := helper.Dontknow(1, 2)
	want := 3

	if got != want {
		t.Errorf("Dontknow(1, 2) = %v; want %v", got, want)
	} else {
        t.Logf("Dontknow(1, 2) = %v; Success", got)
    }
}

func TestCanHandleNegativeNumbers(t *testing.T) {
	got := helper.Dontknow(-1, -2)
	want := -3

	if got != want {
		t.Errorf("Dontknow(-1, -2) = %v; want %v", got, want)
	} else {
        t.Logf("Dontknow(-1, -2) = %v; Success", got)
    }
}
