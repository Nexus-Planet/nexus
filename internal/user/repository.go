package user

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

func (r *Repository) CreateUser(ctx context.Context, id string) (*db.User, error) {
	err := r.q.CreateUser(ctx, id)
	if err != nil {
		return nil, err
	}

	user, err := r.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) FindOne(ctx context.Context, id string) (*db.User, error) {
	user, err := r.q.FindOneUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) FindAll(ctx context.Context) ([]db.User, error) {
	users, err := r.q.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) SoftDeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *Repository) DeactivateUser(ctx context.Context, id string) error {
	return nil
}
func (r *Repository) ReactivateUser(ctx context.Context, id string) error {
	return nil
}

func (r *Repository) SetUsername(ctx context.Context, params *db.SetUserNameParams) (*db.User, error) {
	err := r.q.SetUserName(ctx, *params)
	if err != nil {
		return nil, err
	}

	user, err := r.FindOne(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
