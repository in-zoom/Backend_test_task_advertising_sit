package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAtributePrice(t *testing.T) {
	inputFields := "price"
	expectedResult := "ORDER BY price"
	resultAtribut, err := validation.ValidateAtribute(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultAtribut, expectedResult)
}

func TestValidAtributeDate(t *testing.T) {
	inputFields := "date"
	expectedResult := "ORDER BY date"
	resultAtribut, err := validation.ValidateAtribute(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultAtribut, expectedResult)
}

func TestValidAtributeEmptyLine(t *testing.T) {
	inputFields := "   "
	expectedResult := ""
	resultAtribut, err := validation.ValidateAtribute(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultAtribut, expectedResult)
}
