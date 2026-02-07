package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) CreateSession(ctx context.Context, data Credentials) (*AuthSession, error) {

	hash, err := HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	session, err := svc.repo.CreateSession(ctx, &CreateSessionParams{ID: id.String(), Email: data.Email, PasswordHash: hash})
	if err != nil {
		return nil, err
	}

	return session.ToAuthSession(), nil
}

func (svc *Service) Login(ctx context.Context, data Credentials) (string, error) {
	session, err := svc.repo.FindOneByEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}

	match := CheckHash(data.Password, session.PasswordHash)

	if !match {
		return "", fmt.Errorf("user not found")
	}

	return MakeToken(data.Email), nil
}
