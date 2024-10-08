package controller

import (
	"encoding/json"
	"microd-api/internal/models"
	"microd-api/internal/service"
	"microd-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type DefaultAPIController struct {
	service service.APIService
}

func NewAPIController(service service.APIService) APIController {
	return &DefaultAPIController{service: service}
}

func (c *DefaultAPIController) CreateAPI(w http.ResponseWriter, r *http.Request) {
	var api models.API
	err := json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	id, err := c.service.CreateAPI(r.Context(), api)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating API")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]int64{"id": id})
}
func (c *DefaultAPIController) GetAPIByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid API ID")
		return
	}

	api, err := c.service.GetAPIByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "API not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, api)
}

func (c *DefaultAPIController) UpdateAPI(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid API ID")
		return
	}

	var api models.API
	err = json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	api.ID = id

	err = c.service.UpdateAPI(r.Context(), api)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error updating API")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "API updated successfully"})
}

func (c *DefaultAPIController) DeleteAPI(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid API ID")
		return
	}

	err = c.service.DeleteAPI(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error deleting API")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "API deleted successfully"})
}

func (c *DefaultAPIController) ListAPIs(w http.ResponseWriter, r *http.Request) {
	apis, err := c.service.ListAPIs(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error listing APIs")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, apis)
}
