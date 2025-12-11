package main

import "testing"

func TestFindAllWaysOut(t *testing.T) {
	testCase := []struct {
		file     string
		expected int
	}{
		{file: "test.txt", expected: 5},
	}
	for _, tc := range testCase {
		result, err := FindAllWaysOut(tc.file)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}

func TestFindAllWaysOutComplex(t *testing.T) {
	testCase := []struct {
		file     string
		expected int
	}{
		{file: "test_2.txt", expected: 2},
	}
	for _, tc := range testCase {
		result, err := FindAllWaysOutComplex(tc.file)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}
