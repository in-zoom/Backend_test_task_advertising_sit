package validation

import (
	"errors"
	"strings"
)

func ValidateTitle(title string) (string, error) {
	titleSpaceRemoval := strings.TrimSpace(title)
	if titleSpaceRemoval == "" {
		return "", errors.New("Введите название объявления")
	} else if len(titleSpaceRemoval) > 200 {
		return "", errors.New("Название не может содержать болие двухсот символов")
	}
	return titleSpaceRemoval, nil
}
