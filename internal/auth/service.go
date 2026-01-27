package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
	"github.com/nexus-planet/nexus-planet-api/internal/user"
)

type Service struct {
	auth *Repository
	user *user.Repository
}

func NewService(auth *Repository, user *user.Repository) *Service {
	return &Service{auth: auth, user: user}
}

func (s *Service) CreateUser(ctx context.Context, data Credentials) (*db.User, error) {

	hash, err := HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	user, err := s.user.CreateUser(ctx, &db.CreateUserParams{ID: id.String(), Email: data.Email, PasswordHash: hash})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, data Credentials) (string, error) {
	user, err := s.user.FindOneByEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}

	match := CheckHash(data.Password, user.PasswordHash)

	if !match {
		return "", fmt.Errorf("user not found")
	}

	return MakeToken(data.Email), nil
}
