package postgres

import (
	"context"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
)

func (repository *PaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	row := repository.db.QueryRow(
		`INSERT INTO payments (date, description, category, value, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		payment.Date, payment.Description, payment.Category, payment.Value, payment.UserId,
	)

	if err := row.Scan(&payment.Id); err != nil {
		return fmt.Errorf("executing an SQL query to write a payment: %w", err)
	}

	return nil
}
