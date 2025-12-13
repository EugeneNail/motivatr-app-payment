package http

import "github.com/EugeneNail/motivatr-app-payment/internal/application/commands"

type Handler struct {
	createPaymentHandler *commands.CreatePaymentHandler
}

func NewHandler(createPaymentHandler *commands.CreatePaymentHandler) *Handler {
	return &Handler{
		createPaymentHandler: createPaymentHandler,
	}
}
