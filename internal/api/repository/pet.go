package repository

import (
	"github.com/jmoiron/sqlx"
)

type PetRepository interface {
}

type petRepository struct {
	db *sqlx.DB
}

func NewPetRepository(db *sqlx.DB) PetRepository {
	return &petRepository{db}
}
