package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/domain/dbmodel"
)

type PetRepository interface {
	CreatePet(tx *sqlx.Tx, petDBModel *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error)
	CreateCategoryIfNotExist(tx *sqlx.Tx, categoryDBModel *dbmodel.CategoryDBModel) (*dbmodel.CategoryDBModel, error)
	CreateTagsIfNotExist(tx *sqlx.Tx, tagDBModels []*dbmodel.TagDBModel) ([]*dbmodel.TagDBModel, error)
	CreatePetTagsIfNotExist(tx *sqlx.Tx, petTagDBModels []*dbmodel.PetTagDBModel) ([]*dbmodel.PetTagDBModel, error)
	GetPetByID(tx *sqlx.Tx, id int64) (*dbmodel.PetDBModel, error)
	GetCategoryByID(tx *sqlx.Tx, id int64) (*dbmodel.CategoryDBModel, error)
	GetTagsByIDs(tx *sqlx.Tx, ids []int64) ([]*dbmodel.TagDBModel, error)
	GetTagIDsByPetID(tx *sqlx.Tx, petID int64) ([]int64, error)
	UpdatePetByID(tx *sqlx.Tx, id int64, pet *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error)
	DeletePetByID(tx *sqlx.Tx, id int64) error
	DeleteUnusedTags(tx *sqlx.Tx) error
	DeletePetTagsByPetID(tx *sqlx.Tx, petID int64) error
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

func (r *petRepository) GetPetByID(tx *sqlx.Tx, id int64) (*dbmodel.PetDBModel, error) {
	query := `SELECT * FROM pets WHERE pet_id = $1`
	petDBModel := &dbmodel.PetDBModel{}

	err := tx.Get(petDBModel, query, id)
	if err != nil {
		return nil, err
	}

	return petDBModel, nil
}

func (r *petRepository) GetCategoryByID(tx *sqlx.Tx, id int64) (*dbmodel.CategoryDBModel, error) {
	query := `SELECT * FROM categories WHERE category_id = $1`
	categoryDBModel := &dbmodel.CategoryDBModel{}

	err := tx.Get(categoryDBModel, query, id)
	if err != nil {
		return nil, err
	}

	return categoryDBModel, nil
}

func (r *petRepository) GetTagsByIDs(tx *sqlx.Tx, ids []int64) ([]*dbmodel.TagDBModel, error) {
	query := `SELECT * FROM tags WHERE tag_id IN (?)`
	tagDBModels := make([]*dbmodel.TagDBModel, 0, len(ids))

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return nil, err
	}

	query = tx.Rebind(query)

	err = tx.Select(&tagDBModels, query, args...)
	if err != nil {
		return nil, err
	}
	return tagDBModels, nil
}

func (r *petRepository) GetTagIDsByPetID(tx *sqlx.Tx, petID int64) ([]int64, error) {
	query := `SELECT tag_id FROM pet_tags WHERE pet_id = $1`
	var tagIDs []int64

	err := tx.Select(&tagIDs, query, petID)
	if err != nil {
		return nil, err
	}

	return tagIDs, nil
}

func (r *petRepository) UpdatePetByID(tx *sqlx.Tx, id int64, pet *dbmodel.PetDBModel) (*dbmodel.PetDBModel, error) {
	query := `UPDATE pets SET category_id = :category_id, pet_name = :pet_name, photo_urls = :photo_urls, status = :status WHERE pet_id = :pet_id`
	rows, err := tx.NamedQuery(query, pet)
	if err != nil {
		return nil, err
	}

	var newPetDBModel *dbmodel.PetDBModel
	for rows.Next() {
		err = rows.StructScan(newPetDBModel)
		if err != nil {
			return nil, err
		}
	}

	return newPetDBModel, nil
}

func (r *petRepository) DeletePetByID(tx *sqlx.Tx, id int64) error {
	query := `DELETE FROM pets WHERE pet_id = $1`

	_, err := tx.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *petRepository) DeleteUnusedTags(tx *sqlx.Tx) error {
	query := `
		-- 現在pet_tagsに存在しないtagを削除
		WITH unused_tags AS (
		  SELECT tag_id FROM tags
		  EXCEPT
		  SELECT tag_id FROM pet_tags
		)
		DELETE FROM tags WHERE tag_id IN (SELECT tag_id FROM unused_tags);`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *petRepository) DeletePetTagsByPetID(tx *sqlx.Tx, petID int64) error {
	query := `DELETE FROM pet_tags WHERE pet_id = $1`

	_, err := tx.Exec(query, petID)
	if err != nil {
		return err
	}

	return nil
}
