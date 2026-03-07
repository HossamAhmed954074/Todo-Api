package handelers

import (
	"net/http"
	"strconv"
	"time"
	"todo-api/internal/models"
	"todo-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed"`

}

type TodoHandler struct {
	repo *repository.TodoRepository
}

func NewTodoHandler(repo *repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := &models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsDeleted:   false,
	}
	newTodo, err := h.repo.CreateTodo(c.Request.Context(), todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": newTodo})
}

// delete a todo by id
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	// Convert id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Check if the todo exists
	exists, err := h.repo.IsTodoExists(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check todo existence"})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	// Delete the todo
	err = h.repo.DeleteTodoById(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

// get all todos
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.repo.GetAllTodos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve todos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
}