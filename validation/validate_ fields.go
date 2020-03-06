package validation

import (
	"strings"
)

func ValidateFields(fields string) string {
	emailSpaceRemoval := strings.TrimSpace(fields)
		
	if emailSpaceRemoval == "fields"{
			return emailSpaceRemoval
		}
	return ""
}