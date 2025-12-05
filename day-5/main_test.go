package main

import "testing"

func TestCountFreshFoods(t *testing.T) {
	testCase := []struct {
		file     string
		expected int
	}{
		{file: "test.txt", expected: 3},
	}
	for _, tc := range testCase {
		result, err := CountFreshFoods(tc.file)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}

func TestCountFreshFoodsComplex(t *testing.T) {
	testCase := []struct {
		file     string
		expected int
	}{
		{file: "test.txt", expected: 14},
	}
	for _, tc := range testCase {
		result, err := CountFreshFoodsComplex(tc.file)
		if err != nil {
			t.Errorf("Not expecting an error, got %s", err.Error())
		}
		if result != tc.expected {
			t.Errorf("Expecting %d, got %d", tc.expected, result)
		}
	}
}
