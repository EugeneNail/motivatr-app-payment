package http

import (
	"github.com/EugeneNail/motivatr-app-payment/internal/application/commands"
	"github.com/EugeneNail/motivatr-app-payment/internal/application/queries"
)

type Handler struct {
	createPaymentHandler *commands.CreatePaymentHandler
	getPaymentHandler    *queries.GetPaymentHandler
}

func NewHandler(
	createPaymentHandler *commands.CreatePaymentHandler,
	getPaymentHandler *queries.GetPaymentHandler,
) *Handler {
	return &Handler{
		createPaymentHandler: createPaymentHandler,
		getPaymentHandler:    getPaymentHandler,
	}
}
