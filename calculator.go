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
	sum := 0

	for _, numberText := range numberTexts {
		number, _ := strconv.Atoi(numberText)
		sum += number
	}

	return sum, nil
}
