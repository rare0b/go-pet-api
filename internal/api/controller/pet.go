package controller

//go:generate mockgen -source=internal/api/controller/pet.go -destination=internal/mock/controller/pet.go -package=mock

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
	"github.com/rare0b/go-pet-api/internal/api/usecase"
	"net/http"
	"strconv"
	"strings"
)

type PetController interface {
	//UploadImage(w http.ResponseWriter, r *http.Request)
	CreatePet(w http.ResponseWriter, r *http.Request)
	//UpdatePet(w http.ResponseWriter, r *http.Request)
	//GetPetsByStatuses(w http.ResponseWriter, r *http.Request)
	GetPetByID(w http.ResponseWriter, r *http.Request)
	UpdatePetByID(w http.ResponseWriter, r *http.Request)
	DeletePetByID(w http.ResponseWriter, r *http.Request)
}

type petController struct {
	petUsecase usecase.PetUsecase
}

func NewPetController(petUsecase usecase.PetUsecase) PetController {
	return &petController{petUsecase}
}

func (c *petController) CreatePet(w http.ResponseWriter, r *http.Request) {
	pet := &entity.Pet{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	pet, err = c.petUsecase.CreatePet(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *petController) UpdatePet(w http.ResponseWriter, r *http.Request) {
	pet := &entity.Pet{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pet, err = c.petUsecase.UpdatePetByID(pet.ID, pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *petController) GetPetsByStatuses(w http.ResponseWriter, r *http.Request) {
	queryStatus := r.URL.Query().Get("status")
	if queryStatus == "" {
		http.Error(w, `{"error": "Status parameter is required"}`, http.StatusBadRequest)
		return
	}

	statuses := strings.Split(queryStatus, ",")

	pets, err := c.petUsecase.GetPetsByStatuses(statuses)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(pets)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *petController) GetPetByID(w http.ResponseWriter, r *http.Request) {
	petID, err := strconv.ParseInt(chi.URLParam(r, "petID"), 10, 64)
	if err != nil {
		http.Error(w, `{"error": "Invalid petID format"}`, http.StatusBadRequest)
		return
	}

	pet, err := c.petUsecase.GetPetByID(petID)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(pet)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *petController) UpdatePetByID(w http.ResponseWriter, r *http.Request) {
	petID, err := strconv.ParseInt(chi.URLParam(r, "petID"), 10, 64)
	if err != nil {
		http.Error(w, `{"error": "Invalid petID format"}`, http.StatusBadRequest)
		return
	}

	pet := &entity.Pet{}
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pet, err = c.petUsecase.UpdatePetByID(petID, pet)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(pet)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *petController) DeletePetByID(w http.ResponseWriter, r *http.Request) {
	petID, err := strconv.ParseInt(chi.URLParam(r, "petID"), 10, 64)
	if err != nil {
		http.Error(w, `{"error": "Invalid petID format"}`, http.StatusBadRequest)
		return
	}

	err = c.petUsecase.DeletePetByID(petID)
	if err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
