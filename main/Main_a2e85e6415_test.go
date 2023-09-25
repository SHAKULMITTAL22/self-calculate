package main

import (
    "fmt"
    "testing"
    "github.com/SHAKULMITTAL22/self-calculate/helper"
)

func TestMain_a2e85e6415(t *testing.T) {
    type args struct {
        a int
        b int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "Test with positive numbers",
            args: args{a: 1, b: 2},
            want: 3,
        },
        {
            name: "Test with zero values",
            args: args{a: 0, b: 0},
            want: 0,
        },
        {
            name: "Test with negative numbers",
            args: args{a: -1, b: -1},
            want: -2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := helper.Dontknow(tt.args.a, tt.args.b); got != tt.want {
                t.Errorf("Dontknow() = %v, want %v", got, tt.want)
            } else {
                t.Logf("Test %v passed", tt.name)
            }
        })
    }
}

func ExampleMain() {
    a := helper.Dontknow(1, 2)
    fmt.Println(a)
}
