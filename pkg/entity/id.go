package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewId() ID {
	return uuid.New()
}

func ParseId(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
