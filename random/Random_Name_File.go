package random

import (
	"Backend_task_advertising_site/validation"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

func RandomFileName(input string) (string, error) {
	splitted := strings.Split(input, ".")
	addition := splitted[len(splitted)-1]
	err := validation.ValidateFormatPhoto(addition)
	if err != nil {
		return "", err
	}
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newNameFile := strings.ToLower(fmt.Sprintf("%v", ulid.MustNew(ulid.Timestamp(t), entropy)))
	return newNameFile + "." + addition, nil
}
