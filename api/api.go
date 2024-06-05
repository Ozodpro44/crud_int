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
			1.CreateUser
			2.Get ALL
			3.Delete User
			0.Exit
			`)
			fmt.Scanln(&x)
	
			switch x {
			case 1:
				h.CreateUser()
			case 2:
				h.GetUsers()
			case 3:
				h.DeleteUser()
			// case 4:
			// 	h.GetUserByUsername(UserRepo)
			// case 5:
			// 	h.GetAllUSers(UserRepo)
			// case 0:
				// return
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