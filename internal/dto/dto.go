package dto

import "github.com/hugomatheus/go-api/pkg/entity"

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductInput struct {
	ID    entity.ID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
