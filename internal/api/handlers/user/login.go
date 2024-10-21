package user

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	user, err := h.service.Login(r.Context(), loginData.Login, loginData.Password)
	if err != nil {
		http.Error(w, "Неверная пара логин/пароль", http.StatusUnauthorized)
		return
	}

	// Здесь должна быть логика создания сессии или JWT токена

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
