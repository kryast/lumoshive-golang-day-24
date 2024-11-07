package repository

import (
	"database/sql"
	"day-24/model"
	"fmt"
)

type AdminRepositoryDB struct {
	DB *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepositoryDB {
	return AdminRepositoryDB{DB: db}
}

func (r *AdminRepositoryDB) CreateAdmin(admin model.Admin) error {
	query := `INSERT INTO admins (name, username, password) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, admin.Name, admin.Username, admin.Password)
	if err != nil {
		return fmt.Errorf("failed to insert admin: %v", err)
	}
	return nil
}
