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
