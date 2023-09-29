package main

import (
	"testing"
	"os"
	"io/ioutil"
)

// TestMain_a2e85e6415 testing main function
func TestMain_a2e85e6415(t *testing.T) {

	// Create a pipe to capture stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	main()

	writer.Close()
	out, _ := ioutil.ReadAll(reader)

	// Test case 1 : Checking correct output
	expectedOutput1 := "2\n" // expected output is 2 as helper.Dontknow(1, 2) is called in main function
	if string(out) != expectedOutput1 {
		t.Errorf("Test Failed. Expected %v, got %v", expectedOutput1, string(out))
	} else {
		t.Logf("Test Success. Expected %v, got %v", expectedOutput1, string(out))
	}
}
