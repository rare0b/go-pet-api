package dbmodel

type TagDBModel struct {
	TagID   int64  `db:"tag_id"`
	PetID   int64  `db:"pet_id"`
	TagName string `db:"tag_name"`
}
