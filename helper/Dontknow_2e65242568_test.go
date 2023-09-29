// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model codechat-bison-32k

package helper

import "testing"

func TestDontknow_2e65242568(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"Test 1", 1, 2, 2},
		{"Test 2", 3, 4, 12},
		{"Test 3", 5, 6, 30},
		{"Test 4", 7, 8, 56},
		{"Test 5", 9, 10, 90},
	}
	for _, tt := range tests {
		t.Logf("Testing with a=%d and b=%d", tt.a, tt.b)
		got := Dontknow(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Dontknow(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		} else {
			t.Logf("Test passed with a=%d and b=%d", tt.a, tt.b)
		}
	}
}
