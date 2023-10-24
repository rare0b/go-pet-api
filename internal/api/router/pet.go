package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/rare0b/go-pet-api/internal/api/controller"
	"net/http"
)

func NewPetRouter(c controller.PetController) http.Handler {
	r := chi.NewRouter()

	r.Post("/{petId}/uploadImage", c.UploadImage)
	r.Post("", c.CreatePet)
	r.Put("", c.UpdatePet)
	r.Get("/findByStatus", c.GetPetsByStatuses)
	r.Get("/{petId}", c.GetPetByID)
	r.Post("/{petId}", c.UpdatePetByID)
	r.Delete("/{petId}", c.DeletePetByID)

	return r
}
