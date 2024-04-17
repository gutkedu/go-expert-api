package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/gutkedu/go-expert-api/configs"
	_ "github.com/gutkedu/go-expert-api/docs"
	"github.com/gutkedu/go-expert-api/internal/entity"
	"github.com/gutkedu/go-expert-api/internal/infra/database"
	"github.com/gutkedu/go-expert-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title go expert api example
// @version 1.0
// @description Product api with auth
// @termsOfService http://swagger.io/terms

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", cfg.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		// Protect /products routes with JWT
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		//Routes
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.FetchProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/signin", userHandler.GetJwt)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	fmt.Printf("Server running on port %s", cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
