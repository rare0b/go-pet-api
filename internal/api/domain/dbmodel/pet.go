package dbmodel

type PetDBModel struct {
	PetID      int64    `db:"pet_id"`
	CategoryID int64    `db:"category_id"`
	PetName    string   `db:"pet_name"`
	PhotoUrls  []string `db:"photo_urls"`
	Status     string   `db:"status"`
}
