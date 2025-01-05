package usecase

import (
	"context"
	"time"

	"wetees.com/domain"
	"wetees.com/internal/crypt"
	"wetees.com/internal/vars"
)

type userUsecase struct {
	repo           domain.UserRepository
	conf           *domain.Config
	contextTimeout time.Duration
}

// Change password
func (uc *userUsecase) ChangePassword(c context.Context, email string, oldPasword string, newPassword string, confirmPassword string) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return vars.ErrUserWithEmailNotFound
	}

	if match := crypt.CheckPasswordHash(oldPasword, user.Password); !match {
		return vars.ErrInvalidCredentials
	}

	if newPassword != confirmPassword {
		return vars.ErrPasswordNotMatch
	}

	hash, _ := crypt.HashPassword(newPassword)
	err = uc.repo.ChangePassword(c, email, hash)

	return err
}

// Login
func (uc *userUsecase) Login(c context.Context, login domain.Login) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(c, time.Duration(uc.contextTimeout)*time.Second)
	defer cancel()

	var (
		result map[string]interface{}
	)

	user, err := uc.repo.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return result, vars.ErrLoginFailed
	}

	if match := crypt.CheckPasswordHash(login.Password, user.Password); !match {
		return result, vars.ErrLoginFailed
	}

	accessToken, _, err := uc.repo.CreateAccessToken(ctx, &user, uc.conf)
	if err != nil {
		return result, vars.ErrLoginFailed
	}

	result = map[string]interface{}{
		"access_token": accessToken,
	}

	return result, err
}

// Register user
func (uc *userUsecase) Register(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	//1. Check user exists
	if exists := uc.repo.UserExists(ctx, user.Email); exists {
		return vars.ErrUserAlreadyExists
	}

	//2. Create user entry
	user.Password, _ = crypt.HashPassword(user.Password)
	if err := uc.repo.CreateUser(c, user); err != nil {
		return vars.ErrRegisterUserFailed
	}

	return nil
}

func NewUserUsecase(repo domain.UserRepository, conf *domain.Config, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		repo:           repo,
		conf:           conf,
		contextTimeout: timeout,
	}
}
