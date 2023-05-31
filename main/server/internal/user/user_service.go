package user

import (
	"context"
	"github.com/go-chat-app/main/server/utils"
	"strconv"
	"time"
)

type service struct {
	repo Repository
	time time.Duration
}

func NewService(repo Repository, time time.Duration) Service {
	return &service{
		repo: repo,
		time: time,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, s.time)
	defer cancel()

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: password,
	}
	r, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUserResponse{
		Id:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil

}
