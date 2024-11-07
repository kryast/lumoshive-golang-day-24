package main

import (
	"day-24/database"
	"day-24/handler"
	"day-24/library"
	"day-24/repository"
	"day-24/service"
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
	repo := repository.NewBookRepository(db)
	bookService := service.NewBookService(repo)
	bookHandler := handler.NewBookHandler(*bookService)

	r := chi.NewRouter()
	r.Use(library.MethodForm)

	r.Post("/create-book", bookHandler.CreateBookHandler)
	r.Put("/edit-book/{id}", bookHandler.UpdateBookHandler)
	r.Delete("/delete-book/{id}", bookHandler.DeleteBookHandler)

	r.Get("/dashboard", handler.Home)
	r.Get("/login", handler.FormLogin)
	r.Get("/book-list", bookHandler.BookListHandler)
	r.Get("/create-book", handler.FormCreateBook)
	r.Get("/edit-book/{id}", bookHandler.FormEditBook)
	r.Get("/logout", handler.Logout)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
