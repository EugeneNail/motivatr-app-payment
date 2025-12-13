package domain

type PaymentRepository interface {
	Create(payment *Payment) error
	List(userId int) ([]*Payment, error)
	Find(id int) (*Payment, error)
	Update(payment *Payment) error
	Delete(id int) error
	Count(userId int) (int, error)
}
