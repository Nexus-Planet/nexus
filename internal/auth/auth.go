package auth

import "time"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSession struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	IsActive  string `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthSessionDB struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	IsActive     string    `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type CreateSessionParams struct {
	ID           string `db:"id"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

func (as *AuthSessionDB) ToAuthSession() *AuthSession {
	return &AuthSession{
		ID:        as.ID,
		UserID:    as.UserID,
		Email:     as.Email,
		IsActive:  as.IsActive,
		CreatedAt: as.CreatedAt.Format(time.RFC3339),
		UpdatedAt: as.UpdatedAt.Format(time.RFC3339),
	}
}
