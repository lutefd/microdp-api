package mocks

import (
	"context"
	"errors"
	"microd-api/internal/models"
	"sync"
)

type MockAPIRepository struct {
	apis   map[int64]models.API
	nextID int64
	mu     sync.Mutex
}

func NewMockAPIRepository() *MockAPIRepository {
	return &MockAPIRepository{
		apis:   make(map[int64]models.API),
		nextID: 1,
	}
}

func (m *MockAPIRepository) CreateAPI(ctx context.Context, api models.API) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	api.ID = m.nextID
	m.apis[api.ID] = api
	m.nextID++
	return api.ID, nil
}

func (m *MockAPIRepository) GetAPIByID(ctx context.Context, id int64) (models.API, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	api, ok := m.apis[id]
	if !ok {
		return models.API{}, errors.New("API not found")
	}
	return api, nil
}

func (m *MockAPIRepository) UpdateAPI(ctx context.Context, api models.API) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.apis[api.ID]; !ok {
		return errors.New("API not found")
	}
	m.apis[api.ID] = api
	return nil
}

func (m *MockAPIRepository) DeleteAPI(ctx context.Context, id int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.apis[id]; !ok {
		return errors.New("API not found")
	}
	delete(m.apis, id)
	return nil
}

func (m *MockAPIRepository) ListAPIs(ctx context.Context) ([]models.API, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	apis := make([]models.API, 0, len(m.apis))
	for _, api := range m.apis {
		apis = append(apis, api)
	}
	return apis, nil
}
