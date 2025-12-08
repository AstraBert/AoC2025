package main

import "testing"

func TestGetCircuits(t *testing.T) {
	testCase := []struct {
		file     string
		numPairs int
		expected int
	}{
		{file: "test.txt", numPairs: 10, expected: 40},
	}
	for _, tc := range testCase {
		result, err := GetCircuits(tc.file, tc.numPairs)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}

func TestGetCircuitsComplex(t *testing.T) {
	testCase := []struct {
		file     string
		expected int
	}{
		{file: "test.txt", expected: 25272},
	}
	for _, tc := range testCase {
		result, err := GetCircuitsComplex(tc.file)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}
