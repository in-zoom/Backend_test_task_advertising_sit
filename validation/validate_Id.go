package validation

import (
	"Backend_task_advertising_site/DB"
	"errors"
	"strconv"
)

func ValidateId(id string) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", errors.New("Задано некорректное значение")
	}
	if idInt <= 0 {
		return "", errors.New("Значение не может быть меньше или равно нулю")
	}
	numberOfRecords, err := DB.GettingNumberOfRecords()
	if err != nil {
		return "", err
	}

	if idInt > numberOfRecords {
		return strconv.Itoa(numberOfRecords), nil
	} else {
		return strconv.Itoa(idInt), nil
	}
	return "", nil
}
