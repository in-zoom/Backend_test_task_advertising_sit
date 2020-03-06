package validation

import (
	"Backend_task_advertising_site/DB"
	"errors"
	"strconv"
)

func ValidateOffset(offset string) (resultOffset string, err error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return "", errors.New("Задано некорректное значение")
	}

	if offsetInt < 0 {
		return "", errors.New("Значение не может быть отрицательным")
	}

	numberOfRecords, err := DB.GettingNumberOfRecords()
	if err != nil {
		return "", err
	}

	if offsetInt >= numberOfRecords {
		return "offset" + " " + strconv.Itoa((numberOfRecords - 1)), nil
	} else {
		return "offset" + " " + offset, nil
	}
	return "", nil
}
