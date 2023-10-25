package dbmodel

type TagDBModel struct {
	TagID   int64  `db:"tag_id"`
	TagName string `db:"tag_name"`
}
