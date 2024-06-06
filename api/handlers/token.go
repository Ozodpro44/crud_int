package handlers

import "github.com/google/uuid"

type Token struct {
	Username	string
	UserId 		uuid.UUID
}

var UserToken *Token