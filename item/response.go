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
