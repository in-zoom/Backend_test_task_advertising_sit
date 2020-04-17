package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestValidIfTheOffsetValueIsGreaterThanTheNumberOfRecords(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "100"
	expectedResult := "offset 2"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfTheOffsetValueIsLessThanTheNumberOfEntries(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "1"
	expectedResult := "offset 1"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfTheOffsetValueIsExactlyTheNumberOfRecords(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := "3"
	expectedResult := "offset 2"
	rows := sqlmock.NewRows([]string{"id"}).AddRow(3)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	actualResult, err := validation.ValidateOffset(input, db)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

func TestValidIfOffsetIsAnEmptyString(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	inputOffset := []string{"", "        "}
	expectedResult := ""
	for _, currentOffset := range inputOffset {
		resultOffset, err := validation.ValidateOffset(currentOffset, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, resultOffset, expectedResult)
	}
}
