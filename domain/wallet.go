package domain

import (
	"github.com/google/uuid"
	"time"
)

type Wallet struct {
	ID        uuid.UUID
	UserId    uuid.UUID
	FirstName string
	LastName  string
	Amount    uint64
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status string

const (
	Disabled Status = "disabled"
	Enabled  Status = "enabled"
)
