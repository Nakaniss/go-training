package usecase

import (
	"todo-app/entity"
	"todo-app/repository"
)

type TodoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepo: repo}
}

func (u *TodoUsecase) CreateTodo(todo *entity.Todo) error {
	if todo.Title == "" {
		return entity.ErrInvalidInput
	}
	return u.todoRepo.Create(todo)
}

func (u *TodoUsecase) GetTodo(id int64) (*entity.Todo, error) {
	return u.todoRepo.GetByID(id)
}

func (u *TodoUsecase) ListTodosByUser(userID int64) ([]*entity.Todo, error) {
	return u.todoRepo.GetByUserID(userID)
}

func (u *TodoUsecase) UpdateTodo(todo *entity.Todo) error {
	return u.todoRepo.Update(todo)
}

func (u *TodoUsecase) DeleteTodo(id int64) error {
	return u.todoRepo.Delete(id)
}
