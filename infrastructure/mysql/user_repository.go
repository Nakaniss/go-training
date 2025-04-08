package mysql

import (
	"database/sql"
	"todo-app/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *entity.User) (*entity.User, error) {
	result, err := r.DB.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`,
		user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var createdUser entity.User
	row := r.DB.QueryRow(`SELECT id, name, email, password FROM users WHERE id = ?`, id)
	err = row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email, &createdUser.Password)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func (r *UserRepository) GetByID(id int64) (*entity.User, error) {
	row := r.DB.QueryRow(`SELECT id, name, email, password FROM users WHERE id = ?`, id)

	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	_, err := r.DB.Exec(`UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?`,
		user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id = ?`, id)
	return err
}
