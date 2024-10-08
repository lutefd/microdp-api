package service

import (
	"context"
	"microd-api/internal/models"
	"microd-api/internal/repository"
)

type DefaultAPIService struct {
	repo repository.APIRepository
}

func NewAPIService(repo repository.APIRepository) APIService {
	return &DefaultAPIService{repo: repo}
}

func (s *DefaultAPIService) CreateAPI(ctx context.Context, api models.API) (int64, error) {
	// TODO: add business logic
	return s.repo.CreateAPI(ctx, api)
}

func (s *DefaultAPIService) GetAPIByID(ctx context.Context, id int64) (models.API, error) {
	// TODO: add business logic
	return s.repo.GetAPIByID(ctx, id)
}

func (s *DefaultAPIService) UpdateAPI(ctx context.Context, api models.API) error {
	// TODO: add business logic
	return s.repo.UpdateAPI(ctx, api)
}

func (s *DefaultAPIService) DeleteAPI(ctx context.Context, id int64) error {
	// TODO: add business logic

	return s.repo.DeleteAPI(ctx, id)
}

func (s *DefaultAPIService) ListAPIs(ctx context.Context) ([]models.API, error) {
	// TODO: add business logic
	return s.repo.ListAPIs(ctx)
}
