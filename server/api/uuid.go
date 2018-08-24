package api

import (
	uuid "github.com/satori/go.uuid"
)

// IsUUID .
func IsUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
