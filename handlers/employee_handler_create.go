package handlers

import (
	"encoding/json"
	"fmt"
	"gc1_phase2/config"
	"gc1_phase2/models"
	"net/http"
)

// CreateEmployee handles POST /employees
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid JSON Format", http.StatusBadRequest)
		return
	}

	if emp.Name == "" || emp.Email == "" || emp.Phone == "" {
		http.Error(w, "All fields (name, email, phone) are required", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO employees (name, email, phone) VALUES (?, ?, ?)`
	result, err := config.DB.Exec(query, emp.Name, emp.Email, emp.Phone)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert employee: %v", err), http.StatusBadRequest)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve inserted ID", http.StatusInternalServerError)
		return
	}
	emp.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Employee created successfully",
		"employee": emp,
	})
}
