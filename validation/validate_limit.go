package validation

import (
	"Backend_task_advertising_site/DB"
	"database/sql"
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

	numberOfRecords, err := gettingNumberOfRecords(DB.Connect())
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

func gettingNumberOfRecords(db *sql.DB) (numbeOfRecords int, err error) {
	query := "SELECT count(*) FROM ad_table "
	rows, err := db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var numberOfRecords int
	for rows.Next() {
		if err = rows.Scan(&numberOfRecords); err != nil {
			return 0, err
		}
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}
	return numberOfRecords, nil
}
