package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/JerryDdie77/Hezze/Server/internal/domain"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func (p *PostgresUserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := "SELECT id, user_name, first_name, surname, password_hash, email, phone, is_blocked FROM users WHERE id = $1"

	var u domain.User

	if err := p.db.QueryRowContext(ctx, query, id).Scan(
		&u.ID,
		&u.UserName,
		&u.FirstName,
		&u.Surname,
		&u.PasswordHash,
		&u.Email,
		&u.Phone,
		&u.IsBlocked,
	); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return &domain.User{}, domain.ErrNotFound
		}

		return &domain.User{}, fmt.Errorf("Scan: %w", err)
	}

	return &u, nil
}
