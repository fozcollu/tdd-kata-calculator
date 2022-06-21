package calculator

import "testing"

func TestAdd_WithUpTo2Numbers(t *testing.T) {
	testModels := []struct {
		input    string
		expected int
	}{
		{input: "", expected: 0},
		{input: "1", expected: 1},
		{input: "1,2", expected: 3},
	}

	stringCalculator := New()

	for i, testModel := range testModels {
		if actual, _ := stringCalculator.Add(testModel.input); testModel.expected != actual {
			t.Fatalf("case %d - Expected: %d, Actual: %d", i+1, testModel.expected, actual)
		}
	}
}

func TestAdd_WithUnknownAmountNumbers(t *testing.T) {
	testModels := []struct {
		input    string
		expected int
	}{
		{input: "1,2,3", expected: 6},
		{input: "3,4,6,8,9", expected: 30},
		{input: "2", expected: 2},
	}

	stringCalculator := New()

	for i, testModel := range testModels {
		if actual, _ := stringCalculator.Add(testModel.input); testModel.expected != actual {
			t.Fatalf("case %d - Expected: %d, Actual: %d", i+1, testModel.expected, actual)
		}
	}
}
