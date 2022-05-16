package repository

import (
	"ElifHW1/internal/domain"
	"context"
	"gorm.io/gorm"
)

type SenderPostgres struct {
	db *gorm.DB
}

func NewSenderPostgres(db *gorm.DB) *SenderPostgres {
	return &SenderPostgres{db: db}
}

func (r *SenderPostgres) Send(ctx context.Context, dialog domain.Dialog) error {
	if res := r.db.WithContext(ctx).Create(&dialog); res.Error != nil {
		return res.Error
	}
	return nil
}
