package handler

import (
	"tools/internals/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	s *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{s: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/order-info", h.OrderInfo).Methods("GET")
	r.HandleFunc("/new-order", h.NewOrder).Methods("POST")
	return r
}