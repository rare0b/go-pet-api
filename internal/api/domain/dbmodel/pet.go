package dbmodel

import "github.com/lib/pq"

type PetDBModel struct {
	PetID      int64          `db:"pet_id"`
	CategoryID int64          `db:"category_id"`
	PetName    string         `db:"pet_name"`
	PhotoUrls  pq.StringArray `db:"photo_urls"`
	Status     string         `db:"status"`
}
