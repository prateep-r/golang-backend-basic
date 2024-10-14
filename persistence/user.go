package persistence

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	userId    uuid.UUID
	userEmail string
	userName  string
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy *string
	UpdatedAt *time.Time
}
