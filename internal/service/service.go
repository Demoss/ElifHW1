package service

import (
	"ElifHW1/internal/command"
	"ElifHW1/internal/domain"
	"ElifHW1/internal/query"
	"ElifHW1/internal/repository"
	"context"
)

type Sender interface {
	Send(ctx context.Context, dialog domain.Dialog) error
	DeleteAllMessages(ctx context.Context, messages command.DeleteAllMessages) error
	UpdateMessage(ctx context.Context, dialog domain.Dialog, message command.UpdateMessage) error
}
type Receiver interface {
	ReceiveLast(ctx context.Context, receive query.Receive) ([]domain.Dialog, error)
}

type Service struct {
	Sender
	Receiver
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Sender: NewSenderService(repo), Receiver: NewReceiverService(repo)}
}
