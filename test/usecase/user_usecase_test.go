package usecase_test

import (
	"testing"
	"todo-app/entity"
	"todo-app/usecase"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Run("should create a user", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			CreateFunc: func(user *entity.User) (*entity.User, error) {
				return &entity.User{
					ID:   1,
					Name: "testuser",

					Email:    "test@example.com",
					Password: "1v3204[r23qvqwrv23#%#%!",
				}, nil
			},
		}
		uc := usecase.NewUserUsecase(mockRepo)

		input := &entity.User{
			ID:   1,
			Name: "testuser",

			Email:    "test@example.com",
			Password: "1v3204[r23qvqwrv23#%#%!",
		}
		user, err := uc.CreateUser(input)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, "testuser", user.Name)
		assert.Equal(t, "test@example.com", user.Email)
		assert.Equal(t, "1v3204[r23qvqwrv23#%#%!", user.Password)
	})

	t.Run("should return error when name is empty", func(t *testing.T) {
		mockRepo := &mockUserRepository{}
		uc := usecase.NewUserUsecase(mockRepo)

		input := &entity.User{
			Name:     "",
			Email:    "test@example.com",
			Password: "1v3204[r23qvqwrv23#%#%!",
		}
		user, err := uc.CreateUser(input)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, entity.ErrInvalidInput, err)
	})

	t.Run("should return error when email is empty", func(t *testing.T) {
		mockRepo := &mockUserRepository{}
		uc := usecase.NewUserUsecase(mockRepo)

		input := &entity.User{
			Name:     "testuser",
			Email:    "",
			Password: "1v3204[r23qvqwrv23#%#%!",
		}
		user, err := uc.CreateUser(input)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, entity.ErrInvalidInput, err)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("should get a user by ID", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			GetByIDFunc: func(id int64) (*entity.User, error) {
				return &entity.User{
					ID:    id,
					Name:  "testuser",
					Email: "test@example.com",
				}, nil
			},
		}

		uc := usecase.NewUserUsecase(mockRepo)

		user, err := uc.GetUser(1)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, "testuser", user.Name)
		assert.Equal(t, "test@example.com", user.Email)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			GetByIDFunc: func(id int64) (*entity.User, error) {
				return nil, entity.ErrUserNotFound
			},
		}

		uc := usecase.NewUserUsecase(mockRepo)

		user, err := uc.GetUser(999)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, entity.ErrUserNotFound, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("should update a user", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			UpdateFunc: func(user *entity.User) error {
				return nil
			},
		}

		uc := usecase.NewUserUsecase(mockRepo)

		input := &entity.User{
			ID:       1,
			Name:     "updateduser",
			Email:    "updated@example.com",
			Password: "newpassword",
		}
		err := uc.UpdateUser(input)

		assert.NoError(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("should delete a user", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			DeleteFunc: func(id int64) error {
				return nil
			},
		}

		uc := usecase.NewUserUsecase(mockRepo)

		err := uc.DeleteUser(1)

		assert.NoError(t, err)
	})
}
