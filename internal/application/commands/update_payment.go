package commands

import (
	"context"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/application"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
	"time"
)

type UpdatePaymentCommand struct {
	// Used in application logic
	PaymentId int64
	UserId    int64

	// Actual payload
	Date        time.Time
	Description string
	Category    domain.Category
	Value       float32
}

type UpdatePaymentHandler struct {
	repository domain.PaymentRepository
}

func NewUpdatePaymentHandler(repository domain.PaymentRepository) *UpdatePaymentHandler {
	return &UpdatePaymentHandler{
		repository: repository,
	}
}

func (handler *UpdatePaymentHandler) Handle(ctx context.Context, command UpdatePaymentCommand) error {
	payment, err := handler.repository.Find(ctx, command.PaymentId)
	if err != nil {
		return fmt.Errorf("retrieving payment %d from the database: %w", command.PaymentId, err)
	}

	if payment == nil {
		return application.ErrNotFound
	}

	if payment.UserId != command.UserId {
		return application.ErrNotFound
	}

	payment.Date = command.Date.Truncate(time.Hour * 24)
	payment.Description = command.Description
	payment.Category = command.Category
	payment.Value = command.Value

	if err := handler.repository.Update(ctx, payment); err != nil {
		return fmt.Errorf("writing changes of payment %d to the database: %w", payment.Id, err)
	}

	return nil
}
