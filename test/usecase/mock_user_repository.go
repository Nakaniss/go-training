package usecase_test

import "todo-app/entity"

type mockUserRepository struct {
	CreateFunc  func(user *entity.User) (*entity.User, error)
	GetByIDFunc func(id int64) (*entity.User, error)
	UpdateFunc  func(user *entity.User) error
	DeleteFunc  func(id int64) error
}

func (m *mockUserRepository) Create(user *entity.User) (*entity.User, error) {
	return m.CreateFunc(user)
}

func (m *mockUserRepository) GetByID(id int64) (*entity.User, error) {
	return m.GetByIDFunc(id)
}

func (m *mockUserRepository) Update(user *entity.User) error {
	return m.UpdateFunc(user)
}

func (m *mockUserRepository) Delete(id int64) error {
	return m.DeleteFunc(id)
}
