package model

import "time"

type Order struct {
	ID           int
	CustomerName string    `json:"customer_name"`
	OrderNumber  string    `json:"order_number"`
	OrderStatus  string    `json:"order_status"`
	OrderDate    time.Time `json:"order_date"`
}
