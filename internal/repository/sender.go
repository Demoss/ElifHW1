package repository

import (
	"ElifHW1/internal/command"
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

func (r *SenderPostgres) DeleteAllMessages(ctx context.Context, messages command.DeleteAllMessages) error {
	if res := r.db.WithContext(ctx).Where("send_from = ? and send_to = ?", messages.From, messages.To).Delete(&domain.Dialog{}).Error; res != nil {
		return res
	}
	return nil
}

func (r *SenderPostgres) UpdateMessage(ctx context.Context, dialog domain.Dialog, message command.UpdateMessage) error {
	if res := r.db.WithContext(ctx).Table("dialogs").Where("send_from = ? and send_to = ? and text = ?", dialog.SendFrom, dialog.SendTo, dialog.Text).Update("text", message.Text).Error; res != nil {
		return res
	}

	return nil
}
