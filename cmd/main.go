package main

import (
	"fmt"
	"github.com/rare0b/go-pet-api/internal/api/wire"
	"net/http"
)

func main() {
	mainRouter := wire.InitializeApp()

	err := http.ListenAndServe(":8080", mainRouter)
	if err != nil {
		fmt.Printf("failed to ListenAndServe: %v\n", err)
		return
	}
}
