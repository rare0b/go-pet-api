package dbmodel

type CategoryDBModel struct {
	CategoryID   int64  `db:"category_id"`
	CategoryName string `db:"category_name"`
}
