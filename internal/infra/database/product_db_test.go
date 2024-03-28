package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hugomatheus/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	product, err := entity.NewProduct("Produto 1", 10.4)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productDB := NewProduct(GetDBTest())
	err = productDB.Create(product)
	assert.Nil(t, err)

}

func TestFindByID(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	product, err := entity.NewProduct("Produto 1", 10.4)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productDB := NewProduct(GetDBTest())
	err = productDB.Create(product)
	assert.Nil(t, err)

	productFind, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.Name, productFind.Name)
}

func TestFindAll(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	dbInstance := GetDBTest()
	for i := 0; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i+1), rand.Float64()*100)
		assert.Nil(t, err)
		assert.NotEmpty(t, product.ID)
		dbInstance.Create(product)
		assert.Nil(t, err)
	}
	productDB := NewProduct(dbInstance)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 5)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 25", products[4].Name)

	products, err = productDB.FindAll(4, 10, "asc")
	assert.Nil(t, err)
	assert.Empty(t, products)
}

func TestUpdateProduct(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	product, err := entity.NewProduct("Product 1", 1030.9)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productDB := NewProduct(GetDBTest())
	err = productDB.Create(product)
	assert.Nil(t, err)

	product.Name = product.Name + "edit"
	err = productDB.Update(product)
	assert.Nil(t, err)

	productFind, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.NotEmpty(t, productFind)
	assert.Equal(t, product.Name, productFind.Name)

}

func TestDelete(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	product, err := entity.NewProduct("Product 1", 1030.9)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productDB := NewProduct(GetDBTest())
	err = productDB.Create(product)
	assert.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	productFind, err := productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Empty(t, productFind)
}
