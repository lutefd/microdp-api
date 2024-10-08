package models

import (
	"time"
)

type API struct {
	ID                int64
	Name              string
	Version           string
	Description       string
	DocumentationLink string
	ForumReference    string
	Tags              string
	Swagger           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
