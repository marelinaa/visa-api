// Package psql implements functions for working with postgres database
package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"
	"github.com/marelinaa/visa-api/services/visa/internal/domain"
)

const ctxTime time.Duration = 5 * time.Second

// UserRepo manages the database operations related to users.
type UserRepo struct {
	db *sql.DB
}

// NewRepository creates a new Repository with the given database connection.
func NewRepository(db *sql.DB) *UserRepo {

	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) Create(ctx context.Context, req domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, ctxTime)
	defer cancel()

	queryInsert := ``
	if req.Role == "applicant" {
		queryInsert = `
		INSERT INTO applicant (first_name, last_name, phone_number, email, password_hash, created_at, updated_at)
 		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
 		RETURNING id;`

	} else {
		queryInsert = `
		INSERT INTO operator (first_name, last_name, phone_number, email, password_hash, created_at, updated_at)
 		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
 		RETURNING id;`
	}

	row := repo.db.QueryRowContext(
		ctx,
		queryInsert,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.Email,
		req.PasswordHash,
	)

	if err := row.Scan(&req.ID); err != nil {
		return domain.ErrCreatingUser
	}

	return nil
}

func (repo *UserRepo) SignIn(ctx context.Context, userSignInInput domain.SignInInput) (int64, string, error) {
	ctx, cancel := context.WithTimeout(ctx, ctxTime)
	defer cancel()

	querySelect := `SELECT id, password_hash 
					FROM users 
					WHERE phone_number=$1 AND deleted_at IS NULL;`
	row := repo.db.QueryRowContext(ctx, querySelect, userSignInInput.PhoneNumber)
	var userID int64
	var hash string
	err := row.Scan(&userID, &hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return -1, "", domain.ErrPhoneNumberNotFound
		}

		return -1, "", err
	}

	return userID, hash, nil
}
