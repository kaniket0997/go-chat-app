package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewUserRepository(db DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {

	var lastInsertId int64
	query := "INSERT INTO users(username, password, email) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.ID, user.Username, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}
	user.ID = lastInsertId

	return user, nil
}