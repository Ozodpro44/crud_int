package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TodoRepo struct {
	conn *pgx.Conn
}

func NewTodoRepo(conn *pgx.Conn) repoi.TodoRepoI {
	todoRepo := &TodoRepo{conn: conn}

	return todoRepo
}

func (t *TodoRepo) CreateTodo(ctx context.Context, todo models.Todo, user_id uuid.UUID) error {
	query := `
		INSERT INTO 
		todos (
			user_id,     
			todo_id,     
			task,    
			created_at,  
			is_completed
			) VALUES (
				$1, $2, $3, $4, $5
			)`
	_, err := t.conn.Exec(ctx, query, user_id, todo.TodoID, todo.Task, todo.CreatedAt, todo.IsCompleted)
	if err != nil {
		log.Println("error on Create todo ", err)
		return err
	}
	return nil
}

func (t *TodoRepo) GetTodos(ctx context.Context, user_id uuid.UUID) (models.GetTodosResp, error) {
	var todos models.GetTodosResp
	var todo models.Todo
	query := `
		SELECT * FROM 
			todos
			WHERE user_id = $1;`
	rows, err := t.conn.Query(ctx, query, user_id)
	if err != nil {
		log.Println("Error on get todos", err)
	}
	for rows.Next() {

		err := rows.Scan(
			&todo.UserId, 
			&todo.TodoID, 
			&todo.Task, 
			&todo.CreatedAt, 
			&todo.IsCompleted,
		)

		defer rows.Close()

		if err != nil {
			log.Println("Error on Scan all todos", err)
			return todos, err
		}
		todos.Todos = append(todos.Todos, todo)
	}

	return todos, nil
}

func (t *TodoRepo) GetTodoById(ctx context.Context, todoId string) (models.Todo, error) {
	var todo models.Todo
	todo_id, uerr := uuid.Parse(todoId)

	if uerr != nil {
		log.Println("Incorrect ID!!!", uerr)
		return todo, uerr
	}

	query := `
		SELECT * FROM 
			todos
			WHERE todo_id = $1;`
	err := t.conn.QueryRow(
		ctx, 
		query, 
		todo_id,
		).
		Scan(
			&todo.UserId, 
			&todo.TodoID, 
			&todo.Task, 
			&todo.CreatedAt, 
			&todo.IsCompleted,
		)
	if err != nil {
		log.Println("Error on get todo", err)
		return todo, err
	}

	return todo, nil
}

func (t *TodoRepo) UpdateTodo(ctx context.Context, todo_id string, complete bool) error {
	todoId, terr := uuid.Parse(todo_id)

	if terr != nil {
		log.Println("incorrect type of ID:", terr)
		return terr
	}

	query := `UPDATE todos 
		SET is_completed = $1
		WHERE 
			todo_id = $2;`
	_, err := t.conn.Exec(
		ctx, 
		query, 
		complete, 
		todoId,
	)
	if err != nil {
		log.Println("error on Updating todo: ", err)
		return err
	}
	return nil
}

func (t *TodoRepo) DeleteTodoById(ctx context.Context, todoId string) error {
	todo_id ,terr := uuid.Parse(todoId)

	if terr != nil {
		log.Println("incorrect todo ID:",terr)
		return terr
	}
	
	query := `
		DELETE 
			FROM todos 
			WHERE todo_id=$1;`

	_, err := t.conn.Exec(ctx, query, todo_id)

	if err != nil {
		log.Println("Error on Deleting todo", err)
		return err
	}

	return nil
}
