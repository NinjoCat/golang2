package botlib

import (
	"errors"
	"strconv"
)

const (
	minBottlesNumber = 0
	maxBottlesNumber = 200
)

func PrintBottle(bottlesNumber int) (string, error) {
	if bottlesNumber < minBottlesNumber || bottlesNumber > maxBottlesNumber {
		return "", errors.New("Неправильное количество бутылок")
	}

	var bottlesByStr string

	bottlesNumberOst := bottlesNumber % 100
	if bottlesNumberOst >= 11 && bottlesNumberOst <= 19 {
		bottlesByStr = strconv.Itoa(bottlesNumber) + " бутылок"
	} else if bottlesNumber%10 == 1 {
		bottlesByStr = strconv.Itoa(bottlesNumber) + " бутылка"
	} else if bottlesNumber%10 == 2 || bottlesNumber%10 == 3 || bottlesNumber%10 == 4 {
		bottlesByStr = strconv.Itoa(bottlesNumber) + " бутылки"
	} else {
		bottlesByStr = strconv.Itoa(bottlesNumber) + " бутылок"
	}

	return bottlesByStr, nil
}
