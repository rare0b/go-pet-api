package repository

import "github.com/rare0b/go-pet-api/api/db"

type PetRepository interface {
}

type petRepository struct {
	db db.DB
}

func NewPetRepository(db db.DB) PetRepository {
	return &petRepository{db}
}
