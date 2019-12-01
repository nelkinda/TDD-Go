package main

import "testing"

func Test_main(t *testing.T) {
	if expected, actual := "Hello, world!", getMessage(); expected != actual {
		t.Errorf("Expected and actual differ:\n<%s>\n<%s>\n", expected, actual)
	}
}