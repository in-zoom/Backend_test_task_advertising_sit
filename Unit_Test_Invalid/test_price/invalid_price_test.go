package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidiIfThePriceIsNotSet(t *testing.T) {
	input := "                  "
	expectedResult := "Введите цену"
	_, err := validation.ValidatePrice(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfNotANumber(t *testing.T) {
	input := "drop"
	expectedResult := "Задано некорректное значение"
	_, err := validation.ValidatePrice(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfSetToZero(t *testing.T) {
	input := "   0          "
	expectedResult := "Цена не может быть отрицательной или равна нулю"
	_, err := validation.ValidatePrice(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfANegativeInteger(t *testing.T) {
	input := "-1"
	expectedResult := "Цена не может быть отрицательной или равна нулю"
	_, err := validation.ValidatePrice(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidIfANegativeFractionalNumber(t *testing.T) {
	input := "   -1.6          "
	expectedResult := "Цена не может быть отрицательной или равна нулю"
	_, err := validation.ValidatePrice(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}
