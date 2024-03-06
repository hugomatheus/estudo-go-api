package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Fulano", "fulano@email.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Fulano", user.Name)
	assert.Equal(t, "fulano@email.com", user.Email)
	assert.NotEqual(t, "123456", user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Fulano", "fulano@email.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123456789"))
	assert.NotEqual(t, "123456", user.Password)
}
