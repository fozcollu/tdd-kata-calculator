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

func TestAdd_HandleNewLinesBetweenNumbers(t *testing.T) {
	input := "1\n2,3"
	expected := 6

	stringCalculator := New()
	if actual, _ := stringCalculator.Add(input); expected != actual {
		t.Fatalf("Expected: %d, Actual: %d", expected, actual)
	}
}

func TestAdd_SupportDifferentDelimiters(t *testing.T) {
	testModels := []struct {
		input    string
		expected int
	}{
		{input: "//;\n1;2", expected: 3},
		{input: "//*\n1*2*8", expected: 11},
	}
	stringCalculator := New()

	for i, testModel := range testModels {
		if actual, _ := stringCalculator.Add(testModel.input); testModel.expected != actual {
			t.Fatalf("case %d - Expected: %d, Actual: %d", i+1, testModel.expected, actual)
		}
	}
}

func TestAdd_NegativeNumbersNotAllowed(t *testing.T) {
	testModels := []struct {
		input            string
		expectedErrorMsg string
	}{
		{input: "-1,2", expectedErrorMsg: "negatives not allowed / -1"},
		{input: "-1,-2,3,-8", expectedErrorMsg: "negatives not allowed / -1,-2,-8"},
	}
	stringCalculator := New()

	for i, testModel := range testModels {
		_, err := stringCalculator.Add(testModel.input)
		if err == nil {
			t.Fatalf("case :%d - Expected Error: %s, but actual no error", i+1, testModel.expectedErrorMsg)
		} else if testModel.expectedErrorMsg != err.Error() {
			t.Fatalf("case :%d - Expected Error: %s, Actual: %s", i+1, testModel.expectedErrorMsg, err.Error())
		}
	}
}

func TestAdd_IgnoreBiggerThan1000Numbers(t *testing.T) {
	stringCalculator := New()
	if actual, _ := stringCalculator.Add("2,1001"); actual != 2 {
		t.Fatalf(" Expected: %d, Actual: %d", 2, actual)
	}
}
