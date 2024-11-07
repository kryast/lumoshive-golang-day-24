package service

import (
	"day-24/model"
	"day-24/repository"
)

type AdminService struct {
	adminRepo repository.AdminRepositoryDB
}

func NewAdminService(repo repository.AdminRepositoryDB) AdminService {
	return AdminService{adminRepo: repo}
}

func (s *AdminService) CreateAdmin(admin model.Admin) error {
	return s.adminRepo.CreateAdmin(admin)
}
