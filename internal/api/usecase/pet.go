package usecase

// mockgen -source="internal/api/usecase/pet.go" -destination="internal/mock/usecase/pet.go" -package=mock

import (
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
	"github.com/rare0b/go-pet-api/internal/api/repository"
)

type PetUsecase interface {
	UploadImage(id string, additionalMetadata string, file string) error //TODO: シグネチャ不明
	CreatePet(pet *entity.Pet) (*entity.Pet, error)
	GetPetsByStatuses(statuses []string) ([]*entity.Pet, error)
	GetPetByID(id int64) (*entity.Pet, error)
	UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error)
	DeletePetByID(id int64) error
}

type petUsecase struct {
	petRepository repository.PetRepository
}

func NewPetUsecase(petRepository repository.PetRepository) PetUsecase {
	return &petUsecase{petRepository}
}

func (u *petUsecase) UploadImage(id string, additionalMetadata string, file string) error {

}

func (u *petUsecase) CreatePet(pet *entity.Pet) (*entity.Pet, error) {

}

func (u *petUsecase) GetPetsByStatuses(statuses []string) ([]*entity.Pet, error) {

}

func (u *petUsecase) GetPetByID(id int64) (*entity.Pet, error) {

}

func (u *petUsecase) UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error) {

}

func (u *petUsecase) DeletePetByID(id int64) error {

}
