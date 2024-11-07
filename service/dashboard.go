package service

import (
	"day-24/repository"
	"fmt"
)

type DashboardService struct {
	dashboardRepo repository.DashboardRepositoryDB
}

func NewDashboardService(repo repository.DashboardRepositoryDB) DashboardService {
	return DashboardService{dashboardRepo: repo}
}

func (bs *DashboardService) GetTotalBooks() (int, error) {
	count, err := bs.dashboardRepo.GetBookCount()
	if err != nil {
		return 0, fmt.Errorf("failed to get book count: %v", err)
	}
	return count, nil
}

func (os *DashboardService) GetTotalOrders() (int, error) {
	count, err := os.dashboardRepo.GetOrderCount()
	if err != nil {
		return 0, fmt.Errorf("failed to get order count: %v", err)
	}
	return count, nil
}
