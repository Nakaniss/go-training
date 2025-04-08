package repository

import "todo-app/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	GetByID(id int64) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int64) error
}
