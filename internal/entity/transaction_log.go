package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransactionLog struct {
	TransactionID uuid.UUID `gorm:"primaryKey"`
	Timestamp     time.Time `gorm:"primaryKey"`
	Status        string
	Value         float64
}
