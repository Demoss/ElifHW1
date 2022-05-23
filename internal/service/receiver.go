package service

import (
	"ElifHW1/internal/domain"
	"ElifHW1/internal/query"
	"ElifHW1/internal/repository"
	"context"
)

type ReceiverService struct {
	repo repository.Receiver
}

func NewReceiverService(repo repository.Receiver) *ReceiverService {
	return &ReceiverService{repo: repo}
}

func (s *ReceiverService) ReceiveLast(ctx context.Context, receive query.Receive) ([]domain.Dialog, error) {
	return s.repo.ReceiveLast(ctx, receive)
}
