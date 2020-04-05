package validation

import (
	"errors"
	"strconv"
	"strings"
)

func ValidatePrice(price string) (float64, error) {
	priceSpaceRemoval := strings.TrimSpace(price)

	if priceSpaceRemoval == "" {
		return 0, errors.New("Введите цену")
	}
	praceInt, err := strconv.ParseFloat(priceSpaceRemoval, 64)
	if err != nil {
		return 0, errors.New("Задано некорректное значение")
	}
	if praceInt <= 0 {
		return 0, errors.New("Цена не может быть отрицательной или равна нулю")
	}
	return praceInt, nil
}
