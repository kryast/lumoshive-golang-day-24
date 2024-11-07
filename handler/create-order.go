package handler

import (
	"day-24/library"
	"day-24/model"
	"day-24/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OrderHandler struct {
	serviceOrder service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{serviceOrder: service}
}

func (oh *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	if order.OrderDate.IsZero() {
		parsedDate, err := time.Parse("2006-01-02", order.OrderDate.String())
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid date format: %v", err), http.StatusBadRequest)
			return
		}
		order.OrderDate = parsedDate
	}

	err = oh.serviceOrder.CreateOrder(order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating order: %v", err), http.StatusInternalServerError)
		return
	}

	library.SuccessResponse(w, "success", order)
}
