package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"wetees.com/domain"
	"wetees.com/internal/jwt"
	"wetees.com/internal/vars"
)

const (
	TABLE_USERS = "users"
)

type userRepository struct {
	Conn *sql.DB
}

// Change user password
func (ur *userRepository) ChangePassword(c context.Context, email string, password string) (err error) {
	var (
		stmt   *sql.Stmt
		result sql.Result
	)

	query := fmt.Sprintf("UPDATE %s set password=? WHERE email = ?", TABLE_USERS)
	if stmt, err = ur.Conn.PrepareContext(c, query); err != nil {
		return
	}

	if result, err = stmt.ExecContext(c, password, email); err != nil {
		return
	}

	_, err = result.RowsAffected()

	return
}

// Create access token
func (u *userRepository) CreateAccessToken(c context.Context, user *domain.User, conf *domain.Config) (string, time.Time, error) {
	return jwt.CreateToken(user.Email, strconv.FormatInt(int64(user.ID), 10), conf.AccessTokenSecret, conf.AccessTokenExpiry, conf.SecretKey, conf.AppIssuer)
}

// Create new user
func (uc *userRepository) CreateUser(c context.Context, user *domain.User) (err error) {
	var (
		stmt   *sql.Stmt
		result sql.Result
	)

	query := fmt.Sprintf(`INSERT INTO %s (username, email, password, created_at)	
		VALUES(?, ?, ?, ?)`, TABLE_USERS)

	if stmt, err = uc.Conn.PrepareContext(c, query); err != nil {
		return err
	}

	if result, err = stmt.ExecContext(c, &user.Username, &user.Email, &user.Password, time.Now().Local().UTC()); err != nil {
		return err
	}

	if _, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

// Get user by email
func (ur *userRepository) GetUserByEmail(c context.Context, email string) (user domain.User, err error) {
	query := fmt.Sprintf("SELECT id, username, email, password, created_at FROM %s WHERE email=?", TABLE_USERS)

	err = ur.Conn.QueryRowContext(c, query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)

	switch {
	case err == sql.ErrNoRows:
		user = domain.User{}
		err = vars.ErrUserWithEmailNotFound
	case err != nil:
		user = domain.User{}
	default:
		err = nil
	}

	return
}

// Check if user already exists
func (ur *userRepository) UserExists(c context.Context, email string) bool {
	var count int
	if err := ur.Conn.QueryRowContext(c, fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE email=?", TABLE_USERS), email).
		Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return count > 1
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		Conn: db,
	}
}
