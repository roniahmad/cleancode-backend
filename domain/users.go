package domain

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        int          `json:"id"`
	Username  string       `json:"username" validate:"required"`
	Email     string       `json:"email" validate:"required,email"`
	Password  string       `json:"password" validate:"required"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type UserRepository interface {
	ChangePassword(c context.Context, email string, password string) error
	CreateAccessToken(c context.Context, user *User, conf *Config) (string, time.Time, error)
	CreateUser(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	UserExists(c context.Context, email string) bool
}

type UserUsecase interface {
	ChangePassword(c context.Context, email, oldPasword, newPassword, confirmPassword string) error
	Login(c context.Context, login Login) (map[string]interface{}, error)
	Register(c context.Context, user *User) error
}
