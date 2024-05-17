package service

import (
	"tools/internals/repository"
)

type Service struct {
	rp *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{rp: repository}
}
