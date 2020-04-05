package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidTitle(t *testing.T) {
	var inputNumberOfCharacters string
	expectedResult := 200
	for i := 0; i < 200; i++ {
		inputNumberOfCharacters += "a"
	}
	otputDescription, err := validation.ValidateTitle(inputNumberOfCharacters)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, len(otputDescription), expectedResult)
}
