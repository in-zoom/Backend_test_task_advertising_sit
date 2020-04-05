package validation

import (
	"Backend_task_advertising_site/DB"
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

func ValidateOffset(offset string, db *sql.DB) (resultOffset string, err error) {
	offsetSpaceRemoval := strings.TrimSpace(offset)

	if offsetSpaceRemoval == "" {
		return "", nil
	}
	offsetInt, err := strconv.Atoi(offsetSpaceRemoval)
	if err != nil {
		return "", errors.New("Задано некорректное значение")
	}

	if offsetInt < 0 {
		return "", errors.New("Значение не может быть отрицательным")
	}

	numberOfRecords, err := DB.GettingNumberOfRecords("count(*)", db)
	if err != nil {
		return "", err
	}

	if offsetInt >= numberOfRecords {
		return "offset" + " " + strconv.Itoa((numberOfRecords - 1)), nil
	} else {
		return "offset" + " " + offset, nil
	}
}
