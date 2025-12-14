package postgres

import (
	"context"
	"fmt"
)

func (repository *PaymentRepository) Count(ctx context.Context, userId int) (int, error) {
	row := repository.db.QueryRow(`SELECT COUNT(id) FROM payments WHERE user_id = $1`, userId)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("executing an SQL query to count payments of user %d: %w", userId, err)
	}

	return count, nil
}
