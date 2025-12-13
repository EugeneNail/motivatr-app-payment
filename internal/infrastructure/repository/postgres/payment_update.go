package postgres

import (
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
)

func (repository *PaymentRepository) Update(payment *domain.Payment) error {
	_, err := repository.db.Exec(`
		UPDATE payments
		SET date = $1, description = $2, category = $3, value = $4, user_id = $5
		WHERE id = $6
	`, payment.Date, payment.Description, payment.Category, payment.Value, payment.UserId, payment.Id)

	if err != nil {
		return fmt.Errorf("executing an SQL query to update payment %d: %w", payment.Id, err)
	}

	return nil
}
