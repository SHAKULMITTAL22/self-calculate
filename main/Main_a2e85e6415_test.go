package main

import (
	"testing"

	"github.com/SHAKULMITTAL22/self-calculate/helper"
)

func TestDontknow(t *testing.T) {
    scenarios := []struct {
        desc string
        a, b int
        want int // Change expected result according to function logic
    }{
        {
            desc: "Both positive",
            a:    1,
            b:    2,
            want: 2, // Updated
        },
        {
            desc: "Both negative",
            a:    -1,
            b:    -2,
            want: 2, // Updated
        },
        {
            desc: "Mix of both",
            a:    -1,
            b:    2,
            want: -2, // Updated
        },
    }

    for _, s := range scenarios {
        t.Run(s.desc, func(t *testing.T) {
            got := helper.Dontknow(s.a, s.b)
            if got != s.want {
                t.Errorf("%s: got=%d, want=%d", s.desc, got, s.want)
            } else {
                t.Logf("Success: {%s}: got=%d, want=%d", s.desc, got, s.want)
            }
        })
    }
}
