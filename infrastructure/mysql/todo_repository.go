package mysql

import (
	"database/sql"
	"fmt"
	"todo-app/entity"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo *entity.Todo) error {
	_, err := r.DB.Exec(`INSERT INTO todos (user_id, title, done, created_at, updated_at)
                         VALUES (?, ?, ?, NOW(), NOW())`,
		todo.User.ID, todo.Title, todo.Done)
	return err
}

func (r *TodoRepository) GetByID(id int64) (*entity.Todo, error) {
	row := r.DB.QueryRow(`
        SELECT t.id, t.title, t.done, t.created_at, t.updated_at,
               u.id, u.name, u.email, u.password
        FROM todos t
        JOIN users u ON t.user_id = u.id
        WHERE t.id = ?
    `, id)

	var todo entity.Todo
	var user entity.User
	var createdAt, updatedAt []byte
	err := row.Scan(
		&todo.ID, &todo.Title, &todo.Done, &createdAt, &updatedAt,
		&user.ID, &user.Name, &user.Email, &user.Password,
	)
	if err != nil {
		return nil, err
	}

	todo.CreatedAt, err = entity.ParseTime(createdAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created_at: %w", err)
	}
	todo.UpdatedAt, err = entity.ParseTime(updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse updated_at: %w", err)
	}

	todo.User = user
	return &todo, nil
}

func (r *TodoRepository) GetByUserID(userID int64) ([]*entity.Todo, error) {
	rows, err := r.DB.Query(`
        SELECT t.id, t.title, t.done, t.created_at, t.updated_at,
               u.id, u.name, u.email, u.password
        FROM todos t
        JOIN users u ON t.user_id = u.id
        WHERE u.id = ?
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*entity.Todo
	for rows.Next() {
		var todo entity.Todo
		var user entity.User
		err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt,
			&user.ID, &user.Name, &user.Email, &user.Password,
		)
		if err != nil {
			return nil, err
		}
		todo.User = user
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (r *TodoRepository) Update(todo *entity.Todo) error {
	_, err := r.DB.Exec(`
        UPDATE todos
        SET title = ?, done = ?, updated_at = NOW()
        WHERE id = ?
    `, todo.Title, todo.Done, todo.ID)
	return err
}

func (r *TodoRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM todos WHERE id = ?`, id)
	return err
}
