package repoi

import (
	"app/models"
	"context"
)

type UserRepoI interface {
	Login(ctx context.Context, username string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, username string, user models.User) error
	DeleteUserByUsername(ctx context.Context, username string) error
}

type TodoRepoI interface {
	CreateTodo(ctx context.Context, todo models.Todo) error
	GetTodos(ctx context.Context, limit, page string) (*models.GetTodosResp, error)
	GetTodoById(ctx context.Context, todoId string) (*models.Todo, error)
	UpdateTodo(ctx context.Context, todo models.Todo) error
	DeleteTodoById(ctx context.Context, todoId string) error
}
