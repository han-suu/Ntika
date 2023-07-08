package item

type ItemInput struct {
	Name        string   `json:"name" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	Stock       []int    `json:"stock" binding:"required"`
}

type Images2Input struct {
	Based      string `json:"based" binding:"required"`
	Product_ID int    `json:"product_id" binding:"required"`
}

type CartInput struct {
	Product_ID int    `json:"product_id" binding:"required"`
	Quantity   int    `json:"qty" binding:"required"`
	Size       string `json:"size" binding:"required"`
}

type StockInput struct {
	Product_ID int `json:"product_id" binding:"required"`
	Size_ID    int `json:"size_id" binding:"required"`
	Stock      int `json:"stock" binding:"required"`
}

// ================================================================
type SizeInput struct {
	Size_Name string `json:"size_name" binding:"required"`
}
