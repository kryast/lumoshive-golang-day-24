package repository

import (
	"database/sql"
	"day-24/model"
	"fmt"
)

type OrderRepositoryDB struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryDB {
	return OrderRepositoryDB{DB: db}
}

func (r *OrderRepositoryDB) CreateOrder(order model.Order) error {
	query := `INSERT INTO orders (customer_name, order_number, order_status, order_date) 
		VALUES ($1, $2, $3, $4)`

	_, err := r.DB.Exec(query, order.CustomerName, order.OrderNumber, order.OrderStatus, order.OrderDate)
	if err != nil {
		return fmt.Errorf("failed to insert sale: %v", err)
	}
	return nil
}
