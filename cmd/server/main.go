package main

import (
	"net/http"

	"github.com/hugomatheus/go-api/configs"
	"github.com/hugomatheus/go-api/internal/entity"
	"github.com/hugomatheus/go-api/internal/infra/database"
	"github.com/hugomatheus/go-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("POST /products", productHandler.CreateProduct)
	http.ListenAndServe(":3333", nil)
}
