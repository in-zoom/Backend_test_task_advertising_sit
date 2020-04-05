package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidDescription(t *testing.T) {
	var inputNumberOfCharacters string
	expectedResult := "Текст объявления не может содержать болие тысячи символов"
	for i := 0; i < 1001; i++ {
		inputNumberOfCharacters += "a"
	}
	_, err := validation.ValidateDescription(inputNumberOfCharacters)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}
