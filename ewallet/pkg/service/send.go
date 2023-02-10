package service

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/GiairoZeppeli/EWallet/pkg/repository"
)

type SendService struct {
	repo repository.Send
}

func NewSendService(repo repository.Send) *SendService {
	return &SendService{repo: repo}
}

func (s *SendService) Send(transaction ewallet.Transaction) error {
	return s.repo.Send(transaction)
}
