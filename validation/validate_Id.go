package validation

import (
	"Backend_task_advertising_site/DB"
	"database/sql"
	"errors"
	"strconv"
)

func ValidateId(id string, db *sql.DB) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", errors.New("Задано некорректное значение")
	}
	if idInt <= 0 {
		return "", errors.New("Значение не может быть меньше или равно нулю")
	}

	numberOfRecords, err := DB.GettingNumberOfRecords("max(id)", db)
	if err != nil {
		return "", err
	}

	if idInt > numberOfRecords {
		return strconv.Itoa(numberOfRecords), nil
	} else {
		return strconv.Itoa(idInt), nil
	}
}
