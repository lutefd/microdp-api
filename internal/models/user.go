package models

import (
	"time"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	Avatar       string
	RefreshToken string
	Role         int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
