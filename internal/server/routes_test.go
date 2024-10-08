package server

import (
	"bytes"
	"encoding/json"
	"microd-api/internal/controller"
	"microd-api/internal/mocks"
	"microd-api/internal/models"
	"microd-api/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIRoutes(t *testing.T) {
	mockRepo := mocks.NewMockAPIRepository()
	apiService := service.NewAPIService(mockRepo)
	apiController := controller.NewAPIController(apiService)
	server := &Server{
		apiController: apiController,
	}

	router := server.RegisterRoutes()

	t.Run("CreateAPI", func(t *testing.T) {
		api := models.API{
			Name:        "Test API",
			Version:     "1.0",
			Description: "Test Description",
		}
		body, _ := json.Marshal(api)
		req, _ := http.NewRequest("POST", "/api/v1/apis", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})

	t.Run("GetAPIByID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/apis/1", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("UpdateAPI", func(t *testing.T) {
		api := models.API{
			Name:        "Updated API",
			Version:     "2.0",
			Description: "Updated Description",
		}
		body, _ := json.Marshal(api)
		req, _ := http.NewRequest("PUT", "/api/v1/apis/1", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("DeleteAPI", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/apis/1", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("ListAPIs", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/apis", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}
