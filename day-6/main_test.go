package main

import "testing"

func TestSolveAllProblems(t *testing.T) {
	testCase := []struct {
		file     string
		expected uint64
		complex  bool
	}{
		{file: "test.txt", expected: 4277556, complex: false},
		{file: "test.txt", expected: 3263827, complex: true},
	}
	for _, tc := range testCase {
		result, err := SolveAllProblems(tc.file, tc.complex)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}
