package repository

import (
	"day-24/model"
)

func (or *OrderRepositoryDB) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order

	query := "SELECT id, customer_name, order_number, order_status, order_date FROM orders"
	rows, err := or.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.ID, &order.CustomerName, &order.OrderNumber, &order.OrderStatus, &order.OrderDate)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
