package http

import (
	"errors"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/application"
	"github.com/EugeneNail/motivatr-app-payment/internal/application/queries"
	"github.com/EugeneNail/motivatr-app-payment/internal/transport/http/dto"
	"github.com/EugeneNail/motivatr-lib-common/pkg/authentication"
	"net/http"
	"strconv"
)

func (handler *Handler) GetPayment(request *http.Request) (int, any) {
	userId, err := authentication.ExtractHttpUserId(request.Context())
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("extracting user id from a request: %w", err)
	}

	paymentId, err := strconv.ParseInt(request.PathValue("id"), 10, 64)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("parsing the \"id\" request parameter: %w", err)
	}

	query := queries.GetPaymentQuery{PaymentId: paymentId, UserId: userId}
	result, err := handler.getPaymentHandler.Handle(request.Context(), query)
	if err != nil && errors.Is(err, application.ErrPermissionDenied) {
		return http.StatusNotFound, nil
	}

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("handling a GetPayment query for id %d: %w", paymentId, err)
	}

	if result.Payment == nil {
		return http.StatusNotFound, nil
	}

	payment := dto.PaymentWithoutUserId{
		Id:          result.Payment.Id,
		Date:        result.Payment.Date.Format("2006-01-02"),
		Description: result.Payment.Description,
		Category:    int(result.Payment.Category),
		Value:       result.Payment.Value,
	}

	return http.StatusOK, payment
}
