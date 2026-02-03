package auth

import (
	"context"

	"github.com/nexus-planet/nexus-planet-api/internal/db"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) CreateSession(ctx context.Context, data *db.CreateSessionParams) (*db.AuthSession, error) {
	err := r.q.CreateSession(ctx, *data)
	if err != nil {
		return nil, err
	}

	session, err := r.FindOne(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *Repository) FindOne(ctx context.Context, id string) (*db.AuthSession, error) {
	session, err := r.q.FindOneSession(ctx, id)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) FindOneByEmail(ctx context.Context, email string) (*db.AuthSession, error) {
	session, err := r.q.FindOneSessionByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) FindAllSessions(ctx context.Context) ([]db.AuthSession, error) {
	sessions, err := r.q.FindAllSessions(ctx)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
