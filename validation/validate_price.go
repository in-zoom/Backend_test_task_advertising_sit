package validation

import (
	"errors"
	"strconv"
	"strings"
)

func ValidatePrice(price string) (float64, error){
	pricePrepare := preparePrice(price)
	
	if pricePrepare == "" {
    return 0, errors.New("Введите цену")
	}
	praceInt, err := strconv.ParseFloat(pricePrepare, 64)
	if err != nil {
	return 0, errors.New("Задано некорректное значение")
	}
		if praceInt < 0 {
			return 0, errors.New("Цена не может быть отрицательной")
		} 
   return praceInt, nil
}

func preparePrice(imputPrice string) (outputPrace string) {
	priceSpaceRemoval := strings.TrimSpace(imputPrice)
	return priceSpaceRemoval
}