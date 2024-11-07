package repository

import (
	"database/sql"
	"day-24/model"
	"fmt"
)

func (r *AdminRepositoryDB) GetAdminByUsername(username string) (model.Admin, error) {
	var admin model.Admin
	query := `SELECT id, name, username, password FROM admins WHERE username = $1`
	err := r.DB.QueryRow(query, username).Scan(&admin.ID, &admin.Name, &admin.Username, &admin.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return admin, nil // Return an empty admin if not found
		}
		return admin, fmt.Errorf("failed to query admin: %v", err)
	}
	return admin, nil
}
