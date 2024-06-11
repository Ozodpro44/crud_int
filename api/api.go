package api

import (
	"app/api/handlers"
	"app/storage"
	"fmt"
)


func Api(storage storage.StorageI) {
	h := handlers.NewHandlers(storage)
	
	
	x := -1
	for x != 0 {
		if handlers.UserToken!= nil {
			fmt.Println(`
               User                    Todo
        ------------------------+------------------------	
        |   1.CreateUser        |   6.Create Todo       |
        |   2.Get ALL           |   7.Get all todos     |
        |   3.Get User          |   8.Get todo by id    |
        |   4.Update User       |   9.Update todo       |
        |   5.Delete User       |   10.Delete todo      |
        ------------------------+------------------------
            11.Log out              0.Exit
			`)
			fmt.Scanln(&x)
	
			switch x {
			case 1:
				h.CreateUser()
			case 2:
				h.GetUsers()
			case 3:
				h.GetUser()
			case 4:
				h.UpdateUser()
			case 5:
				h.DeleteUser()
			case 6:
				h.CreateTodo()
			case 7:
				h.GetTodos()
			case 8:
				h.GetTodoById()
			case 9:
				h.UpdateTodo()
			case 10:
				h.DeleteTodoById()
			case 11:
				h.Logout()
			case 0:
				return
			}
		}else {
			fmt.Println(`
			1.Login
			2.Registration`)
			fmt.Scanln(&x)

			switch x {
			case 1:
				h.Login()
			case 2:
				h.CreateUser()
			}
		}
	}
}