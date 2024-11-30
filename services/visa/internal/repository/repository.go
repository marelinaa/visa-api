// Package cart implements functions for working with database
package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/marelinaa/visa-api/services/visa/internal/domain"
)

// Repository manages the database operations related to carts.
type Applicant struct {
	db *sql.DB
}

// NewRepository creates a new Repository with the given database connection.
func NewRepository(db *sql.DB) *Applicant {
	return &Applicant{
		db: db,
	}
}

func (repo *Applicant) AddApplication(ctx context.Context, application domain.Application) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return nil
}
