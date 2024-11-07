package handler

import (
	"day-24/model"
	"fmt"
	"net/http"
)

func (oh *OrderHandler) OrderListHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := oh.serviceOrder.GetAllOrders()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching orders: %v", err), http.StatusInternalServerError)
		return
	}

	data := struct {
		Orders []model.Order
	}{
		Orders: orders,
	}

	err = templates.ExecuteTemplate(w, "order-list-view", data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %v", err), http.StatusInternalServerError)
		return
	}
}
