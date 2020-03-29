package store_test

import (
	"github.com/pprisn/http_rest_api/internal/app/model"
	"github.com/pprisn/http_rest_api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindeByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

}
