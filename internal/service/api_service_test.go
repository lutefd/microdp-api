package service

import (
	"context"
	"microd-api/internal/cache"
	"microd-api/internal/mocks"
	"microd-api/internal/models"
	"reflect"
	"testing"
	"time"
)

func TestAPIService(t *testing.T) {
	mockRepo := mocks.NewMockAPIRepository()
	service := NewAPIService(mockRepo)
	ctx := context.Background()

	t.Run("CreateAPI", func(t *testing.T) {
		api := models.API{
			Name:              "Test API",
			Version:           "1.0",
			Description:       "Test Description",
			DocumentationLink: "http://docs.example.com",
			ForumReference:    "http://forum.example.com",
			ApmLink:           "http://apm.example.com",
			Team:              "Test Team",
			Tags:              "test,api",
			Swagger:           "http://swagger.example.com",
		}

		id, err := service.CreateAPI(ctx, api)
		if err != nil {
			t.Fatalf("error creating API: %v", err)
		}
		if id <= 0 {
			t.Errorf("expected positive ID, got %d", id)
		}
	})

	t.Run("GetAPIByID", func(t *testing.T) {
		api, err := service.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("error getting API: %v", err)
		}
		if api.Name != "Test API" {
			t.Errorf("expected API name 'Test API', got '%s'", api.Name)
		}

		cachedAPI, err := service.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("error getting API from cache: %v", err)
		}
		if !reflect.DeepEqual(api, cachedAPI) {
			t.Errorf("cached API does not match original API")
		}

		modifiedAPI := api
		modifiedAPI.Name = "Modified API"
		mockRepo.UpdateAPI(ctx, modifiedAPI)

		cachedAPI, err = service.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("error getting API from cache: %v", err)
		}
		if cachedAPI.Name != "Test API" {
			t.Errorf("expected cached API name 'Test API', got '%s'", cachedAPI.Name)
		}
	})

	t.Run("UpdateAPI", func(t *testing.T) {
		api := models.API{
			ID:                1,
			Name:              "Updated API",
			Version:           "2.0",
			Description:       "Updated Description",
			DocumentationLink: "http://updated-docs.example.com",
			ForumReference:    "http://updated-forum.example.com",
			ApmLink:           "http://updated-apm.example.com",
			Team:              "Updated Team",
			Tags:              "updated,api",
			Swagger:           "http://updated-swagger.example.com",
		}

		err := service.UpdateAPI(ctx, api)
		if err != nil {
			t.Fatalf("error updating API: %v", err)
		}

		updatedAPI, err := service.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("error getting updated API: %v", err)
		}
		if updatedAPI.Name != "Updated API" {
			t.Errorf("expected updated API name 'Updated API', got '%s'", updatedAPI.Name)
		}
	})

	t.Run("ListAPIs", func(t *testing.T) {
		apis, err := service.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("error listing APIs: %v", err)
		}
		if len(apis) != 1 {
			t.Errorf("expected 1 API, got %d", len(apis))
		}

		cachedAPIs, err := service.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("error listing APIs from cache: %v", err)
		}
		if !reflect.DeepEqual(apis, cachedAPIs) {
			t.Errorf("cached APIs do not match original APIs")
		}

		newAPI := models.API{Name: "New API"}
		_, err = service.CreateAPI(ctx, newAPI)
		if err != nil {
			t.Fatalf("error creating new API: %v", err)
		}

		updatedAPIs, err := service.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("error listing updated APIs: %v", err)
		}
		if len(updatedAPIs) != 2 {
			t.Errorf("expected 2 APIs after adding new one, got %d", len(updatedAPIs))
		}
	})

	t.Run("DeleteAPI", func(t *testing.T) {
		err := service.DeleteAPI(ctx, 1)
		if err != nil {
			t.Fatalf("error deleting API: %v", err)
		}

		_, err = service.GetAPIByID(ctx, 1)
		if err == nil {
			t.Error("expected error when fetching deleted API, got nil")
		}

		apis, err := service.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("error listing APIs after deletion: %v", err)
		}
		if len(apis) != 1 {
			t.Errorf("expected 1 API after deletion, got %d", len(apis))
		}
	})

	t.Run("CacheExpiration", func(t *testing.T) {
		mockRepo := mocks.NewMockAPIRepository()
		shortCache := cache.NewCache(100 * time.Millisecond)
		shortCacheService := &DefaultAPIService{
			repo:  mockRepo,
			cache: shortCache,
		}

		ctx := context.Background()

		api := models.API{Name: "Expiring API"}
		id, err := shortCacheService.CreateAPI(ctx, api)
		if err != nil {
			t.Fatalf("error creating API: %v", err)
		}

		fetchedAPI, err := shortCacheService.GetAPIByID(ctx, id)
		if err != nil {
			t.Fatalf("error fetching API: %v", err)
		}
		if fetchedAPI.Name != "Expiring API" {
			t.Errorf("expected API name 'Expiring API', got '%s'", fetchedAPI.Name)
		}

		modifiedAPI := fetchedAPI
		modifiedAPI.Name = "Modified Expiring API"
		err = mockRepo.UpdateAPI(ctx, modifiedAPI)
		if err != nil {
			t.Fatalf("error updating API in repository: %v", err)
		}

		time.Sleep(150 * time.Millisecond)

		refetchedAPI, err := shortCacheService.GetAPIByID(ctx, id)
		if err != nil {
			t.Fatalf("error refetching API: %v", err)
		}
		if refetchedAPI.Name != "Modified Expiring API" {
			t.Errorf("expected API name 'Modified Expiring API' after cache expiration, got '%s'", refetchedAPI.Name)
		}

		if reflect.DeepEqual(fetchedAPI, refetchedAPI) {
			t.Errorf("refetched API should be different from originally fetched API after cache expiration")
		}
	})
}
