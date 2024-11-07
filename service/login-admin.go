package service

import (
	"day-24/model"
	"fmt"
)

func (s *AdminService) ValidateAdmin(username, password string) (model.Admin, error) {
	admin, err := s.adminRepo.GetAdminByUsername(username)
	if err != nil {
		return model.Admin{}, fmt.Errorf("failed to get admin: %v", err)
	}

	return admin, nil
}
