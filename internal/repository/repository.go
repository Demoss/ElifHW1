package repository

import (
	"ElifHW1/internal/domain"
	"ElifHW1/internal/query"
	"context"
	"gorm.io/gorm"
)

type Sender interface {
	Send(ctx context.Context, dialog domain.Dialog) error
}

type Receiver interface {
	ReceiveLast(ctx context.Context, receive query.Receive) ([]domain.Dialog, error)
}

type Repository struct {
	Sender
	Receiver
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Sender: NewSenderPostgres(db), Receiver: NewReceiverPostgres(db)}
}
