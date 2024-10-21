package order

import (
	"encoding/json"
	"net/http"

	"gofermart/internal/repository"
)

func (h *Handler) UploadOrder(w http.ResponseWriter, r *http.Request) {
	var orderNumber string
	if err := json.NewDecoder(r.Body).Decode(&orderNumber); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(string)

	err := h.service.UploadOrder(r.Context(), userID, orderNumber)
	if err != nil {
		switch err {
		case repository.ErrInvalidOrderNumber:
			http.Error(w, "Неверный формат номера заказа", http.StatusUnprocessableEntity)
		case repository.ErrOrderAlreadyUploaded:
			w.WriteHeader(http.StatusOK)
		case repository.ErrOrderUploadedByAnotherUser:
			http.Error(w, "Номер заказа уже был загружен другим пользователем", http.StatusConflict)
		default:
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
