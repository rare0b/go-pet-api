package repository

import (
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/domain/dbmodel"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
)

type PetRepository interface {
	UploadImage(id string, additionalMetadata string, file string) error //TODO: シグネチャ不明
	CreatePet(petDBModel *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error)
	CreateCategoryIfNotExist(categoryDBModel *dbmodel.CategoryDBModel) (*dbmodel.CategoryDBModel, error)
	CreateTagsIfNotExist(tagDBModels []*dbmodel.TagDBModel) ([]*dbmodel.TagDBModel, error)
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
	query := `INSERT INTO pets (id, category_id, name, photo_urls, status) VALUES (:Category, :Name, :PhotoUrls, :Status)`
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
	query := `INSERT INTO categories (category_id, category_name) VALUES (:CategoryID, :CategoryName) ON CONFLICT DO NOTHING`
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
	query := `INSERT INTO tags (tag_id, pet_id, tag_name) VALUES (:TagID, :PetID, :TagName) ON CONFLICT DO NOTHING`
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
