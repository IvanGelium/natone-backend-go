package postgress

import (
	"context"

	"github.com/IvanGelium/main-service/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	createUserQuery = `
		INSERT INTO users (id, email, password_hash, account_id, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at;
	`
	getUserByEmailQuery = `
		SELECT id, email, password_hash, account_id, role, created_at, updated_at
		FROM users
		WHERE email = $1;
	`
)

type authRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) domain.AuthRepository {
	return &authRepository{pool: pool}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var u domain.User
	err := r.pool.QueryRow(ctx, getUserByEmailQuery, email).
		Scan(
			&u.ID,
			&u.Email,
			&u.PasswordHash,
			&u.AccountId,
			&u.Role,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt)
	if err != nil {
		return domain.User{}, nil
	}
	return domain.User{}, nil
}

func (r *authRepository) CreateUser(ctx context.Context, u domain.User, id uuid.UUID) (domain.User, error) {

	err := r.pool.QueryRow(ctx, createUserQuery, id, u.Email, u.PasswordHash, u.AccountId, u.Role).
		Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}
