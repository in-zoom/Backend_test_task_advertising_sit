package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidTitleMoreThanTwoHundredCharacters(t *testing.T) {
	var inputNumberOfCharacters string
	expectedResult := "Название не может содержать болие двухсот символов"
	for i := 0; i < 201; i++ {
		inputNumberOfCharacters += "a"
	}
	_, err := validation.ValidateTitle(inputNumberOfCharacters)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidTitleEmptyLine(t *testing.T) {
	inputTitle := []string{"", "   "}
	expectedResult := "Введите название объявления"
	for _, currentTitle := range inputTitle {
		_, err := validation.ValidateTitle(currentTitle)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}
