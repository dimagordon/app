package identity

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID `json:"id"`
}
