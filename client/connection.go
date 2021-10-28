package client

import uuid "github.com/google/uuid"

type Connection struct {
	ID uuid.UUID
}

func NewConnection() *Connection {
	return &Connection{ID: uuid.New()}
}
