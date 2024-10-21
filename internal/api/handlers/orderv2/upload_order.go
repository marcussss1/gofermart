package orderv2

import (
	"fmt"
	"net/http"
)

func (h *Handler) UploadOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload order")

	err := h.service.UploadOrder(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//switch err {
		//case repository.ErrInvalidOrderNumber:
		//	http.Error(w, "Неверный формат номера заказа", http.StatusUnprocessableEntity)
		//case repository.ErrOrderAlreadyUploaded:
		//	w.WriteHeader(http.StatusOK)
		//case repository.ErrOrderUploadedByAnotherUser:
		//	http.Error(w, "Номер заказа уже был загружен другим пользователем", http.StatusConflict)
		//default:
		//	http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		//}
		//return
	}

	w.WriteHeader(http.StatusAccepted)
}
