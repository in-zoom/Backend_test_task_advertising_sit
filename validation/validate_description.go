package validation

import "errors"

func ValidateDescription(description string) (string, error){
	if len(description) > 1000 {
		return "", errors.New("Текст объявления не может содержать болие тысячи символов")
	}
	return description, nil
}