package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestValidIfTheEnteredIdIsGreaterThanWhatIsInTheTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "100"
	expectedResult := "3"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateId(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfTheEnteredIdIsEqualIdInTheTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "1"
	expectedResult := "1"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateId(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
