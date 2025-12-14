package http

import (
	"encoding/json"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/application/commands"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
	"github.com/EugeneNail/motivatr-lib-common/pkg/authentication"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation/rules"
	"net/http"
	"time"
)

type CreatePaymentInput struct {
	Date        string  `json:"date"`
	Category    int     `json:"category"`
	Description string  `json:"description"`
	Value       float32 `json:"value"`
}

func (handler *Handler) CreatePayment(request *http.Request) (int, any) {
	userId, err := authentication.ExtractHttpUserId(request.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("extracting user id from context: %w", err)
	}

	input := CreatePaymentInput{}
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("parsing the input: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"date":     input.Date,
		"category": input.Category,
		"value":    input.Value,
	}, map[string][]rules.RuleFunc{
		"date":     {rules.Required(), rules.Date()},
		"category": {rules.Required()},
		"value":    {rules.Required()},
	})

	if err := validator.Validate(); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("validating the input: %w", err)
	}

	if validator.Failed() {
		return http.StatusUnprocessableEntity, validator.Errors()
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		validator.AddError("date", "The date field format is invalid")
		return http.StatusUnprocessableEntity, validator.Errors()
	}

	command := commands.CreatePaymentCommand{
		Date:        date,
		Category:    domain.Category(input.Category),
		Description: input.Description,
		Value:       input.Value,
		UserId:      userId,
	}

	result, err := handler.createPaymentHandler.Handle(request.Context(), command)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("handling a createPayment command: %w", err)
	}

	if len(result.ValidationErrors) > 0 {
		return http.StatusUnprocessableEntity, result.ValidationErrors
	}

	return http.StatusOK, result.Id
}
