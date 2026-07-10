package products

import (
	"net/http"
	"production-ecommerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	// 1. call the service -> ListProducts
	// 2. Return JSON in an HTTP response
	products := []string{"hello", "world"}
	json.Write(w, 200, products)
}
