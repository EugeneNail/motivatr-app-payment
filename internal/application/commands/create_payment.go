package commands

import (
	"context"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation/rules"
	"time"
)

type CreatePaymentCommand struct {
	UserId      int64
	Date        time.Time
	Category    domain.Category
	Description string
	Value       float32
}

type CreatePaymentResult struct {
	Id               int64
	ValidationErrors map[string]string
}

type CreatePaymentHandler struct {
	repository domain.PaymentRepository
}

func NewCreatePaymentHandler(repository domain.PaymentRepository) *CreatePaymentHandler {
	return &CreatePaymentHandler{
		repository: repository,
	}
}

func (handler *CreatePaymentHandler) Handle(ctx context.Context, command CreatePaymentCommand) (*CreatePaymentResult, error) {
	result := &CreatePaymentResult{}

	validator := validation.NewValidator(map[string]any{
		"description": command.Description,
		"category":    command.Category,
		"value":       command.Value,
	}, map[string][]rules.RuleFunc{
		"description": {rules.Max(50)},
		"category":    {rules.Min(1), rules.Max(10)},
		"value":       {rules.Min(-1000000), rules.Max(1000000)},
	})

	if err := validator.Validate(); err != nil {
		return nil, fmt.Errorf("validating the command: %w", err)
	}

	if validator.Failed() {
		result.ValidationErrors = validator.Errors()
		return result, nil
	}

	payment := domain.Payment{
		Date:        command.Date,
		Description: command.Description,
		Value:       command.Value,
		Category:    command.Category,
		UserId:      command.UserId,
	}

	if err := handler.repository.Create(ctx, &payment); err != nil {
		return nil, fmt.Errorf("writing a payment to the DB: %w", err)
	}

	result.Id = payment.Id
	return result, nil
}
