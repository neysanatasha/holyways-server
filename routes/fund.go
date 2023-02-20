package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func FundRoutes(r *mux.Router) {
	fundRepository := repositories.RepositoryFund(mysql.DB)
	h := handlers.HandlerFund(fundRepository)

	r.HandleFunc("/funds", h.FindFund).Methods("GET")
	r.HandleFunc("/fund/{id}", h.GetFund).Methods("GET")
	// r.HandleFunc("/fundby-status/{id}", h.GetFundStatus).Methods("GET")
	r.HandleFunc("/fund", middleware.Auth(middleware.UploadFile(h.CreateFund))).Methods("POST")
	r.HandleFunc("/fund/{id}", h.UpdateFund).Methods("PATCH")
	r.HandleFunc("/funds", h.DeleteFund).Methods("DELETE")
}
