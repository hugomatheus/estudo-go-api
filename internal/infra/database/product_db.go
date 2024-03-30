package database

import (
	"github.com/hugomatheus/go-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

type FindAllProductResponse struct {
	Products []entity.Product `json:"products"`
	Total    int64            `json:"total"`
}

func (p *Product) FindAll(page, limit int, sort string) (FindAllProductResponse, error) {
	var products []entity.Product
	var err error
	var total int64

	err = p.DB.Model(&entity.Product{}).Count(&total).Error
	if err != nil {
		return FindAllProductResponse{}, err
	}

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + " " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at" + sort).Find(&products).Error
	}

	if err != nil {
		return FindAllProductResponse{}, err
	}

	return FindAllProductResponse{
		Products: products,
		Total:    total,
	}, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)

	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
