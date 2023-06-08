package item

type ItemInput struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description" binding:"required"`
}
