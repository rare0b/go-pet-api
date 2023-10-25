package usecase

//go:generate mockgen -source="internal/api/usecase/pet.go" -destination="internal/mock/usecase/pet.go" -package=mock

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/rare0b/go-pet-api/internal/api/domain/dbmodel"
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
	//TODO
	return errors.New(500, fmt.Sprintf("not implemented in petUsecase.UploadImage"))
}

func (u *petUsecase) CreatePet(pet *entity.Pet) (*entity.Pet, error) {
	categoryDBModel := PetEntityToCategoryDBModel(pet)
	petDBModel := PetEntityToPetDBModel(pet)
	tagDBModels := PetEntityToTagDBModels(pet)
	petTagDBModels := PetEntityToPetTagDBModels(pet)

	categoryDBModel, err := u.petRepository.CreateCategoryIfNotExist(categoryDBModel)
	if err != nil {
		return nil, err
	}

	petDBModel, err = u.petRepository.CreatePet(petDBModel)
	if err != nil {
		return nil, err
	}

	tagDBModels, err = u.petRepository.CreateTagsIfNotExist(tagDBModels)
	if err != nil {
		return nil, err
	}

	petTagDBModels, err = u.petRepository.CreatePetTagsIfNotExist(petTagDBModels)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (u *petUsecase) GetPetsByStatuses(statuses []string) ([]*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petUsecase.GetPetsByStatuses"))
}

func (u *petUsecase) GetPetByID(id int64) (*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petUsecase.GetPetByID"))
}

func (u *petUsecase) UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petUsecase.UpdatePetByID"))
}

func (u *petUsecase) DeletePetByID(id int64) error {
	//TODO
	return errors.New(500, fmt.Sprintf("not implemented in petUsecase.DeletePetByID"))
}

func PetEntityToCategoryDBModel(pet *entity.Pet) *dbmodel.CategoryDBModel {
	return &dbmodel.CategoryDBModel{
		CategoryID:   pet.Category.ID,
		CategoryName: pet.Category.Name,
	}
}

func PetEntityToPetDBModel(pet *entity.Pet) *dbmodel.PetDBModel {
	return &dbmodel.PetDBModel{
		PetID:      pet.ID,
		CategoryID: pet.Category.ID,
		PetName:    *pet.Name,
		PhotoUrls:  pet.PhotoUrls,
		Status:     pet.Status,
	}
}

func PetEntityToTagDBModels(pet *entity.Pet) []*dbmodel.TagDBModel {
	var tagDBModels []*dbmodel.TagDBModel
	for _, tag := range pet.Tags {
		tagDBModels = append(tagDBModels, &dbmodel.TagDBModel{
			TagID:   tag.ID,
			TagName: tag.Name,
		})
	}
	return tagDBModels
}

func PetEntityToPetTagDBModels(pet *entity.Pet) []*dbmodel.PetTagDBModel {
	var petTagDBModels []*dbmodel.PetTagDBModel
	for _, tag := range pet.Tags {
		petTagDBModels = append(petTagDBModels, &dbmodel.PetTagDBModel{
			PetID: pet.ID,
			TagID: tag.ID,
		})
	}
	return petTagDBModels
}
