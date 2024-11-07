package library

import "net/http"

func MethodForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.FormValue("_method") == "PUT" {
			r.Method = http.MethodPut
		} else if r.Method == http.MethodPost && r.FormValue("_method") == "DELETE" {
			r.Method = http.MethodDelete
		}
		next.ServeHTTP(w, r)
	})
}
