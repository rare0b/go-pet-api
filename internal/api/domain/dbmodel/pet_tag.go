package dbmodel

type PetTagDBModel struct {
	PetID int64 `db:"pet_id"`
	TagID int64 `db:"tag_id"`
}
