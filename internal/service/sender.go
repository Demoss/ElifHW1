package service

import (
	"ElifHW1/internal/command"
	"ElifHW1/internal/domain"
	"ElifHW1/internal/repository"
	"context"
)

type SenderService struct {
	repo repository.Sender
}

func NewSenderService(repo repository.Sender) *SenderService {
	return &SenderService{repo: repo}
}
func (s *SenderService) Send(ctx context.Context, dialog domain.Dialog) error {
	return s.repo.Send(ctx, dialog)
}

func (s *SenderService) DeleteAllMessages(ctx context.Context, messages command.DeleteAllMessages) error {
	return s.repo.DeleteAllMessages(ctx, messages)
}
func (s *SenderService) UpdateMessage(ctx context.Context, dialog domain.Dialog, message command.UpdateMessage) error {
	return s.repo.UpdateMessage(ctx, dialog, message)
}
