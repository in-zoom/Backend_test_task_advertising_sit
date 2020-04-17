package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidIfAnInteger(t *testing.T) {
	inputInteger := "1"
	var expectedResult float64 = 1
	resultNumber, err := validation.ValidatePrice(inputInteger)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultNumber, expectedResult)
}

func TestValidIfAFractionalNumber(t *testing.T) {
	inputAFractionalNumber := "1.6"
	expectedResult := 1.6
	resultNumber, err := validation.ValidatePrice(inputAFractionalNumber)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultNumber, expectedResult)
}

func TestValidIfNumberWithSpaces(t *testing.T) {
	inputNumber := "   1.6          "
	expectedResult := 1.6
	resultNumber, err := validation.ValidatePrice(inputNumber)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultNumber, expectedResult)
}
