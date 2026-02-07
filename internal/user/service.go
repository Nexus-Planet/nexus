package user

import (
	"context"
	"time"

	"github.com/nexus-planet/nexus-planet-api/internal/db"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) CreateUser(ctx context.Context, data UserCredentials) (*User, error) {
	user, err := svc.repo.CreateUser(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

func (svc *Service) FindOneUser(ctx context.Context, id string) (*User, error) {
	user, err := svc.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

func (svc *Service) FindAllUsers(ctx context.Context) ([]*User, error) {
	userDBs, err := svc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*User, len(userDBs))
	for i, user := range userDBs {
		if user != nil {
			users[i] = user.ToUser()
		}
	}

	return users, nil
}

func (svc *Service) SetUsername(ctx context.Context, data *SetUsernameCredentials) (*User, error) {
	user, err := svc.repo.SetUsername(ctx, &SetUsernameParams{ID: data.ID, Username: data.Username, UsernameChangedAt: time.Now()})
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

func (svc *Service) UpdateUserData(ctx context.Context, data *UpdateUserCredentials) (*User, error) {

	user, err := svc.repo.UpdateUserData(ctx, &UpdateUserParams{ID: *data.ID, DisplayName: *db.ToNullString(data.DisplayName)})
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

func (svc *Service) SoftDeleteUser(ctx context.Context, id string) error {
	err := svc.repo.SoftDeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeactivateUser(ctx context.Context, id string) error {
	err := svc.repo.DeactivateUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) Reactivate(ctx context.Context, id string) error {
	err := svc.repo.ReactivateUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
