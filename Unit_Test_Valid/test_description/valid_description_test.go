package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDescription(t *testing.T) {
	var inputNumberOfCharacters string
	expectedResult := 1000
	for i := 0; i < 1000; i++ {
		inputNumberOfCharacters += "a"
	}
	otputDescription, err := validation.ValidateDescription(inputNumberOfCharacters)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, len(otputDescription), expectedResult)
}
