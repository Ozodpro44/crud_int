package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID
	Fullname string
	Username string
	Gmail    string
	Password string
}

type Todo struct {
	TodoID      uuid.UUID
	Task        string
	CreatedAt   time.Time
	IsCompleted bool
}

type GetTodosResp struct {
	Todos []Todo
	Count int
}
