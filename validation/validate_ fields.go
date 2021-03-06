package validation

import (
	"errors"
	"strings"
)

func ValidateFields(fields []string) (string, error) {

	if len(fields) > 2 {
		return "", errors.New("Можно задать не больше двух параметров")
	} else if len(fields) == 0 {
		return "links[1:1]", nil
	}

	var totalFields string
	for _, currentField := range fields {
		outputField, err := checkField(currentField)
		if err != nil {
			return "", err
		}
		totalFields += " " + outputField + ","
	}
	resultFields := prepareFields(totalFields)
	return resultFields, nil
}

func checkField(inputField string) (string, error) {
	whiteFields := []string{"photos", "description"}
	for _, currentWhiteField := range whiteFields {
		if currentWhiteField == inputField {
			switch inputField {
			case "photos":
				return "links[1:3]", nil
			case "description":
				return "announcement_text", nil
			}
		}
	}
	return "", errors.New("Параметр задан неверно")
}

func prepareFields(inputFields string) string {
	fieldslSpaceRemoval := strings.TrimSpace(inputFields)
	fieldsTrimSuffix := strings.TrimSuffix(fieldslSpaceRemoval, ",")
	return fieldsTrimSuffix
}
