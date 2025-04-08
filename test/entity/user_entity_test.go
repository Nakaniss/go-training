package entity_test

import (
	"testing"
	"todo-app/entity"

	"github.com/stretchr/testify/assert"
)

func TestUser_IsValidEmail(t *testing.T) {
	t.Run("valid email", func(t *testing.T) {
		u := &entity.User{Email: "test@example.com"}
		assert.True(t, u.IsValidEmail())
	})

	t.Run("invalid email without @", func(t *testing.T) {
		u := &entity.User{Email: "invalid-email"}
		assert.False(t, u.IsValidEmail())
	})
}
