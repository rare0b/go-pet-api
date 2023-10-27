package repository

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/domain/dbmodel"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
)

type PetRepository interface {
	CreatePet(petDBModel *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error)
	CreateCategoryIfNotExist(categoryDBModel *dbmodel.CategoryDBModel) (*dbmodel.CategoryDBModel, error)
	CreateTagsIfNotExist(tagDBModels []*dbmodel.TagDBModel) ([]*dbmodel.TagDBModel, error)
	CreatePetTagsIfNotExist(petTagDBModels []*dbmodel.PetTagDBModel) ([]*dbmodel.PetTagDBModel, error)
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

func (r *petRepository) UploadImage(id string, additionalMetadata string, file string) error {
	//TODO
	return errors.New(500, fmt.Sprintf("not implemented in petRepository.UploadImage"))
}

func (r *petRepository) CreatePet(petDBModel *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error) {
	query := `INSERT INTO pets (pet_id, category_id, pet_name, photo_urls, status) VALUES (:pet_id, :category_id, :pet_name, :photo_urls, :status)`
	rows, err := r.db.NamedQuery(query, petDBModel)
	if err != nil {
		return nil, err
	}

	var NewPetDBModel *dbmodel.PetDBModel
	for rows.Next() {
		err = rows.StructScan(petDBModel)
		if err != nil {
			return nil, err
		}
	}

	return NewPetDBModel, nil
}

func (r *petRepository) CreateCategoryIfNotExist(categoryDBModel *dbmodel.CategoryDBModel) (*dbmodel.CategoryDBModel, error) {
	query := `INSERT INTO categories (category_id, category_name) VALUES (:category_id, :category_name) ON CONFLICT DO NOTHING`
	rows, err := r.db.NamedQuery(query, categoryDBModel)
	if err != nil {
		return nil, err
	}

	var NewCategoryDBModel *dbmodel.CategoryDBModel
	for rows.Next() {
		err = rows.StructScan(NewCategoryDBModel)
		if err != nil {
			return nil, err
		}
	}

	return NewCategoryDBModel, nil
}

func (r *petRepository) CreateTagsIfNotExist(tagDBModels []*dbmodel.TagDBModel) ([]*dbmodel.TagDBModel, error) {
	query := `INSERT INTO tags (tag_id, tag_name) VALUES (:tag_id, :tag_name) ON CONFLICT DO NOTHING`
	var NewTagDBModels []*dbmodel.TagDBModel

	//TODO:バルクインサートにしたい
	for _, tagDBModel := range tagDBModels {
		rows, err := r.db.NamedQuery(query, tagDBModel)
		if err != nil {
			return nil, err
		}

		NewTagDBModel := &dbmodel.TagDBModel{}
		for rows.Next() {
			err = rows.StructScan(tagDBModel)
			if err != nil {
				return nil, err
			}
		}
		NewTagDBModels = append(NewTagDBModels, NewTagDBModel)
	}

	return NewTagDBModels, nil
}

func (r *petRepository) CreatePetTagsIfNotExist(petTagDBModels []*dbmodel.PetTagDBModel) ([]*dbmodel.PetTagDBModel, error) {
	// Updateにも使うのでIfNotExist
	query := `INSERT INTO pet_tags (pet_id, tag_id) VALUES (:pet_id, :tag_id) ON CONFLICT DO NOTHING`
	var NewPetTagDBModels []*dbmodel.PetTagDBModel

	for _, petTagDBModel := range petTagDBModels {
		rows, err := r.db.NamedQuery(query, petTagDBModel)
		if err != nil {
			return nil, err
		}

		NewPetTagDBModel := &dbmodel.PetTagDBModel{}
		for rows.Next() {
			err = rows.StructScan(petTagDBModel)
			if err != nil {
				return nil, err
			}
		}
		NewPetTagDBModels = append(NewPetTagDBModels, NewPetTagDBModel)
	}

	return NewPetTagDBModels, nil
}

func (r *petRepository) GetPetsByStatuses(statuses []string) ([]*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petRepository.GetPetsByStatuses"))
}

func (r *petRepository) GetPetByID(id int64) (*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petRepository.GetPetByID"))
}

func (r *petRepository) UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petRepository.UpdatePetByID"))
}

func (r *petRepository) DeletePetByID(id int64) error {
	//TODO
	return errors.New(500, fmt.Sprintf("not implemented in petRepository.DeletePetByID"))
}
