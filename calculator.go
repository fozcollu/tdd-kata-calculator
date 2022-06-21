package calculator

import (
	"regexp"
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
	delimiter := ","
	hasDelimiter, _ := regexp.MatchString("^//(.)\n(.*)$", text)

	if hasDelimiter == true {
		endIndex := strings.Index(text, "\n")
		delimiter = text[2:endIndex]
	}

	text = strings.ReplaceAll(text, "\n", delimiter)
	numberTexts := strings.Split(text, delimiter)
	sum := 0

	for _, numberText := range numberTexts {
		number, _ := strconv.Atoi(numberText)
		sum += number
	}

	return sum, nil
}
