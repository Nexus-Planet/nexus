package user

import (
	"context"

	"github.com/nexus-planet/nexus-planet-api/internal/db"
)

type Service struct {
	r *Repository
}

func NewService(r *Repository) *Service {
	return &Service{r: r}
}

func (svc *Service) CreateUser(ctx context.Context, data UserCredentials) (*db.User, error) {
	user, err := svc.r.CreateUser(ctx, data.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *Service) FindOneUser(ctx context.Context, id string) (*db.User, error) {
	user, err := svc.r.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *Service) FindAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := svc.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (svc *Service) DeleteUser(ctx context.Context, id string) error {
	err := svc.r.SoftDeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeactivateUser(ctx context.Context, id string) error {
	err := svc.r.DeactivateUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Reactivate(ctx context.Context, id string) error {
	err := svc.r.ReactivateUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
