package user

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
}
