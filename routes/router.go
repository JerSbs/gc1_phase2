package routes

import (
	"net/http"
	"p2-graded-challenge-1-JerSbs/handlers"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// ✅ Root route to respond when accessing "/"
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ Employee API is running"))
	})

	// Handle POST and GET /employees
	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateEmployee(w, r)
		} else if r.Method == http.MethodGet {
			handlers.GetAllEmployees(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle GET, PUT, DELETE /employees/:id
	mux.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetEmployeeByID(w, r)
		case http.MethodPut:
			handlers.UpdateEmployee(w, r)
		case http.MethodDelete:
			handlers.DeleteEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
