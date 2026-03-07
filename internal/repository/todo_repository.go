package repository

import (
	"context"
	"fmt"
	"time"
	"todo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `INSERT INTO todos (title, description, completed, created_at, updated_at, is_deleted) 
          VALUES ($1, $2, $3, $4, $5, $6) 
          RETURNING id, title, description, completed, created_at, updated_at, is_deleted`
	var newTodo models.Todo
	err := r.db.QueryRow(ctx, query, todo.Title, todo.Description, todo.Completed, todo.CreatedAt, todo.UpdatedAt, todo.IsDeleted).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Description, &newTodo.Completed, &newTodo.CreatedAt, &newTodo.UpdatedAt, &newTodo.IsDeleted)
	return &newTodo, err
}

// if todo is exists
func (r *TodoRepository) IsTodoExists(ctx context.Context, id int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `SELECT COUNT(1) FROM todos WHERE id = $1 AND is_deleted = false`
	var count int
	err := r.db.QueryRow(ctx, query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TodoRepository) GetTodoByID(ctx context.Context, id int) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `SELECT id, title, description, completed, created_at, updated_at, deleted_at, is_deleted 
			  FROM todos WHERE id = $1 AND is_deleted = false`
	var todo models.Todo
	err := r.db.QueryRow(ctx, query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt, &todo.IsDeleted)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) GetAllTodos(ctx context.Context) ([]*models.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `SELECT id, title, description, completed, created_at, updated_at, deleted_at, is_deleted 
			  FROM todos WHERE is_deleted = false`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
    fmt.Println("Query executed successfully, processing results... %v ", rows.Err())
	var todos []*models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt, &todo.IsDeleted)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (r *TodoRepository) UpdateTodo(ctx context.Context, todo *models.Todo) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `UPDATE todos SET title = $1, description = $2, completed = $3, updated_at = $4 
			  WHERE id = $5 AND is_deleted = false`
	_, err := r.db.Exec(ctx, query, todo.Title, todo.Description, todo.Completed, todo.UpdatedAt, todo.ID)
	return err
}

func (r *TodoRepository) DeleteTodoById(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `UPDATE todos SET is_deleted = true, deleted_at = $1 WHERE id = $2 AND is_deleted = false`
	res, err := r.db.Exec(ctx, query, time.Now(), id)
	if err != nil {
		return err
	}

	// بدلاً من الاستعلام مرتين، نتحقق هل تم تحديث شيء فعلاً؟
	if res.RowsAffected() == 0 {
		return nil // أو ارمِ خطأ "not found" مخصص
	}
	return nil
}

func (r *TodoRepository) DeleteAllTodos(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := `UPDATE todos SET is_deleted = true, deleted_at = $1 WHERE is_deleted = false`
	res, err := r.db.Exec(ctx, query, time.Now())
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return nil // أو ارمِ خطأ "not found" مخصص
	}

	return err
}
