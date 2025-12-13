package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/EugeneNail/motivatr-app-payment/internal/domain"
)

func (repository *PaymentRepository) Find(id int) (*domain.Payment, error) {
	payment := domain.Payment{}

	row := repository.db.QueryRow(`SELECT id, date, description, category, value, user_id FROM payments WHERE id = $1`, id)
	err := row.Scan(&payment.Id, &payment.Date, &payment.Description, &payment.Category, &payment.Value, &payment.UserId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("executing an SQL query to retrieve payment %d from the db: %w", id, err)
	}

	return &payment, nil
}
