package controller

import (
	"bytes"
	"encoding/json"
	"microd-api/internal/mocks"
	"microd-api/internal/models"
	"microd-api/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestAPIController(t *testing.T) {
	mockRepo := mocks.NewMockAPIRepository()
	apiService := service.NewAPIService(mockRepo)
	controller := NewAPIController(apiService)

	t.Run("CreateAPI", func(t *testing.T) {
		api := models.API{
			Name:        "Test API",
			Version:     "1.0",
			Description: "Test Description",
		}
		body, _ := json.Marshal(api)
		req, _ := http.NewRequest("POST", "/apis", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.CreateAPI(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var response map[string]int64
		json.Unmarshal(rr.Body.Bytes(), &response)
		if _, exists := response["id"]; !exists {
			t.Errorf("response doesn't contain id field")
		}
	})

	t.Run("CreateAPI_InvalidJSON", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/apis", bytes.NewBufferString("invalid json"))
		rr := httptest.NewRecorder()

		controller.CreateAPI(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("GetAPIByID", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/apis/1", nil)
		rr := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Get("/apis/{id}", controller.GetAPIByID)
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response models.API
		json.Unmarshal(rr.Body.Bytes(), &response)
		if response.Name != "Test API" {
			t.Errorf("handler returned unexpected body: got %v want %v", response.Name, "Test API")
		}
	})

	t.Run("GetAPIByID_NotFound", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/apis/999", nil)
		rr := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Get("/apis/{id}", controller.GetAPIByID)
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})

	t.Run("UpdateAPI", func(t *testing.T) {
		api := models.API{
			Name:        "Updated API",
			Version:     "2.0",
			Description: "Updated Description",
		}
		body, _ := json.Marshal(api)
		req, _ := http.NewRequest("PUT", "/apis/1", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Put("/apis/{id}", controller.UpdateAPI)
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("UpdateAPI_InvalidJSON", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/apis/1", bytes.NewBufferString("invalid json"))
		rr := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Put("/apis/{id}", controller.UpdateAPI)
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("DeleteAPI", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/apis/1", nil)
		rr := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Delete("/apis/{id}", controller.DeleteAPI)
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("DeleteAPI_NotFound", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/apis/999", nil)
		rr := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Delete("/apis/{id}", controller.DeleteAPI)
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
		}
	})

	t.Run("ListAPIs", func(t *testing.T) {
		api := models.API{
			Name:        "Test API",
			Version:     "1.0",
			Description: "Test Description",
		}
		body, _ := json.Marshal(api)
		req, _ := http.NewRequest("POST", "/apis", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		controller.CreateAPI(rr, req)

		req, _ = http.NewRequest("GET", "/apis", nil)
		rr = httptest.NewRecorder()

		controller.ListAPIs(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response []models.API
		json.Unmarshal(rr.Body.Bytes(), &response)
		if len(response) != 1 {
			t.Errorf("handler returned unexpected number of apis: got %v want %v", len(response), 1)
		}
	})
}
