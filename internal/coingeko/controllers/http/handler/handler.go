package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

var timeout = time.Second * 3

type handler struct {
	svc Service
}

func New(svc Service) *handler {
	return &handler{
		svc: svc,
	}
}
func (h *handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/price", h.Price)
}
func (h *handler) Price(w http.ResponseWriter, r *http.Request) {
	coin := r.FormValue("coin")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	coinResp, err := h.svc.GetPrice(ctx, coin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	rBytes, err := json.Marshal(coinResp.Current_price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(rBytes)
}
