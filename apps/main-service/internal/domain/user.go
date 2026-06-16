package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `json:"id"`
	AccountId    uuid.UUID  `json:"account_id"`
	Email        string     `json:"email"`
	Role         Role       `json:"role"`
	PasswordHash string     `json:"password"`
	Nickname     *string    `json:"nickname"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)
