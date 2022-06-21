package calculator

import (
	"strconv"
	"strings"
)

type ICalculator interface {
	Add(numbers string) (int, error)
}

type Calculator struct {
}

func New() ICalculator {
	return Calculator{}
}

func (c Calculator) Add(text string) (int, error) {
	numberTexts := strings.Split(text, ",")
	firstNumber := 0
	secondNumber := 0

	if len(numberTexts) > 0 {
		firstNumber, _ = strconv.Atoi(numberTexts[0])
	}
	if len(numberTexts) > 1 {
		secondNumber, _ = strconv.Atoi(numberTexts[1])
	}

	return firstNumber + secondNumber, nil
}
