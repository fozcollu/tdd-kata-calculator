package calculator

type ICalculator interface {
	Add(numbers string) (int, error)
}

type Calculator struct {
}

func New() ICalculator {
	return Calculator{}
}

func (c Calculator) Add(text string) (int, error) {
	return 1, nil
}
