package service

import (
	"context"
	"encoding/json"
	"fmt"
	"microd-api/internal/cache"
	"microd-api/internal/models"
	"microd-api/internal/repository"
	"time"
)

type DefaultAPIService struct {
	repo  repository.APIRepository
	cache *cache.Cache
}

func NewAPIService(repo repository.APIRepository) APIService {
	return &DefaultAPIService{
		repo:  repo,
		cache: cache.NewCache(5 * time.Minute),
	}
}

func (s *DefaultAPIService) CreateAPI(ctx context.Context, api models.API) (int64, error) {
	id, err := s.repo.CreateAPI(ctx, api)
	if err != nil {
		return 0, err
	}

	s.cache.Clear()

	return id, nil
}

func (s *DefaultAPIService) GetAPIByID(ctx context.Context, id int64) (models.API, error) {
	cacheKey := fmt.Sprintf("api:%d", id)

	if cachedData, ok := s.cache.Get(cacheKey); ok {
		var api models.API
		err := json.Unmarshal(cachedData, &api)
		if err == nil {
			return api, nil
		}
	}

	api, err := s.repo.GetAPIByID(ctx, id)
	if err != nil {
		return models.API{}, err
	}

	cachedData, _ := json.Marshal(api)
	s.cache.Set(cacheKey, cachedData)

	return api, nil
}

func (s *DefaultAPIService) UpdateAPI(ctx context.Context, api models.API) error {
	err := s.repo.UpdateAPI(ctx, api)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("api:%d", api.ID)
	s.cache.Clear()

	cachedData, _ := json.Marshal(api)
	s.cache.Set(cacheKey, cachedData)

	return nil
}

func (s *DefaultAPIService) DeleteAPI(ctx context.Context, id int64) error {
	err := s.repo.DeleteAPI(ctx, id)
	if err != nil {
		return err
	}

	s.cache.Clear()

	return nil
}

func (s *DefaultAPIService) ListAPIs(ctx context.Context) ([]models.API, error) {
	cacheKey := "api:list"

	if cachedData, ok := s.cache.Get(cacheKey); ok {
		var apis []models.API
		err := json.Unmarshal(cachedData, &apis)
		if err == nil {
			return apis, nil
		}
	}

	apis, err := s.repo.ListAPIs(ctx)
	if err != nil {
		return nil, err
	}

	cachedData, _ := json.Marshal(apis)
	s.cache.Set(cacheKey, cachedData)

	return apis, nil
}
