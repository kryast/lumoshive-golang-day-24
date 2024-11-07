package repository

import (
	"database/sql"
	"fmt"
)

type DashboardRepositoryDB struct {
	DB *sql.DB
}

func NewDashboardRepository(db *sql.DB) DashboardRepositoryDB {
	return DashboardRepositoryDB{DB: db}
}

func (r *DashboardRepositoryDB) GetBookCount() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM books"
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get book count: %v", err)
	}
	return count, nil
}

func (r *DashboardRepositoryDB) GetOrderCount() (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM orders"
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get order count: %v", err)
	}
	return count, nil
}
