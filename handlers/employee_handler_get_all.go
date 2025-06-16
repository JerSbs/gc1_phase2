package handlers

import (
	"encoding/json"
	"net/http"
	"gc1_phase2/config"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := `SELECT id, name, email FROM employees`
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []map[string]interface{}
	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			http.Error(w, "Error reading employee data", http.StatusInternalServerError)
			return
		}
		employees = append(employees, map[string]interface{}{
			"id":    id,
			"name":  name,
			"email": email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
