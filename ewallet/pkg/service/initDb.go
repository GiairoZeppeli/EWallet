package service

import "github.com/GiairoZeppeli/EWallet/pkg/repository"

type InitDbService struct {
	repo repository.InitDb
}

func NewInitDbService(repo repository.InitDb) *InitDbService {
	return &InitDbService{repo: repo}
}

func (s *InitDbService) InitDb() error {
	return s.repo.InitDb()
}
