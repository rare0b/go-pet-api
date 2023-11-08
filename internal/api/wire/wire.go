package wire

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rare0b/go-pet-api/internal/api/controller"
	"github.com/rare0b/go-pet-api/internal/api/repository"
	"github.com/rare0b/go-pet-api/internal/api/router"
	"github.com/rare0b/go-pet-api/internal/api/usecase"
)

func Wire(
	db *sqlx.DB,
) *chi.Mux {
	wire.Build(
		repository.NewPetRepository,
		usecase.NewPetUsecase,
		controller.NewPetController,
		router.NewPetRouter,
		router.NewMainRouter,
	)
	return &chi.Mux{}
}
