package validation

import (
	"errors"
	"strings"
)

func ValidateAtribute(attribute string) (string, error) {
	witeAttribute := []string{"price", "date"}
	prepareAttribute := prepare(attribute)
	if prepareAttribute == "" {
		return "", nil
	}
	for _, currentAttribute := range witeAttribute {
		if prepareAttribute == currentAttribute {
			return "ORDER BY" + " " + attribute, nil
		}

	}
	return "", errors.New("Неверный параметр")
}

func ValidateOrder(order string) (string, error) {
	witeOrder := []string{"asc", "desc"}
	prepareOrder := prepare(order)
	if prepareOrder == "" {
		return "", nil
	}
	for _, currentOrder := range witeOrder {
		if prepareOrder == currentOrder {
			return order, nil
		}
	}
	return "", errors.New("Неверный параметр сортировки")
}

func prepare(imput string) string {
	emailSpaceRemoval := strings.TrimSpace(imput)
	return emailSpaceRemoval
}
