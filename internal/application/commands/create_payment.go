package commands

import "github.com/EugeneNail/motivatr-app-payment/internal/domain"

type CreatePaymentCommand struct {
}

type CreatePaymentResult struct {
}

type CreatePaymentHandler struct {
	repository domain.PaymentRepository
}

func NewCreatePaymentHandler(repository domain.PaymentRepository) *CreatePaymentHandler {
	return &CreatePaymentHandler{
		repository: repository,
	}
}
