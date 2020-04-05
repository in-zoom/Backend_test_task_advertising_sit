package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInvalidIfTheEnteredIdIsNotADigit(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	inputId := []string{"", "      ", "drop", "delete", "$", "+", "="}
	expectedResult := "Задано некорректное значение"
	for _, currentId := range inputId {
		_, err := validation.ValidateId(currentId, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}

func TestInvalidIfTheEnteredIdIsLessThanOrEqualToZero(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	inputId := []string{"-3", "-2", "-1", "0"}
	expectedResult := "Значение не может быть меньше или равно нулю"
	for _, currentId := range inputId {
		_, err := validation.ValidateId(currentId, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}
