package main

import (
	"fmt"
	"github.com/rare0b/go-pet-api/internal/api/controller"
	"github.com/rare0b/go-pet-api/internal/api/db"
	"github.com/rare0b/go-pet-api/internal/api/repository"
	"github.com/rare0b/go-pet-api/internal/api/router"
	"github.com/rare0b/go-pet-api/internal/api/usecase"
	"net/http"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Printf("failed to NewDB: %v\n", err)
		return
	}
	defer db.Close()

	petRepository := repository.NewPetRepository(db)
	petUsecase := usecase.NewPetUsecase(petRepository)
	petController := controller.NewPetController(petUsecase)
	petRouter := router.NewPetRouter(petController)
	mainRouter := router.NewMainRouter(petRouter)

	err = http.ListenAndServe(":8080", mainRouter)
	if err != nil {
		fmt.Printf("failed to ListenAndServe: %v\n", err)
		return
	}
}
