package validation

import (
	"fmt"

	"github.com/google/uuid"
)

func ValidateUUID(id string, fieldName string) error {
	if id == "" {
		return fmt.Errorf("%s is empty", fieldName)
	}
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid %s format: %w", fieldName, err)
	}
	return nil
}
