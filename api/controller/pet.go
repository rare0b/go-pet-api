package controller

import "github.com/rare0b/go-pet-api/api/usecase"

type PetController interface {
}

type petController struct {
	petUsecase usecase.PetUsecase
}

func NewPetController(petUsecase usecase.PetUsecase) PetController {
	return &petController{petUsecase}
}
