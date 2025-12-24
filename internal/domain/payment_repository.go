package domain

import "context"

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	List(ctx context.Context, userId int64) ([]*Payment, error)
	Find(ctx context.Context, id int64) (*Payment, error)
	Update(ctx context.Context, payment *Payment) error
	Delete(ctx context.Context, id int64) error
	Count(ctx context.Context, userId int64) (int, error)
}
