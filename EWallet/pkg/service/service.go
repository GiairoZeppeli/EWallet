package service

import "github.com/GiairoZeppeli/EWallet/pkg/repository"

type Send interface {
}

type GetLast interface {
}

type GetBalance interface {
}

type Service struct {
	Send
	GetLast
	GetBalance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
