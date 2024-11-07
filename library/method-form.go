package library

import "net/http"

func MethodForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cek apakah method POST dan ada field _method dalam form data
		if r.Method == http.MethodPost && r.FormValue("_method") == "PUT" {
			r.Method = http.MethodPut
		} else if r.Method == http.MethodPost && r.FormValue("_method") == "DELETE" {
			r.Method = http.MethodDelete
		}
		next.ServeHTTP(w, r)
	})
}
