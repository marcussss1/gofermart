package api

import (
	"github.com/gorilla/mux"
	"gofermart/internal/api/handlers/orderv2"
)

func NewRouter(
	//userHandler user.Handler,
	//orderHandler order.Handler,
	//balanceHandler balance.Handler,
	orderV2Handler *orderv2.Handler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/upload", orderV2Handler.UploadOrder).Methods("POST")

	//r.HandleFunc("/api/user/register", userHandler.Register).Methods("POST")
	//r.HandleFunc("/api/user/login", userHandler.Login).Methods("POST")
	//
	//r.HandleFunc("/api/user/orders", orderHandler.UploadOrder).Methods("POST")
	//r.HandleFunc("/api/user/orders", orderHandler.GetOrders).Methods("GET")
	//
	//r.HandleFunc("/api/user/balance", balanceHandler.GetBalance).Methods("GET")
	//r.HandleFunc("/api/user/balance/withdraw", balanceHandler.Withdraw).Methods("POST")
	//r.HandleFunc("/api/user/withdrawals", balanceHandler.GetWithdrawals).Methods("GET")

	return r
}
