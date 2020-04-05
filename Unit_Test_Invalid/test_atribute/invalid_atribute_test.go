package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidAtribute(t *testing.T) {
	inputAtribute := []string{"drop", "delete"}
	expectedResult := "Неверный параметр"
	for _, currentAtribute := range inputAtribute {
		_, err := validation.ValidateAtribute(currentAtribute)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}
