package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidOneFields(t *testing.T) {
	inputFields := []string{"drop"}
	expectedResult := "Параметр задан неверно"
	_, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidTwoFields(t *testing.T) {
	inputFields := []string{"drop", "delete"}
	expectedResult := "Параметр задан неверно"
	_, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}

func TestInvalidMoreThanThreeFields(t *testing.T) {
	inputFields := []string{"photos", "description", "drop", "delete"}
	expectedResult := "Можно задать не больше двух параметров"
	_, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, expectedResult)
}
