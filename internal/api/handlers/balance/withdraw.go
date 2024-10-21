package balance

import (
	"encoding/json"
	"net/http"

	"gofermart/internal/models"
	"gofermart/internal/repository"
)

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var withdrawal models.Withdrawal
	if err := json.NewDecoder(r.Body).Decode(&withdrawal); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(string)

	err := h.service.Withdraw(r.Context(), userID, withdrawal.Order, withdrawal.Sum)
	if err != nil {
		switch err {
		case repository.ErrInvalidOrderNumber:
			http.Error(w, "Неверный номер заказа", http.StatusUnprocessableEntity)
		case repository.ErrInsufficientFunds:
			http.Error(w, "На счету недостаточно средств", http.StatusPaymentRequired)
		default:
			http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
