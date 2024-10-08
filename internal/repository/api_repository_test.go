package repository

import (
	"context"
	"database/sql"
	"microd-api/internal/models"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE apis (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			version TEXT,
			description TEXT,
			documentation_link TEXT,
			forum_reference TEXT,
			apm_link TEXT,
			team TEXT,
			tags TEXT,
			swagger TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		t.Fatalf("Error creating test table: %v", err)
	}

	return db
}

func TestAPIRepository(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewSQLiteAPIRepository(db)
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

		id, err := repo.CreateAPI(ctx, api)
		if err != nil {
			t.Fatalf("Error creating API: %v", err)
		}
		if id <= 0 {
			t.Errorf("Expected positive ID, got %d", id)
		}
	})

	t.Run("GetAPIByID", func(t *testing.T) {
		api, err := repo.GetAPIByID(ctx, 1)
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

		err := repo.UpdateAPI(ctx, api)
		if err != nil {
			t.Fatalf("Error updating API: %v", err)
		}

		updatedAPI, err := repo.GetAPIByID(ctx, 1)
		if err != nil {
			t.Fatalf("Error getting updated API: %v", err)
		}
		if updatedAPI.Name != "Updated API" {
			t.Errorf("Expected updated API name 'Updated API', got '%s'", updatedAPI.Name)
		}
	})

	t.Run("ListAPIs", func(t *testing.T) {
		apis, err := repo.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("Error listing APIs: %v", err)
		}
		if len(apis) != 1 {
			t.Errorf("Expected 1 API, got %d", len(apis))
		}
	})

	t.Run("DeleteAPI", func(t *testing.T) {
		err := repo.DeleteAPI(ctx, 1)
		if err != nil {
			t.Fatalf("Error deleting API: %v", err)
		}

		apis, err := repo.ListAPIs(ctx)
		if err != nil {
			t.Fatalf("Error listing APIs after deletion: %v", err)
		}
		if len(apis) != 0 {
			t.Errorf("Expected 0 APIs after deletion, got %d", len(apis))
		}
	})
}
