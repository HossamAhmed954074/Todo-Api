package models

type Todo struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Completed   bool   `json:"completed" db:"completed"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	DeletedAt   string `json:"deleted_at,omitempty" db:"deleted_at"`
	IsDeleted   bool   `json:"is_deleted" db:"is_deleted"`
}
