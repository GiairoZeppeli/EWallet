package service

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/GiairoZeppeli/EWallet/pkg/repository"
)

type Send interface {
	Send(transaction ewallet.Transaction) error
}

type GetLast interface {
	GetLast(count int) ([]ewallet.Transaction, error)
}

type GetBalance interface {
	GetBalance(address string) (float32, error)
}
type InitDb interface {
	InitDb() error
}

type Service struct {
	Send
	GetLast
	GetBalance
	InitDb
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Send:       NewSendService(repos.Send),
		GetBalance: NewGetBalanceService(repos.GetBalance),
		InitDb:     NewInitDbService(repos.InitDb),
		GetLast:    NewGetLastService(repos.GetLast),
	}
}
