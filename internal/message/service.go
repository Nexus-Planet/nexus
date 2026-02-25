package message

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) SendMessage(ctx context.Context, data Message) (*Message, error) {
	id := uuid.New()

	msg, err := svc.repo.CreateMessage(ctx, CreateMessageParams{ID: id.String(), Content: data.Content, Type: data.Type, Attachments: data.Attachments})
	if err != nil {
		return nil, err
	}
	return msg.ToMessage(), err
}

func (svc *Service) FindOne(ctx context.Context, id string) (*Message, error) {
	msg, err := svc.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return msg.ToMessage(), nil
}

func (svc *Service) FindAll(ctx context.Context) ([]*Message, error) {
	msgsDB, err := svc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	msgs := make([]*Message, len(msgsDB))
	for i, msg := range msgsDB {
		if msg != nil {
			msgs[i] = msg.ToMessage()
		}
	}

	return msgs, nil
}

func (svc *Service) UpdateData(ctx context.Context, data UpdateMessage) (*Message, error) {
	msg, err := svc.repo.UpdateData(ctx, &UpdateMessageParams{ID: data.ID, Content: data.Content})
	if err != nil {
		return nil, err
	}

	return msg.ToMessage(), nil
}

func (svc *Service) SoftDelete(ctx context.Context, id string) error {
	err := svc.repo.SoftDelete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
