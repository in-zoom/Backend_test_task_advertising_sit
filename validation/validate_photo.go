package validation

import (
	"errors"
)

func ValidateFormatPhoto(foto string) error {
	formatPhoto := []string{"jpg", "jpeg", "png", "bmp"}
	for _, currentPhoto := range formatPhoto {
		if foto == currentPhoto {
			return nil
		}
	}
	return errors.New("Выбран неверный формат файла")
}

func TheNumberOfLinksToThePhoto(links []string) error {
	if len(links) > 3 {
		return errors.New("Можно загрузить только 3 файла")
	}
	return nil
}
