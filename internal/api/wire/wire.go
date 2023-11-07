package wire

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/rare0b/go-pet-api/internal/api/controller"
	"github.com/rare0b/go-pet-api/internal/api/db"
	"github.com/rare0b/go-pet-api/internal/api/repository"
	"github.com/rare0b/go-pet-api/internal/api/router"
	"github.com/rare0b/go-pet-api/internal/api/usecase"
)

func InitializeApp() *chi.Mux {
	wire.Build(
		db.NewDB,
		repository.NewPetRepository,
		usecase.NewPetUsecase,
		controller.NewPetController,
		router.NewPetRouter,
		router.NewMainRouter,
	)
	return nil
}
