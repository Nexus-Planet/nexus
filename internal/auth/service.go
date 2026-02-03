package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
)

type Service struct {
	r *Repository
}

func NewService(r *Repository) *Service {
	return &Service{r: r}
}

func (svc *Service) CreateSession(ctx context.Context, data Credentials) (*db.AuthSession, error) {

	hash, err := HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	session, err := svc.r.CreateSession(ctx, &db.CreateSessionParams{ID: id.String(), Email: data.Email, PasswordHash: hash})
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (svc *Service) Login(ctx context.Context, data Credentials) (string, error) {
	session, err := svc.r.FindOneByEmail(ctx, data.Email)
	if err != nil {
		return "", err
	}

	match := CheckHash(data.Password, session.PasswordHash)

	if !match {
		return "", fmt.Errorf("user not found")
	}

	return MakeToken(data.Email), nil
}
