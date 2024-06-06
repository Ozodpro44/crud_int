package repoi

import (
	"app/models"
	"context"

	"github.com/google/uuid"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	UpdateUser(ctx context.Context, username string, user models.User) error
	DeleteUserByUsername(ctx context.Context, username string) error
}

type TodoRepoI interface {
	CreateTodo(ctx context.Context, todo models.Todo, user_id uuid.UUID) error
	GetTodos(ctx context.Context, user_id uuid.UUID) (models.GetTodosResp, error)
	GetTodoById(ctx context.Context, todoId string) (models.Todo, error)
	UpdateTodo(ctx context.Context, todo_id string, complete bool) error
	DeleteTodoById(ctx context.Context, todoId string) error
}
