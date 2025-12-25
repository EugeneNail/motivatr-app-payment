package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/application"
	"github.com/EugeneNail/motivatr-app-payment/internal/application/commands"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
	"github.com/EugeneNail/motivatr-lib-common/pkg/authentication"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation"
	"github.com/EugeneNail/motivatr-lib-common/pkg/validation/rules"
	"net/http"
	"strconv"
	"time"
)

type UpdatePaymentInput struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Category    int     `json:"category"`
	Value       float32 `json:"value"`
}

func (handler *Handler) UpdatePayment(request *http.Request) (int, any) {
	userId, err := authentication.ExtractHttpUserId(request.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("extracting user id from a request: %w", err)
	}

	id, err := strconv.ParseInt(request.PathValue("id"), 10, 64)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("parsing the 'id' request parameter: %w", err)
	}

	input := UpdatePaymentInput{}
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("parsing the input: %w", err)
	}

	validator := validation.NewValidator(map[string]any{
		"date":        input.Date,
		"description": input.Description,
		"category":    input.Category,
		"value":       input.Value,
	}, map[string][]rules.RuleFunc{
		"date":        {rules.Required(), rules.Min(10), rules.Max(10), rules.Date()},
		"description": {rules.Max(50)},
		"category":    {rules.Required(), rules.Min(1), rules.Max(10)},
		"value":       {rules.Required(), rules.Min(-1000000), rules.Max(1000000)},
	})

	if err := validator.Validate(); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("validating the input: %w", err)
	}

	if validator.Failed() {
		return http.StatusUnprocessableEntity, validator.Errors()
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("parsing a date '%s': %w", input.Date, err)
	}

	command := commands.UpdatePaymentCommand{
		PaymentId:   id,
		UserId:      userId,
		Date:        date,
		Description: input.Description,
		Category:    domain.Category(input.Category),
		Value:       input.Value,
	}

	err = handler.updatePaymentHandler.Handle(request.Context(), command)
	if err != nil && errors.Is(err, application.ErrNotFound) {
		return http.StatusNotFound, nil
	}

	if err != nil && errors.Is(err, application.ErrPermissionDenied) {
		return http.StatusNotFound, nil
	}
	
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("handling an UpdatePayment command: %w", err)
	}

	return http.StatusNoContent, nil
}
