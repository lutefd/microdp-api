package controller

import (
	"net/http"
)

type APIController interface {
	CreateAPI(w http.ResponseWriter, r *http.Request)
	GetAPIByID(w http.ResponseWriter, r *http.Request)
	UpdateAPI(w http.ResponseWriter, r *http.Request)
	DeleteAPI(w http.ResponseWriter, r *http.Request)
	ListAPIs(w http.ResponseWriter, r *http.Request)
}
