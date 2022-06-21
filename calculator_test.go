package calculator

import "testing"

func TestCalculator_WithUpTo2Numbers(t *testing.T) {
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
