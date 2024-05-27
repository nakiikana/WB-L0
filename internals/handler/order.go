package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (h *Handler) OrderInfo(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		logrus.Errorf("no id found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order, err := h.s.OrderInfo(uuid)
	if err != nil {
		logrus.Errorf("error while retrieving the order: %v", err)
		return
	}
	orderInfo, err := json.Marshal(order)
	if err != nil {
		logrus.Errorf("error while marshalling order")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(orderInfo); err != nil {
		logrus.Errorf("error when writing orderInfo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
