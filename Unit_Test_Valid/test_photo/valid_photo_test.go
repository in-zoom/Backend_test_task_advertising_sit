package validation

import (
	"Backend_task_advertising_site/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidFileFormatPhoto(t *testing.T) {
	inputFormatPhoto := []string{"jpg", "jpeg", "png", "bmp"}
	for _, currentFormatPhoto := range inputFormatPhoto {
		err := validation.ValidateFormatPhoto(currentFormatPhoto)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, nil)
	}
}

func TestValidIfOnePhoto(t *testing.T) {
	inputPhoto := []string{"jpg"}
	err := validation.TheNumberOfLinksToThePhoto(inputPhoto)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, nil)
}

func TestValidIfTwoPhotos(t *testing.T) {
	inputPhotos := []string{"jpg", "jpeg"}
	err := validation.TheNumberOfLinksToThePhoto(inputPhotos)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, nil)
}

func TestValidifThreePhotos(t *testing.T) {
	inputPhotos := []string{"jpg", "jpeg", "png"}
	err := validation.TheNumberOfLinksToThePhoto(inputPhotos)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, err, nil)
}
