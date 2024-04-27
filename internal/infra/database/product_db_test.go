package database

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/hugomatheus/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockProductDB struct {
	ProductDB *gorm.DB
}

func NewMockProductDB(db *gorm.DB) *MockProductDB {
	return &MockProductDB{ProductDB: db}
}

func (m *MockProductDB) FindByID(id string) (*entity.Product, error) {
	return nil, errors.New("produto não foi encontrado")
}

func (m *MockProductDB) Delete(id string) error {
	return errors.New("produto não foi deletado")
}

func (m *MockProductDB) Create(product *entity.Product) error {
	return errors.New("produto não foi criado")
}

func TestDeleteProductWithError(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	product, err := entity.NewProduct("Product 1", 1030.9)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productDB := NewProduct(GetDBTest())
	err = productDB.Create(product)
	assert.Nil(t, err)

	mockProductDB := NewMockProductDB(GetDBTest())

	err = mockProductDB.Delete(product.ID.String())
	assert.NotNil(t, err)
	assert.NotEmpty(t, err)
	assert.Error(t, err, "produto não foi deletado")
}

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

func TestFindByIDError(t *testing.T) {
	SetupDBTest(&entity.Product{})
	defer CloseDBTest()

	product, err := entity.NewProduct("Produto 1", 10.4)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productDB := NewProduct(GetDBTest())
	err = productDB.Create(product)
	assert.Nil(t, err)

	mockProductDB := NewMockProductDB(GetDBTest())
	_, err = mockProductDB.FindByID("2333")
	assert.NotNil(t, err)
	assert.NotEmpty(t, err)
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
	assert.Len(t, products.Products, 10)
	assert.Equal(t, "Product 1", products.Products[0].Name)
	assert.Equal(t, "Product 10", products.Products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products.Products, 10)
	assert.Equal(t, "Product 11", products.Products[0].Name)
	assert.Equal(t, "Product 20", products.Products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products.Products, 5)
	assert.Equal(t, "Product 21", products.Products[0].Name)
	assert.Equal(t, "Product 25", products.Products[4].Name)

	products, err = productDB.FindAll(4, 10, "asc")
	assert.Nil(t, err)
	assert.Empty(t, products.Products)

	products, err = productDB.FindAll(0, 0, "")
	assert.Nil(t, err)
	assert.Len(t, products.Products, 25)
	assert.Equal(t, "Product 1", products.Products[0].Name)
	assert.Equal(t, "Product 25", products.Products[24].Name)

	products, err = productDB.FindAll(1, 25, "order")
	assert.Nil(t, err)
	assert.Len(t, products.Products, 25)
	assert.Equal(t, "Product 1", products.Products[0].Name)
	assert.Equal(t, "Product 25", products.Products[24].Name)
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
