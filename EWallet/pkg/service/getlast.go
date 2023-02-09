package service

import (
	"github.com/GiairoZeppeli/EWallet"
	"github.com/GiairoZeppeli/EWallet/pkg/repository"
)

type GetLastService struct {
	repo repository.GetLast
}

func NewGetLastService(repo repository.GetLast) *GetLastService {
	return &GetLastService{repo: repo}
}

func (s *GetLastService) GetLast(count int) ([]EWallet.Transaction, error) {
	return s.repo.GetLast(count)
}
