package main

import (
    "testing"
)

// TestConvert tests the Convert function.
func TestConvert(t *testing.T) {
    // test cases...
}func TestConvert(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "Divisible by 3", input: 3, expected: "Fizz"},
		{name: "Divisible by 5", input: 5, expected: "Buzz"},
		{name: "Divisible by 3 and 5", input: 15, expected: "FizzBuzz"},
		{name: "Not divisible by 3 or 5", input: 7, expected: ""},
		{name: "Zero input", input: 0, expected: "Error: input must be greater than 0"},
		{name: "Negative input", input: -5, expected: "Error: input must be greater than 0"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Convert(tc.input)
			if err != nil && result != tc.expected {
				t.Errorf("Convert(%d) returned error %v; expected %s", tc.input, err, tc.expected)
			} else if err == nil && result != tc.expected {
				t.Errorf("Convert(%d) = %s; expected %s", tc.input, result, tc.expected)
			}
		})
	}
}