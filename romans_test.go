package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	expectation := "III"
	result := RomanToInt("III")
	if result != expectation {
		t.Fatalf("No good, because result was %v isntead of %v", result, expectation)
	}
}
