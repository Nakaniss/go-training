package repository

import "todo-app/entity"

type TodoRepository interface {
	Create(todo *entity.Todo) error
	GetByID(id int64) (*entity.Todo, error)
	GetByUserID(userID int64) ([]*entity.Todo, error)
	Update(todo *entity.Todo) error
	Delete(id int64) error
}
