package main

import (
	"day-24/database"
	"day-24/handler"
	"day-24/library"
	"day-24/middleware"
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

	repoBook := repository.NewBookRepository(db)
	bookService := service.NewBookService(repoBook)
	bookHandler := handler.NewBookHandler(bookService)

	repoOrder := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repoOrder)
	orderHandler := handler.NewOrderHandler(orderService)

	repoAdmin := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(repoAdmin)
	adminHandler := handler.NewAdminHandler(adminService)

	repoDashboard := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(repoDashboard)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	r := chi.NewRouter()

	r.Use(library.MethodForm)

	// Rute untuk Halaman Login
	r.Get("/", handler.FormLogin)               // Form login
	r.Post("/login", adminHandler.LoginHandler) // Login

	r.With(middleware.CheckLoginMiddleware).Group(func(r chi.Router) {

		// Rute untuk Admin (Dashboard, dsb)
		r.Get("/dashboard", dashboardHandler.DashboardHandler) // Halaman Dashboard
		r.Get("/book-list", bookHandler.BookListHandler)       // Daftar Buku

		// Rute untuk Buku
		r.Get("/create-book", handler.FormCreateBook)                // Form Create Book
		r.Post("/create-book", bookHandler.CreateBookHandler)        // Create Book
		r.Get("/edit-book/{id}", bookHandler.FormEditBook)           // Form Edit Book
		r.Put("/edit-book/{id}", bookHandler.UpdateBookHandler)      // Update Book
		r.Delete("/delete-book/{id}", bookHandler.DeleteBookHandler) // Delete Book

		// Rute untuk Penjualan (Orders)
		r.Get("/order-list", orderHandler.OrderListHandler)      // Daftar Orders
		r.Post("/create-order", orderHandler.CreateOrderHandler) // Create Order

		// Rute untuk Admin CRUD
		r.Post("/create-admin", adminHandler.CreateAdminHandler) // Create Admin

		// Rute untuk logout
		r.Get("/logout", adminHandler.LogoutHandler) // Logout
		r.Get("/logout-view", handler.Logout)        // Logout view (Redirect ke login)
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
