package order

import (
	"gofermart/internal/models"
)

func (s *Service) GetUserOrders(userID int) ([]models.Order, error) {
	return s.repository.GetByUserID(userID)
}
