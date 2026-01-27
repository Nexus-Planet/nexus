package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, data Credentials) (*db.User, error) {

	hash, err := HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	user, err := s.repo.CreateUser(ctx, &db.CreateUserParams{ID: id.String(), Email: data.Email, PasswordHash: hash})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, data Credentials) (string, error) {
	user, err := s.repo.FindOneByEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}

	match := CheckHash(data.Password, user.PasswordHash)

	if !match {
		return "", fmt.Errorf("user not found")
	}

	return MakeToken(data.Email), nil
}
