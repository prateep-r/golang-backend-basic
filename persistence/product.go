package persistence

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ProductId   uuid.UUID
	ProductName string
	Price       float64
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedBy   string
	UpdatedAt   time.Time
}
