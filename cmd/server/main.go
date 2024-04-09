package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gutkedu/go-expert-api/configs"
	"github.com/gutkedu/go-expert-api/internal/entity"
	"github.com/gutkedu/go-expert-api/internal/infra/database"
	"github.com/gutkedu/go-expert-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)

	http.ListenAndServe(":"+cfg.WebServerPort, r)
}
