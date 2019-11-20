package auth

import "github.com/google/uuid"

type User struct {
	ID       int
	UserID   uuid.UUID
	Username string
}
