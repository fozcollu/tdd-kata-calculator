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
	numberTexts := strings.Split(extractedText, delimiter)

	numbers := c.convertNumberTextsToInteger(numberTexts)
	numbers = c.ignoreBiggerThan1000(numbers)

	//negatives not allowed
	negatives := c.findNegatives(numbers)
	if len(negatives) > 0 {
		return -1, errors.New("negatives not allowed / " + strings.Join(negatives, ","))
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum, nil
}

func (c *Calculator) convertNumberTextsToInteger(numberTexts []string) []int {
	var numbers []int
	for _, numberText := range numberTexts {
		number, _ := strconv.Atoi(numberText)
		numbers = append(numbers, number)
	}
	return numbers
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

func (c *Calculator) findNegatives(numbers []int) []string {
	var negatives []string
	for _, number := range numbers {
		if number < 0 {
			negatives = append(negatives, strconv.Itoa(number))
		}
	}
	return negatives
}

func (c *Calculator) ignoreBiggerThan1000(numbers []int) []int {
	var validNumbers []int

	for _, number := range numbers {
		if number < 1000 {
			validNumbers = append(validNumbers, number)
		}
	}

	return validNumbers

}
