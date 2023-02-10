package service

import (
	"github.com/GiairoZeppeli/EWallet/pkg/repository"
)

type GetBalanceService struct {
	repo repository.GetBalance
}

func NewGetBalanceService(repo repository.GetBalance) *GetBalanceService {
	return &GetBalanceService{repo: repo}
}

func (s *GetBalanceService) GetBalance(address string) (float32, error) {
	return s.repo.GetBalance(address)
}
