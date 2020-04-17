package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidFieldsPhotos(t *testing.T) {
	inputFields := []string{"photos"}
	expectedResult := "links[1:3]"
	resultFields, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultFields, expectedResult)
}

func TestValidFieldsDescription(t *testing.T) {
	inputFields := []string{"description"}
	expectedResult := "announcement_text"
	resultFields, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultFields, expectedResult)
}

func TestValidFieldsPhotosAndDescription(t *testing.T) {
	inputFields := []string{"photos", "description"}
	expectedResult := "links[1:3], announcement_text"
	resultFields, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultFields, expectedResult)
}

func TestValidFieldsEmptyArray(t *testing.T) {
	inputFields := []string{}
	expectedResult := "links[1:1]"
	resultFields, err := validation.ValidateFields(inputFields)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultFields, expectedResult)
}
