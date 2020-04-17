package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidFileFormatPhoto(t *testing.T) {
	inputFormatPhoto := []string{"docx", "rar", "gif", "psd"}
	expectedResult := "Выбран неверный формат файла"
	for _, currentFormatPhoto := range inputFormatPhoto {
		err := validation.ValidateFormatPhoto(currentFormatPhoto)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}

func TestInvalidIfMoreThanThreePhotos(t *testing.T) {
	inputPhoto := []string{"docx", "rar", "gif", "psd"}
	expectedResult := "Можно загрузить только 3 файла"
	err := validation.TheNumberOfLinksToThePhoto(inputPhoto)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}
