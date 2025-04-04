package handlers

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "github.com/go-chi/chi/v5"
)

func getUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Obtener userID de la URL
        userIDStr := chi.URLParam(r, "userID")
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
            return
        }

        var userResp UserResponse // Usa la struct segura para respuestas
        err = db.QueryRow("SELECT id, username FROM users WHERE id = ?", userID).Scan(&userResp.ID, &userResp.Username)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "Usuario no encontrado", http.StatusNotFound)
            } else {
                log.Printf("Error consultando datos de usuario %d: %v", userID, err)
                http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
            }
            return
        }


  w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(userResp)
    }
}


