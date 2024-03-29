package database

import (
	"testing"

	"github.com/hugomatheus/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	SetupDBTest(&entity.User{})
	defer CloseDBTest()

	dbInstance := GetDBTest()
	user, _ := entity.NewUser("Fulano", "fulano@email.com", "123456")
	userDB := NewUser(dbInstance)
	err := userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = dbInstance.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
	assert.NotEqual(t, "123456", userFound.Password)
	assert.NotEqual(t, "123456", user.Password)
	assert.Equal(t, user.Password, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	SetupDBTest(&entity.User{})
	defer CloseDBTest()

	user, _ := entity.NewUser("Fulano", "fulano@email.com", "123456")
	userDB := NewUser(GetDBTest())
	err := userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEqual(t, "123456", userFound.Password)
	assert.NotEqual(t, "123456", user.Password)
	assert.Equal(t, user.Password, userFound.Password)
}

func TestWhenNotFindByEmail(t *testing.T) {
	SetupDBTest(&entity.User{})
	defer CloseDBTest()

	userDB := NewUser(GetDBTest())
	userFound, err := userDB.FindByEmail("email@email.com")
	assert.Nil(t, userFound)
	assert.NotNil(t, err)
}
