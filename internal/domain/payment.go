package domain

import "time"

type Payment struct {
	Id          int64
	Date        time.Time
	Description string
	Category    Category
	Value       float32
	UserId      int64
}
