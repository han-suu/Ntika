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
	ID         int
	User_ID    int
	Product_ID int
	Quantity   int
	Size       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Orders struct {
	ID              int
	User_ID         int
	Sub_Total       int
	Shipping_Method string
	Shipping_Fee    int
	Total_Price     int
	Address         string
	City            string
	Status          string
	StartDate       time.Time
	EndDate         time.Time
	Durasi          int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ID         int
	Order_ID   int
	Product_ID int
	Quantity   int
	Size       string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Images struct {
	ID         int
	Path       int
	Product_ID int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Images2 struct {
	ID         int
	Based      string
	Product_ID int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Size struct {
	ID        int
	Size_Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product_size_stock struct {
	ID         int
	Size_ID    int
	Product_ID int
	Stock      int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Ongkir_tb struct {
	ID     int
	Ongkir int

	CreatedAt time.Time
	UpdatedAt time.Time
}
