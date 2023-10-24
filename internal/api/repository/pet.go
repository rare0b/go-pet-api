package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
)

type PetRepository interface {
	UploadImage(id string, additionalMetadata string, file string) error //TODO: シグネチャ不明
	CreatePet(pet *entity.Pet) (*entity.Pet, error)
	GetPetsByStatuses(statuses []string) ([]*entity.Pet, error)
	GetPetByID(id int64) (*entity.Pet, error)
	UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error)
	DeletePetByID(id int64) error
}

type petRepository struct {
	db *sqlx.DB
}

func NewPetRepository(db *sqlx.DB) PetRepository {
	return &petRepository{db}
}

func UploadImage(id string, additionalMetadata string, file string) error {
	//TODO
	return nil
}

func (r *petRepository) CreatePet(pet *entity.Pet) (*entity.Pet, error) {

}

func (r *petRepository) GetPetsByStatuses(statuses []string) ([]*entity.Pet, error) {

}

func (r *petRepository) GetPetByID(id int64) (*entity.Pet, error) {

}

func (r *petRepository) UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error) {

}

func (r *petRepository) DeletePetByID(id int64) error {

}
