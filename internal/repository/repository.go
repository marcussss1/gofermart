package repository

import (
	"errors"

	"gofermart/internal/models"
)

var (
	ErrUserNotFound               = errors.New("пользователь не найден")
	ErrOrderNotFound              = errors.New("заказ не найден")
	ErrBalanceNotFound            = errors.New("баланс не найден")
	ErrInsufficientFunds          = errors.New("недостаточно средств")
	ErrInvalidOrderNumber         = errors.New("неверный формат номера заказа")
	ErrOrderAlreadyUploaded       = errors.New("заказ уже был загружен этим пользователем")
	ErrOrderUploadedByAnotherUser = errors.New("номер заказа уже был загружен другим пользователем")
)

type UserRepository interface {
	Create(user *models.User) error
	GetByLogin(login string) (*models.User, error)
	GetByID(id string) (*models.User, error)
}

type OrderRepository interface {
	Create(order *models.Order) error
	GetByNumber(number string) (*models.Order, error)
	GetByUserID(userID string) ([]models.Order, error)
	UpdateStatus(orderID string, status string) error
}

type BalanceRepository interface {
	GetBalance(userID string) (*models.Balance, error)
	UpdateBalance(userID string, amount float64) error
	Withdraw(withdrawal *models.Withdrawal) error
	GetWithdrawals(userID string) ([]models.Withdrawal, error)
}
