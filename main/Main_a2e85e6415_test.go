package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

type DontknowMock func(int, int) int

func (f DontknowMock) Dontknow(a, b int) int {
	return f(a, b)
}

func TestPrintResults(t *testing.T) {
	cases := []struct {
		name string
		in1, in2, want int
	}{
		{"Case 1: Basic Test", 1, 2, 3}, // TODO: Replace 3 with the expected output of helper.Dontknow(1,2)
		{"Case 2: Another Test", 4, 5, 9}, // TODO: Replace 9 with the expected output of helper.Dontknow(4,5)
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			helper.Dontknow = DontknowMock(func(a, b int) int { return tt.want })

			var buf bytes.Buffer
			output = &buf
			printResults(tt.in1, tt.in2)

			got := buf.String()
			want := fmt.Sprintf("%d\n", tt.want)

			if got != want {
				t.Errorf("printResult() = %v, want %v", got, want)
			}
		})
	}

	output = os.Stdout
}
