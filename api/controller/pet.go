package controller

import (
	"encoding/json"
	"github.com/rare0b/go-pet-api/api/domain/entity"
	"github.com/rare0b/go-pet-api/api/usecase"
	"net/http"
)

type PetController interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
	CreatePet(w http.ResponseWriter, r *http.Request)
	UpdatePet(w http.ResponseWriter, r *http.Request)
	GetPetsByStatus(w http.ResponseWriter, r *http.Request)
	GetPetByID(w http.ResponseWriter, r *http.Request)
	UpdatePetByID(w http.ResponseWriter, r *http.Request)
	DeletePetByID(w http.ResponseWriter, r *http.Request)
}

type petController struct {
	petUsecase usecase.PetUsecase
}

func NewPetController(petUsecase *usecase.PetUsecase) PetController {
	return &petController{petUsecase}
}

func (c *petController) UploadImage(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (c *petController) CreatePet(w http.ResponseWriter, r *http.Request) {
	pet := &entity.Pet{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(pet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pet, err = c.petUsecase.CreatePet(pet)
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

func (c *petController) GetPetsByStatus(w http.ResponseWriter, r *http.Request) {
	//TODO: query param
}

func (c *petController) GetPetByID(w http.ResponseWriter, r *http.Request) {
	//TODO: path param
}

func (c *petController) UpdatePetByID(w http.ResponseWriter, r *http.Request) {
	//TODO: path param
}

func (c *petController) DeletePetByID(w http.ResponseWriter, r *http.Request) {
	//TODO: path param
}
