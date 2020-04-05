package validation

import (
	"errors"
	"strings"
)

func ValidateAtribute(attribute string) (string, error) {
	witeAttribute := []string{"price", "date"}
	attributeSpaceRemoval := strings.TrimSpace(attribute)

	if attributeSpaceRemoval == "" {
		return "", nil
	}

	for _, currentAttribute := range witeAttribute {
		if attributeSpaceRemoval == currentAttribute {
			return "ORDER BY" + " " + attribute, nil
		}

	}
	return "", errors.New("Неверный параметр")
}

func ValidateOrder(order string) (string, error) {
	witeOrder := []string{"asc", "desc"}
	orderSpaceRemoval := strings.TrimSpace(order)

	if orderSpaceRemoval == "" {
		return "", nil
	}

	for _, currentOrder := range witeOrder {
		if orderSpaceRemoval == currentOrder {
			return order, nil
		}
	}
	return "", errors.New("Неверный параметр сортировки")
}
