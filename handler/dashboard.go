package handler

import (
	"day-24/service"
	"fmt"
	"net/http"
)

type DashboardHandler struct {
	serviceDashboard service.DashboardService
}

func NewDashboardHandler(service service.DashboardService) *DashboardHandler {
	return &DashboardHandler{serviceDashboard: service}
}

func (dh *DashboardHandler) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	totalBooks, err := dh.serviceDashboard.GetTotalBooks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching total books: %v", err), http.StatusInternalServerError)
		return
	}

	totalOrder, err := dh.serviceDashboard.GetTotalOrders()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching total sales: %v", err), http.StatusInternalServerError)
		return
	}

	data := struct {
		TotalBooks int
		TotalOrder int
	}{
		TotalBooks: totalBooks,
		TotalOrder: totalOrder,
	}

	templates.ExecuteTemplate(w, "dashboard-view", data)
}
