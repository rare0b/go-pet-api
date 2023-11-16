package usecase

//go:generate mockgen -source="internal/api/usecase/pet.go" -destination="internal/mock/usecase/pet.go" -package=mock

import (
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/domain/dbmodel"
	"github.com/rare0b/go-pet-api/internal/api/domain/entity"
	"github.com/rare0b/go-pet-api/internal/api/repository"
)

type PetUsecase interface {
	CreatePet(pet *entity.Pet) (*entity.Pet, error)
	GetPetByID(id int64) (*entity.Pet, error)
	UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error)
	DeletePetByID(id int64) error
}

type petUsecase struct {
	db            *sqlx.DB
	petRepository repository.PetRepository
}

func NewPetUsecase(db *sqlx.DB, petRepository repository.PetRepository) PetUsecase {
	return &petUsecase{db, petRepository}
}

func (u *petUsecase) CreatePet(pet *entity.Pet) (*entity.Pet, error) {
	categoryDBModel := petEntityToCategoryDBModel(pet)
	petDBModel := petEntityToPetDBModel(pet)
	tagDBModels := petEntityToTagDBModels(pet)
	petTagDBModels := petEntityToPetTagDBModels(pet)

	tx, err := u.db.Beginx()
	if err != nil {
		return nil, err
	}

	categoryDBModel, err = u.petRepository.CreateCategoryIfNotExist(tx, categoryDBModel)
	if err != nil {
		return nil, err
	}

	petDBModel, err = u.petRepository.CreatePet(tx, petDBModel)
	if err != nil {
		return nil, err
	}

	tagDBModels, err = u.petRepository.CreateTagsIfNotExist(tx, tagDBModels)
	if err != nil {
		return nil, err
	}

	petTagDBModels, err = u.petRepository.CreatePetTagsIfNotExist(tx, petTagDBModels)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return pet, nil
}

func (u *petUsecase) GetPetByID(id int64) (*entity.Pet, error) {
	//TODO: 共有ロックにする

	tx, err := u.db.Beginx()
	if err != nil {
		return nil, err
	}

	petDBModel, err := u.petRepository.GetPetByID(tx, id)
	if err != nil {
		return nil, err
	}

	categoryDBModel, err := u.petRepository.GetCategoryByID(tx, petDBModel.CategoryID)
	if err != nil {
		return nil, err
	}

	tagIDs, err := u.petRepository.GetTagIDsByPetID(tx, id)
	if err != nil {
		return nil, err
	}

	tagDBModels, err := u.petRepository.GetTagsByIDs(tx, tagIDs)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return dbModelsToPetEntity(categoryDBModel, petDBModel, tagDBModels), nil
}

func (u *petUsecase) UpdatePetByID(id int64, pet *entity.Pet) (*entity.Pet, error) {
	tx, err := u.db.Beginx()
	if err != nil {
		return nil, err
	}

	// 更新対象のpetが存在するか確認
	_, err = u.petRepository.GetPetByID(tx, id)

	if err != nil {
		pet, err = u.CreatePet(pet)
		if err != nil {
			return nil, err
		}

		return pet, nil
	} else {
		categoryDBModel := petEntityToCategoryDBModel(pet)
		petDBModel := petEntityToPetDBModel(pet)
		tagDBModels := petEntityToTagDBModels(pet)
		petTagDBModels := petEntityToPetTagDBModels(pet)

		tx, err := u.db.Beginx()
		if err != nil {
			return nil, err
		}

		// categoryは他petと共用なので、UpdateではなくCreateのみの想定
		categoryDBModel, err = u.petRepository.CreateCategoryIfNotExist(tx, categoryDBModel)
		if err != nil {
			return nil, err
		}

		petDBModel, err = u.petRepository.UpdatePetByID(tx, id, petDBModel)
		if err != nil {
			return nil, err
		}

		// tagは他petと共用なので、UpdateではなくCreateのみの想定
		tagDBModels, err = u.petRepository.CreateTagsIfNotExist(tx, tagDBModels)
		if err != nil {
			return nil, err
		}

		// Update前後のtagの紐づきを削除→作成
		err = u.petRepository.DeletePetTagsByPetID(tx, id)
		if err != nil {
			return nil, err
		}

		petTagDBModels, err = u.petRepository.CreatePetTagsIfNotExist(tx, petTagDBModels)
		if err != nil {
			return nil, err
		}

		// pet更新後、他のどのpetも持っていないtagは削除
		err = u.petRepository.DeleteUnusedTags(tx)
		if err != nil {
			return nil, err
		}

		err = tx.Commit()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		return pet, nil
	}
}

func (u *petUsecase) DeletePetByID(id int64) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}

	// pet_tagsも参照整合性制約で削除される
	err = u.petRepository.DeletePetByID(tx, id)
	if err != nil {
		return err
	}

	// pet削除後、他のどのpetも持っていないtagは削除
	err = u.petRepository.DeleteUnusedTags(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func petEntityToCategoryDBModel(pet *entity.Pet) *dbmodel.CategoryDBModel {
	return &dbmodel.CategoryDBModel{
		CategoryID:   pet.Category.ID,
		CategoryName: pet.Category.Name,
	}
}

func petEntityToPetDBModel(pet *entity.Pet) *dbmodel.PetDBModel {
	return &dbmodel.PetDBModel{
		PetID:      pet.ID,
		CategoryID: pet.Category.ID,
		PetName:    *pet.Name,
		PhotoUrls:  pet.PhotoUrls,
		Status:     pet.Status,
	}
}

func petEntityToTagDBModels(pet *entity.Pet) []*dbmodel.TagDBModel {
	var tagDBModels []*dbmodel.TagDBModel
	for _, tag := range pet.Tags {
		tagDBModels = append(tagDBModels, &dbmodel.TagDBModel{
			TagID:   tag.ID,
			TagName: tag.Name,
		})
	}
	return tagDBModels
}

func petEntityToPetTagDBModels(pet *entity.Pet) []*dbmodel.PetTagDBModel {
	var petTagDBModels []*dbmodel.PetTagDBModel
	for _, tag := range pet.Tags {
		petTagDBModels = append(petTagDBModels, &dbmodel.PetTagDBModel{
			PetID: pet.ID,
			TagID: tag.ID,
		})
	}
	return petTagDBModels
}

func dbModelsToPetEntity(
	categoryDBModel *dbmodel.CategoryDBModel,
	petDBModel *dbmodel.PetDBModel,
	tagDBModels []*dbmodel.TagDBModel,
) *entity.Pet {
	category := &entity.Category{
		ID:   categoryDBModel.CategoryID,
		Name: categoryDBModel.CategoryName,
	}
	var tags []*entity.Tag
	for _, tagDBModel := range tagDBModels {
		tags = append(tags, &entity.Tag{
			ID:   tagDBModel.TagID,
			Name: tagDBModel.TagName,
		})
	}
	return &entity.Pet{
		ID:        petDBModel.PetID,
		Category:  category,
		Name:      &petDBModel.PetName,
		PhotoUrls: petDBModel.PhotoUrls,
		Tags:      tags,
		Status:    petDBModel.Status,
	}
}
