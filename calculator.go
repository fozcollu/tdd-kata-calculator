package calculator

import (
	"errors"
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
	delimiter, extractedText := c.extractDelimiterAndNumbers(text)

	extractedText = strings.ReplaceAll(extractedText, "\n", delimiter)
	numbers := strings.Split(extractedText, delimiter)
	sum := 0
	var negatives []string
	for _, numberText := range numbers {
		number, _ := strconv.Atoi(numberText)

		if number < 0 {
			negatives = append(negatives, numberText)
		}
		sum += number
	}

	if len(negatives) > 0 {
		return -1, errors.New("negatives not allowed / " + strings.Join(negatives, ","))
	}

	return sum, nil
}

func (c *Calculator) extractDelimiterAndNumbers(text string) (string, string) {
	//default delimiter is ','
	delimiter := ","
	delimiterPattern := "^//(.)\n(.*)$"
	hasDifferentDelimiter, _ := regexp.MatchString(delimiterPattern, text)
	numbers := text
	if hasDifferentDelimiter == true {
		endIndex := strings.Index(text, "\n")
		delimiter = (text)[2:endIndex]
		numbers = text[endIndex:]

	}
	return delimiter, numbers
}
