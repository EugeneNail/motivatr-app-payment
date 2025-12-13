package domain

import "time"

type Payment struct {
	Id          int
	Date        time.Time
	Description string
	Category    Category
	Value       float32
	UserId      int
}
