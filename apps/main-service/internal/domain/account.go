package domain

import (
	"context"

	"github.com/google/uuid"
)

type Account struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type UserJWT struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	AccountId uuid.UUID `json:"account_id"`
}

type AuthService interface {
	Login(ctx context.Context, data LoginInput) (UserJWT, error)
	SignUp(ctx context.Context, data SignupInput) (UserJWT, error)
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	CreateUser(ctx context.Context, user User, id uuid.UUID) (User, error)
}
