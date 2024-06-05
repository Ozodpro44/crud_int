package handlers

import (
	"app/models"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

var ctx = context.Background()

func (h *Handlers) Login() {
	var username string
	var password string

	fmt.Print("Enter Username:")
	fmt.Scanln(&username)
	fmt.Print("Enter Password:")
	fmt.Scanln(&password)

	user, err := h.storage.GetUserRepo().Login(ctx, username)
	if err != nil {
		log.Println("Error user:",err)
		return
	}
	if password==user.Password {
		UserToken.UserId=user.UserID.String()
		UserToken.Username=user.Username
		fmt.Println("You are Loged in")
	}
	fmt.Println("Incorrect password!!!")
}

func (h *Handlers) CreateUser() {
	var user models.User
	user.UserID = uuid.New()
	fmt.Print("Enter Fullname:")
	fmt.Scanln(&user.Fullname)
	fmt.Print("\nEnter user name:")
	fmt.Scanln(&user.Username)
	fmt.Print("Enter password:")
	fmt.Scanln(&user.Password)
	for {
		fmt.Print("\nEnter Gmail:")
		fmt.Scanln(&user.Gmail)
		check := CheckGmail(user.Gmail)
		if check {
			break
		}
		fmt.Println("Invaid email adress (good@exmaple.com)!!!")
	}

	err := h.storage.GetUserRepo().CreateUser(ctx, user)
	if err != nil {
		log.Println("Error on Create user ", err)
		return
	}
	log.Println("User created!!", err)
	fmt.Println("You are registred!!!")

}

func (h *Handlers) GetUsers() {
	users, err := h.storage.GetUserRepo().GetUsers(ctx)

	if err != nil {
		log.Println("Error on get all users:", err)
		return
	}

	fmt.Println("rows |             user_id                    | username |    fullname   |            gmail               |         password      ")
	fmt.Println("-----+----------------------------------------+----------+---------------+--------------------------------+-----------------------")
	for i, v := range users {
		id := v.UserID.String()
		fmt.Printf("%4d | %38s | %8s | %13s | %-30s | %-26s\n", i+1, id, v.Username, v.Fullname, v.Gmail, v.Password)

	}
}

func (h *Handlers) DeleteUser() {
	var username string
	fmt.Print("\nEnter Deleting username:")
	fmt.Scanln(&username)

	err := h.storage.GetUserRepo().DeleteUserByUsername(ctx, username)

	if err != nil {
		log.Println("error on deleting", err)
		return
	}
	fmt.Println("Deleted")
}

func (h *Handlers) UpdateUser() {
	username := UserToken.Username
	var user2 models.User

	fmt.Print("enter New username:")
	fmt.Scanln(&user2.Username)
	fmt.Print("enter New Fullname:")
	fmt.Scanln(&user2.Fullname)
	fmt.Print("enter New Gmail:")
	fmt.Scanln(&user2.Gmail)
	fmt.Print("enter New Password:")
	fmt.Scanln(&user2.Password)

	err := h.storage.GetUserRepo().UpdateUser(ctx, username, user2)
	if err != nil {
		log.Println("error on updating", err)
		return
	}
}
