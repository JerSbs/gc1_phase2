package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"p2-graded-challenge-1-JerSbs/config"
	"strconv"
	"strings"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL: /employees/3
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID in URL", http.StatusBadRequest)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var emp struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if emp.Name == "" || emp.Email == "" || emp.Phone == "" {
		http.Error(w, "All fields (name, email, phone) are required", http.StatusBadRequest)
		return
	}

	query := `UPDATE employees SET name = ?, email = ?, phone = ? WHERE id = ?`
	result, err := config.DB.Exec(query, emp.Name, emp.Email, emp.Phone, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update employee: %v", err), http.StatusBadRequest)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Employee updated successfully",
		"employee": map[string]interface{}{
			"id":    id,
			"name":  emp.Name,
			"email": emp.Email,
			"phone": emp.Phone,
		},
	})
}
