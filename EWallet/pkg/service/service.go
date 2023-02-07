package service

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/GiairoZeppeli/EWallet/pkg/repository"
)

type Send interface {
	Send(transaction EWallet.Transaction) error
}

type GetLast interface {
	GetLast(count int) ([]EWallet.Transaction, error)
}

type GetBalance interface {
	GetBalance(address string) (float32, error)
}

type Service struct {
	Send
	GetLast
	GetBalance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Send:       NewSendService(repos.Send),
		GetBalance: NewGetBalanceService(repos.GetBalance),
	}
}
