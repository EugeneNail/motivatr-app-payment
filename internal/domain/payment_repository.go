package domain

import "context"

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	List(ctx context.Context, userId int) ([]*Payment, error)
	Find(ctx context.Context, id int) (*Payment, error)
	Update(ctx context.Context, payment *Payment) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context, userId int) (int, error)
}
