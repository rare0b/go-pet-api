package response

import "github.com/rare0b/go-pet-api/internal/api/domain/entity"

type PetResponse struct {
	ID        int64            `json:"id,omitempty"`
	Category  *entity.Category `json:"category,omitempty"`
	Name      *string          `json:"name"`
	PhotoUrls []string         `json:"photoUrls" xml:"photoUrls"`
	Tags      []*entity.Tag    `json:"tags" xml:"tags"`
	Status    string           `json:"status,omitempty"`
}
