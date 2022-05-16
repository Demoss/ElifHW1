package repository

import (
	"ElifHW1/internal/domain"
	"ElifHW1/internal/query"
	"context"
	"gorm.io/gorm"
)

type ReceiverPostgres struct {
	db *gorm.DB
}

func NewReceiverPostgres(db *gorm.DB) *ReceiverPostgres {
	return &ReceiverPostgres{db: db}
}

func (r *ReceiverPostgres) ReceiveLast(ctx context.Context, receive query.Receive) ([]domain.Dialog, error) {
	var dialog []domain.Dialog

	err := r.db.WithContext(ctx).Table("dialogs").Where("send_from = ? and send_to = ?", receive.From, receive.To).Find(&dialog).Error
	if err != nil {
		return nil, err
	}
	return dialog, nil
}
