package controller

//go:generate mockgen -source=internal/api/controller/pet.go -destination=internal/mock/controller/pet.go -package=mock

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
	"github.com/rare0b/go-pet-api/internal/api/response"
	"github.com/rare0b/go-pet-api/internal/api/usecase"
	"net/http"
	"strconv"
)

type PetController interface {
	CreatePet(w http.ResponseWriter, r *http.Request)
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

	response, err := json.Marshal(petEntityToPetResponse(pet))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (c *petController) GetPetByID(w http.ResponseWriter, r *http.Request) {
	petID, err := strconv.ParseInt(chi.URLParam(r, "petId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	pet, err := c.petUsecase.GetPetByID(petID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(petEntityToPetResponse(pet))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *petController) UpdatePetByID(w http.ResponseWriter, r *http.Request) {
	petID, err := strconv.ParseInt(chi.URLParam(r, "petId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	pet := &entity.Pet{}
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	pet, err = c.petUsecase.UpdatePetByID(petID, pet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(petEntityToPetResponse(pet))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *petController) DeletePetByID(w http.ResponseWriter, r *http.Request) {
	petID, err := strconv.ParseInt(chi.URLParam(r, "petId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.petUsecase.DeletePetByID(petID)
	if err != nil {
		if err.Error() == "pet not found" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func petEntityToPetResponse(pet *entity.Pet) *response.PetResponse {
	return &response.PetResponse{
		ID:        pet.ID,
		Category:  pet.Category,
		Name:      pet.Name,
		PhotoUrls: pet.PhotoUrls,
		Tags:      pet.Tags,
		Status:    pet.Status,
	}
}
