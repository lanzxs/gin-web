package utils

import uuid "github.com/satori/go.uuid"

func CreateUUID() string {
	id := uuid.NewV4()

	return id.String()
}
