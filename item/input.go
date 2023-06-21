package item

type ItemInput struct {
	Name        string   `json:"name" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	Size_Chart  string   `json:"size_chart" binding:"required"`
}

type Images2Input struct {
	Based      string `json:"based" binding:"required"`
	Product_ID int    `json:"product_id" binding:"required"`
}
