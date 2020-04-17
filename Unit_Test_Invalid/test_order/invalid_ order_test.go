package validation

import (
	"backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidOrder(t *testing.T) {
	inputOrder := []string{"drop", "delete"}
	expectedResult := "Неверный параметр сортировки"
	for _, currentOrder := range inputOrder {
		_, err := validation.ValidateOrder(currentOrder)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}
