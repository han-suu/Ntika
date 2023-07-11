package item

type TagResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CatResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

type DetailResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Images      []Images2 `json:"images"`
}

type CartItemResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Qty   int    `json:"qty"`
	Size  string `json:"size"`
}

type OrderItemResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"qty"`
	Size     string `json:"size"`
	Image    string `json:"image"`
}

type HistoryResponse struct {
	ID              int                 `json:"id"`
	Shipping_Method string              `json:"shipping_method"`
	Total           int                 `json:"total"`
	Start           string              `json:"start"`
	End             string              `json:"end"`
	Status          string              `json:"status"`
	Items           []OrderItemResponse `json:"items"`
}
