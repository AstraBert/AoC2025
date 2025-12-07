package main

import "testing"

func TestCountSplits(t *testing.T) {
	testCase := []struct {
		file     string
		expected int
	}{
		{file: "test.txt", expected: 21},
	}
	for _, tc := range testCase {
		result, err := CountSplits(tc.file)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}
