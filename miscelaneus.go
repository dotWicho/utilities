package utilities

import "github.com/google/uuid"

// IsValidUUID just validate if an uuid is valid (Version 1 or Version 4)
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
