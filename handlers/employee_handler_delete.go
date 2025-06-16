package handlers

import (
	"database/sql"
	"encoding/json"
	"gc1_phase2/config"
	"net/http"
	"strconv"
	"strings"
)

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	// Get employee data before deletion for response
	var emp struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	query := `SELECT id, name, email, phone FROM employees WHERE id = ?`
	err = config.DB.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Phone)
	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to find employee", http.StatusInternalServerError)
		return
	}

	// Proceed to delete
	deleteQuery := `DELETE FROM employees WHERE id = ?`
	_, err = config.DB.Exec(deleteQuery, id)
	if err != nil {
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Employee deleted successfully",
		"employee": emp,
	})
}
