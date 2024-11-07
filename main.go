package main

import (
	"day-24/database"
	"day-24/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Inisialisasi repository, service, dan handler untuk User

	r := chi.NewRouter()

	r.Get("/login-view", handler.FormLogin)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
