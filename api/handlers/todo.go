package handlers

import (
	"app/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func (h *Handlers) CreateTodo() {
	var todo models.Todo
	task := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Task:")
	task.Scan()
	todo.Task = task.Text()
	todo.UserId = UserToken.UserId
	todo.TodoID = uuid.New()
	todo.CreatedAt = time.Now()
	todo.IsCompleted = false
	err := h.storage.GetTodoRepo().CreateTodo(ctx, todo, UserToken.UserId)

	if err != nil {
		log.Println("Error on creating todo:", err)
		return
	}

	fmt.Println("Task created!!!")
}

func (h *Handlers) GetTodos() {
	todos, err := h.storage.GetTodoRepo().GetTodos(ctx, UserToken.UserId)

	if err != nil {
		log.Println("Error on gett all todos:", err)
		return
	}
	fmt.Println("rows |             user_id                    |                 todo_id                |                Task               | is completed |       create at      ")
	fmt.Println("-----+----------------------------------------+----------------------------------------+-----------------------------------+--------------+-----------------------")
	for i, v := range todos.Todos {
		fmt.Printf("%4d | %38s | %38s | %-33s | %-12v | %-26s\n", i+1, v.UserId, v.TodoID, v.Task, v.IsCompleted, v.CreatedAt)
	}
}

func (h *Handlers) GetTodoById() {
	var todoId string
	fmt.Print("Enter todo ID to get:")
	fmt.Scanln(&todoId)

	todo, err := h.storage.GetTodoRepo().GetTodoById(ctx, todoId)

	if err != nil {
		log.Println("Error on get todo by id:", err)
		return
	}
	fmt.Println("            user_id                    |                 todo_id                |                Task               | is completed |       create at      ")
	fmt.Println("---------------------------------------+----------------------------------------+-----------------------------------+--------------+-----------------------")
	fmt.Printf("%38s | %38s | %-33s | %-12v | %-26s\n", todo.UserId, todo.TodoID, todo.Task, todo.IsCompleted, todo.CreatedAt)

}

func (h *Handlers) UpdateTodo() {
	var todoId string
	var complete string
	var todo_complete bool

	fmt.Print("Enter todo id to update:")
	fmt.Scanln(&todoId)
	fmt.Print("Completed(y/n):")
	fmt.Scanln(&complete)

	if complete == "y" || complete == "Y" {
		todo_complete = true
	} else if complete == "n" || complete == "N" {
		todo_complete = false
	} else {
		fmt.Println("incorrect choice!!!")
		h.UpdateTodo()
	}

	err := h.storage.GetTodoRepo().UpdateTodo(ctx, todoId, todo_complete)

	if err != nil {
		log.Println("Error on updating todo:", err)
		return
	}
	fmt.Println("Todo updated!!!")
}

func (h *Handlers) DeleteTodoById() {
	var todoId string

	fmt.Print("Enter todo ID to delete:")
	fmt.Scanln(&todoId)

	err := h.storage.GetTodoRepo().DeleteTodoById(ctx, todoId)

	if err != nil {
		log.Println("Error on deleting todo:", err)
		return
	}

	fmt.Println("Todo deleted!!!")
}
