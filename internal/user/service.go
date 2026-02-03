package user

import (
	"context"
)

type Service struct {
	r *Repository
}

func NewService(r *Repository) *Service {
	return &Service{r: r}
}

func (svc *Service) CreateUser(ctx context.Context, data UserCredentials) {

}

func (svc *Service) FindOneUser(ctx context.Context) {

}

func (svc *Service) FindAllUsers(ctx context.Context) {

}

func (svc *Service) DeleteUser(ctx context.Context) {

}

func (svc *Service) DeactivateUser(ctx context.Context) {

}
