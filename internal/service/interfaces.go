package service

import (
	"context"
	"microd-api/internal/models"
)

type APIService interface {
	CreateAPI(ctx context.Context, api models.API) (int64, error)
	GetAPIByID(ctx context.Context, id int64) (models.API, error)
	UpdateAPI(ctx context.Context, api models.API) error
	DeleteAPI(ctx context.Context, id int64) error
	ListAPIs(ctx context.Context) ([]models.API, error)
}
