package item

import "time"

type Item struct {
	ID          int
	Name        string
	Category    string
	Price       int
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID          int
	Name        string
	Category    string
	Price       int
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}
