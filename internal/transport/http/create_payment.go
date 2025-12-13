package http

import "net/http"

func (handler *Handler) CreatePayment(request *http.Request) (int, any) {
	return http.StatusOK, "hello"
}
