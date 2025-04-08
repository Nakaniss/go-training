package usecase

import (
	"todo-app/entity"
	"todo-app/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (u *UserUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	if user.Name == "" || user.Email == "" {
		return nil, entity.ErrInvalidInput
	}
	u.userRepo.Create(user)
	return user, nil
}

func (u *UserUsecase) GetUser(id int64) (*entity.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserUsecase) UpdateUser(user *entity.User) error {
	return u.userRepo.Update(user)
}

func (u *UserUsecase) DeleteUser(id int64) error {
	return u.userRepo.Delete(id)
}
