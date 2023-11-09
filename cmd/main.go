package main

import (
	"fmt"
	"github.com/rare0b/go-pet-api/internal/api/db"
	"github.com/rare0b/go-pet-api/internal/api/wire"
	"net/http"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Printf("failed to NewDB: %v\n", err)
		return
	}
	defer db.Close()

	mainRouter := wire.Wire(db)

	err = http.ListenAndServe(":8080", mainRouter)
	if err != nil {
		fmt.Printf("failed to ListenAndServe: %v\n", err)
		return
	}
}
