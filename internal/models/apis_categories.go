package models

import (
	"time"
)

type APICategory struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type APICategoryMapping struct {
	ApiID      int64
	CategoryID int64
}
