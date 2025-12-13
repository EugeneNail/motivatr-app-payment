package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
)

func (repository *PaymentRepository) List(userId int) ([]*domain.Payment, error) {
	payments := make([]*domain.Payment, 0)
	rows, err := repository.db.Query(`
		SELECT id, date, description, category, value, user_id 
		FROM payments
		WHERE user_id = $1
		ORDER BY date DESC
	`, userId)
	defer rows.Close()

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return payments, nil
	}

	if err != nil {
		return nil, fmt.Errorf("executing an SQL query to retrieve payments: %w", err)
	}

	for rows.Next() {
		var payment domain.Payment
		if err := rows.Scan(&payment.Id, &payment.Date, &payment.Description, &payment.Category, &payment.Value, &payment.UserId); err != nil {
			return nil, fmt.Errorf("scanning a row into a payment: %w", err)
		}
		payments = append(payments, &payment)
	}

	return payments, nil
}
