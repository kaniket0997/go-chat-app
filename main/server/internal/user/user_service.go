package user

import (
	"context"
	"errors"
	"github.com/go-chat-app/main/server/utils"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type service struct {
	repo Repository
	time time.Duration
}

type MyJWTClaims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	secretKey = "secret"
)

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

func (s *service) LoginUser(ctx context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {

	context, cancel := context.WithTimeout(ctx, s.time)
	defer cancel()

	u, err := s.repo.GetUserByEmail(context, request.Email)
	if err != nil {
		return nil, err
	}
	match := utils.CheckPassword(request.Password, u.Password)
	if !match {
		return nil, errors.New("invalid password, can't login")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		Id:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &LoginUserResponse{
		Id:          strconv.Itoa(int(u.ID)),
		Username:    u.Username,
		AccessToken: ss,
	}, nil
}
