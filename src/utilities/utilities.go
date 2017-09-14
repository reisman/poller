package utilities

import "github.com/google/uuid"

func CreateIdentity() string {
	return uuid.New().String()
}
