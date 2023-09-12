package main

import "testing"

func TestFactorial(t *testing.T) {
	result := factorial(5)
	if result != 120 {
		t.Error(("Response from factorial is inexpected value"))
	}
}
