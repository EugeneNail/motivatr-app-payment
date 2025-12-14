package postgres

import (
	"context"
	"fmt"
)

func (repository *PaymentRepository) Delete(ctx context.Context, id int) error {
	if _, err := repository.db.Exec(`DELETE FROM payments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("executing an SQL query to delete payment with id %d: %w", id, err)
	}

	return nil
}
