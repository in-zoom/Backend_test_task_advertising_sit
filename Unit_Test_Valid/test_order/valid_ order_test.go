package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidOrderAsc(t *testing.T) {
	inputOrder := "asc"
	expectedResult := "asc"
	resultOrder, err := validation.ValidateOrder(inputOrder)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultOrder, expectedResult)
}

func TestValidOrderDesc(t *testing.T) {
	inputOrder := "desc"
	expectedResult := "desc"
	resultOrder, err := validation.ValidateOrder(inputOrder)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultOrder, expectedResult)
}

func TestValidOrderEmptyLine(t *testing.T) {
	inputOrder := "    "
	expectedResult := ""
	resultOrder, err := validation.ValidateOrder(inputOrder)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, resultOrder, expectedResult)
}
