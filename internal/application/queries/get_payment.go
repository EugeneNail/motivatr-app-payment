package queries

import (
	"context"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/application"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
)

type GetPaymentQuery struct {
	UserId    int64
	PaymentId int64
}

type GetPaymentResult struct {
	Payment *domain.Payment
}

type GetPaymentHandler struct {
	repository domain.PaymentRepository
}

func NewGetPaymentHandler(repository domain.PaymentRepository) *GetPaymentHandler {
	return &GetPaymentHandler{
		repository: repository,
	}
}

func (handler *GetPaymentHandler) Handle(ctx context.Context, query GetPaymentQuery) (*GetPaymentResult, error) {
	result := &GetPaymentResult{}

	payment, err := handler.repository.Find(ctx, query.PaymentId)
	if err != nil {
		return nil, fmt.Errorf("retrieving payment %d from the database: %w", query.PaymentId, err)
	}

	if payment == nil {
		return result, nil
	}

	if payment.UserId != query.UserId {
		return nil, application.ErrPermissionDenied
	}

	result.Payment = payment
	return result, nil
}
