package item

import (
	"fmt"
	"ntika/auth"
)

// import "fmt"

type Service interface {
	FindAll(filter string, sort string) ([]Item, error)

	Create(tagInput ItemInput) (Item, error)
	AddToCart(tagInput CartInput, user auth.User) (CartItem, error)
	// Pap(tagInput Images2Input) (Images2, error)
	// Pap(tagInput string) (Images2, error)
	AddSize(tagInput SizeInput) (Size, error)
	UpdateStock(stockInput StockInput) (Product_size_stock, error)
	// GetAllItemStock()
	GetItemStock(ID int) ([]Product_size_stock, error)
	Thumbnail(ID int) (Images2, error)
	ItemDetail(ID int) (Item, error)
	FindImages(ID int) ([]Images2, error)
	GetCart(user auth.User) ([]CartItem, error)
	DeleteCart(id int) (CartItem, error)
	Order(ID int, orderInput OrderInput) (Orders, error)
	CreateOrderItem(item CartItem, ID int) (OrderItem, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll(filter string, sort string) ([]Item, error) {
	items, err := s.repository.FindAll(filter, sort)
	return items, err

}

func (s *service) Create(itemInput ItemInput) (Item, error) {

	// ADD ITEM
	item := Item{
		Name:        itemInput.Name,
		Category:    itemInput.Category,
		Price:       itemInput.Price,
		Description: itemInput.Description,
	}
	newtag, err := s.repository.Create(item)

	if err != nil {
		fmt.Println(err)
	}
	// ADD IMAGES
	search, err := s.repository.Find(newtag)

	for _, based := range itemInput.Images {
		image := Images2{
			Based:      based,
			Product_ID: search.ID,
		}
		_, _ = s.repository.Pap(image)
	}

	// ADD STOCK
	for i := 1; i < 5; i++ {
		stock := Product_size_stock{
			Size_ID:    i,
			Product_ID: search.ID,
			Stock:      itemInput.Stock[i-1],
		}
		_, _ = s.repository.AddStock(stock)
	}
	return newtag, err
}

func (s *service) AddToCart(itemInput CartInput, user auth.User) (CartItem, error) {

	print("SEROIS")
	// ADD ITEM
	item := CartItem{
		Product_ID: itemInput.Product_ID,
		Quantity:   itemInput.Quantity,
		Size:       itemInput.Size,
		User_ID:    user.ID,
	}
	// print(item)
	newtag, err := s.repository.AddToCart(item)

	if err != nil {
		fmt.Println(err)
	}

	return newtag, err
}

func (s *service) UpdateStock(stockInput StockInput) (Product_size_stock, error) {

	// item_stock := Product_size_stock{
	// 	Size_ID    : stockInput.Size_ID,
	// 	Product_ID : stockInput.Size_ID,
	// 	Stock      int
	// }

	item, err := s.repository.FindStock(stockInput.Product_ID, stockInput.Size_ID)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	item.Stock += stockInput.Stock
	// user.Address = addressInput.Address

	newstock, err := s.repository.UpdateStock(item)
	fmt.Printf("error service %s", err)
	fmt.Println("")
	return newstock, err
	// item := Size{
	// 	Size_Name: tagInput.Size_Name,
	// }
	// newtag, err := s.repository.AddSize(item)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func (s *service) GetItemStock(ID int) ([]Product_size_stock, error) {
	items, err := s.repository.GetItemStock(ID)
	return items, err

}

// func (s *service) Pap(itemInput Images2Input) (Images2, error) {

// 	img := Images2{
// 		Based:      itemInput.Based,
// 		Product_ID: itemInput.Product_ID,
// 	}
// 	newtag, err := s.repository.Pap(img)
// 	return newtag, err
// }

// func (s *service) Pap(itemInput []string) (Images2, error) {

// 	for base,_ := range itemInput {
// 		img := Images2{
// 			Based:      itemInput.Based,
// 			Product_ID: itemInput.Product_ID,
// 		}
// 		newtag, err := s.repository.Pap(img)
// 	}

//		return newtag, err
//	}
func (s *service) AddSize(tagInput SizeInput) (Size, error) {

	item := Size{
		Size_Name: tagInput.Size_Name,
	}
	newtag, err := s.repository.AddSize(item)

	if err != nil {
		fmt.Println(err)
	}

	return newtag, err
}

func (s *service) ItemDetail(ID int) (Item, error) {
	pic, err := s.repository.ItemDetail(ID)
	return pic, err

}

func (s *service) GetCart(user auth.User) ([]CartItem, error) {
	cart, err := s.repository.GetCart(user.ID)
	return cart, err

}

func (s *service) DeleteCart(id int) (CartItem, error) {
	item, err := s.repository.FindCartItem(id)
	if err != nil {
		fmt.Println(err)
	}

	deleteItem, err := s.repository.DeleteCart(item)

	return deleteItem, err
}

func (s *service) Order(ID int, orderInput OrderInput) (Orders, error) {

	// ADD ITEM
	order := Orders{
		User_ID:         ID,
		Sub_Total:       orderInput.Sub_Total,
		Shipping_Method: orderInput.Shipping_Method,
		Shipping_Fee:    orderInput.Shipping_Fee,
		Total_Price:     orderInput.Total_Price,
		Address:         orderInput.Address,
		Status:          "Menunggu Konfirmasi Admin",
		StartDate:       orderInput.StartDate,
		EndDate:         orderInput.EndDate,
		Durasi:          orderInput.Durasi,
	}
	neworder, err := s.repository.Order(order)

	if err != nil {
		fmt.Println(err)
	}
	return neworder, err
}

func (s *service) CreateOrderItem(item CartItem, ID int) (OrderItem, error) {

	// ADD ITEM
	orderitem := OrderItem{
		Order_ID:   ID,
		Product_ID: item.Product_ID,
		Quantity:   item.Quantity,
		Size:       item.Size,
	}
	neworderitem, err := s.repository.CreateOrderItem(orderitem)

	if err != nil {
		fmt.Println(err)
	}

	return neworderitem, err
}

// ============================================================

func (s *service) Thumbnail(ID int) (Images2, error) {
	pic, err := s.repository.Thumbnail(ID)
	return pic, err

}

func (s *service) FindImages(ID int) ([]Images2, error) {
	pics, err := s.repository.FindImages(ID)
	return pics, err

}
