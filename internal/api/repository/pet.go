package repository

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/domain/dbmodel"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
)

type PetRepository interface {
	CreatePet(tx *sqlx.Tx, petDBModel *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error)
	CreateCategoryIfNotExist(tx *sqlx.Tx, categoryDBModel *dbmodel.CategoryDBModel) (*dbmodel.CategoryDBModel, error)
	CreateTagsIfNotExist(tx *sqlx.Tx, tagDBModels []*dbmodel.TagDBModel) ([]*dbmodel.TagDBModel, error)
	CreatePetTagsIfNotExist(tx *sqlx.Tx, petTagDBModels []*dbmodel.PetTagDBModel) ([]*dbmodel.PetTagDBModel, error)
	GetPetByID(tx *sqlx.Tx, id int64) (*entity.Pet, error)
	UpdatePetByID(tx *sqlx.Tx, id int64, pet *entity.Pet) (*entity.Pet, error)
	DeletePetByID(tx *sqlx.Tx, id int64) error
}

type petRepository struct {
	db *sqlx.DB
}

func NewPetRepository(db *sqlx.DB) PetRepository {
	return &petRepository{db}
}

func (r *petRepository) CreatePet(tx *sqlx.Tx, petDBModel *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error) {
	query := `INSERT INTO pets (pet_id, category_id, pet_name, photo_urls, status) VALUES (:pet_id, :category_id, :pet_name, :photo_urls, :status)`
	rows, err := tx.NamedQuery(query, petDBModel)
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

func (r *petRepository) CreateCategoryIfNotExist(tx *sqlx.Tx, categoryDBModel *dbmodel.CategoryDBModel) (*dbmodel.CategoryDBModel, error) {
	query := `INSERT INTO categories (category_id, category_name) VALUES (:category_id, :category_name) ON CONFLICT DO NOTHING`
	rows, err := tx.NamedQuery(query, categoryDBModel)
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

func (r *petRepository) CreateTagsIfNotExist(tx *sqlx.Tx, tagDBModels []*dbmodel.TagDBModel) ([]*dbmodel.TagDBModel, error) {
	query := `INSERT INTO tags (tag_id, tag_name) VALUES (:tag_id, :tag_name) ON CONFLICT DO NOTHING`
	var NewTagDBModels []*dbmodel.TagDBModel

	//TODO:バルクインサートにしたい
	for _, tagDBModel := range tagDBModels {
		rows, err := tx.NamedQuery(query, tagDBModel)
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

func (r *petRepository) CreatePetTagsIfNotExist(tx *sqlx.Tx, petTagDBModels []*dbmodel.PetTagDBModel) ([]*dbmodel.PetTagDBModel, error) {
	// Updateにも使うのでIfNotExist
	query := `INSERT INTO pet_tags (pet_id, tag_id) VALUES (:pet_id, :tag_id) ON CONFLICT DO NOTHING`
	var NewPetTagDBModels []*dbmodel.PetTagDBModel

	for _, petTagDBModel := range petTagDBModels {
		rows, err := tx.NamedQuery(query, petTagDBModel)
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

func (r *petRepository) GetPetByID(tx *sqlx.Tx, id int64) (*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petRepository.GetPetByID"))
}

func (r *petRepository) UpdatePetByID(tx *sqlx.Tx, id int64, pet *entity.Pet) (*entity.Pet, error) {
	//TODO
	return nil, errors.New(500, fmt.Sprintf("not implemented in petRepository.UpdatePetByID"))
}

func (r *petRepository) DeletePetByID(tx *sqlx.Tx, id int64) error {
	//TODO
	return errors.New(500, fmt.Sprintf("not implemented in petRepository.DeletePetByID"))
}
