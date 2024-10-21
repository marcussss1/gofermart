package user

import (
	"encoding/json"
	"net/http"

	"gofermart/internal/models"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	err := h.service.Register(r.Context(), &user)
	if err != nil {
		// Обработка ошибок
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
