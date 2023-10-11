package usecase

import "github.com/rare0b/go-pet-api/api/repository"

type PetUsecase interface {
}

type petUsecase struct {
	petRepository repository.PetRepository
}

func NewPetUsecase(petRepository repository.PetRepository) PetUsecase {
	return &petUsecase{petRepository}
}
