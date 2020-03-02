package validation

import (
	"errors"
	"strings"
)

func ValidateTitle(title string) (string, error){
	titlePrepare := preparePrice(title)
	if titlePrepare == "" {
	return "", errors.New("Введите название объявления")
	} else if len(titlePrepare) > 200 {
	return "", errors.New("Название не может содержать болие двухсот символов")	
	}
  return titlePrepare, nil
}

func prepareTitle(imputTitle string) (outputTitle string) {
	priceSpaceRemoval := strings.TrimSpace(imputTitle)
	return priceSpaceRemoval
}