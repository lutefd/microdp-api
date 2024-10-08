package service

import (
	"context"
	"microd-api/internal/mocks"
	"microd-api/internal/models"
	"testing"
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
			t.Fatalf("Error creating API: %v", err)
		}
		if id <= 0 {
			t.Errorf("Expected positive ID, got %d", id)
		}
	})

	t.Run("GetAPIByID", func(t *testing.T) {
		api, err := service.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("Error getting API: %v", err)
		}
		if api.Name != "Test API" {
			t.Errorf("Expected API name 'Test API', got '%s'", api.Name)
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
			t.Fatalf("Error updating API: %v", err)
		}

		updatedAPI, err := service.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("Error getting updated API: %v", err)
		}
		if updatedAPI.Name != "Updated API" {
			t.Errorf("Expected updated API name 'Updated API', got '%s'", updatedAPI.Name)
		}
	})

	t.Run("ListAPIs", func(t *testing.T) {
		apis, err := service.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("Error listing APIs: %v", err)
		}
		if len(apis) != 1 {
			t.Errorf("Expected 1 API, got %d", len(apis))
		}
	})

	t.Run("DeleteAPI", func(t *testing.T) {
		err := service.DeleteAPI(ctx, 1)
		if err != nil {
			t.Fatalf("Error deleting API: %v", err)
		}

		apis, err := service.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("Error listing APIs after deletion: %v", err)
		}
		if len(apis) != 0 {
			t.Errorf("Expected 0 APIs after deletion, got %d", len(apis))
		}
	})
}
