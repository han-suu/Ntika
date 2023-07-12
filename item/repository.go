package item

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Find(item Item) (Item, error)
	FindAll(filter string, sort string) ([]Item, error)
	Create(item Item) (Item, error)
	AddToCart(cart CartItem) (CartItem, error)
	FindStock(product_id int, size_id int) (Product_size_stock, error)
	UpdateStock(item_stock Product_size_stock) (Product_size_stock, error)
	GetItemStock(ID int) ([]Product_size_stock, error)
	ItemDetail(ID int) (Item, error)
	GetCart(ID int) ([]CartItem, error)
	FindCartItem(ID int) (CartItem, error)
	DeleteCart(item CartItem) (CartItem, error)
	Order(order Orders) (Orders, error)
	CreateOrderItem(item OrderItem) (OrderItem, error)
	UserHistory(ID int) ([]Orders, error)
	CountSell(ID int) (int64, error)
	NewArr() ([]Item, error)
	// -----------------------------------
	AdminOrder() ([]Orders, error)
	GetOrder(ID int) (Orders, error)
	AdminUpdateOrder(order Orders) (Orders, error)
	// ============================
	Pap(img Images2) (Images2, error)
	AddSize(size Size) (Size, error)
	AddStock(size Product_size_stock) (Product_size_stock, error)
	Thumbnail(ID int) (Images2, error)
	FindImages(ID int) ([]Images2, error)
	GetOrderItem(ID int) ([]OrderItem, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(filter string, sort string) ([]Item, error) {
	var item []Item
	base := r.db.Debug()
	if filter != "" {
		base = base.Where(&Item{Category: filter}) // Adding this condition just if someThing is true
	}
	if sort != "" {
		list := strings.Split(sort, "_")
		val := list[0]
		order := list[1]
		query := fmt.Sprintf("%s %s", val, order)
		base = base.Order(query)
	}
	err := base.Find(&item).Error // Query your results
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return item, err

}

func (r *repository) Create(item Item) (Item, error) {

	err := r.db.Create(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return item, err
}

func (r *repository) AddToCart(cart CartItem) (CartItem, error) {

	err := r.db.Debug().Create(&cart).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return cart, err
}

func (r *repository) Find(item Item) (Item, error) {

	err := r.db.Find(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return item, err
}

func (r *repository) FindStock(product_id int, size_id int) (Product_size_stock, error) {
	var item_stock Product_size_stock

	err := r.db.Where("product_id = ? AND size_id >= ?", product_id, size_id).First(&item_stock).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE FB-EMAIL")
		println("=====================")
	}

	return item_stock, err
}

func (r *repository) UpdateStock(item_stock Product_size_stock) (Product_size_stock, error) {
	err := r.db.Save(&item_stock).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}
	fmt.Printf("error repo %s", err)
	fmt.Println("")
	return item_stock, err
}

func (r *repository) Pap(img Images2) (Images2, error) {

	err := r.db.Create(&img).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return img, err
}

func (r *repository) GetItemStock(ID int) ([]Product_size_stock, error) {
	var items []Product_size_stock
	base := r.db.Debug()
	base = base.Where(&Product_size_stock{Product_ID: ID})
	err := base.Find(&items).Error // Query your results
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return items, err

}

func (r *repository) ItemDetail(ID int) (Item, error) {
	var item Item
	err := r.db.Where(&Item{ID: ID}).First(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return item, err
}

func (r *repository) GetCart(ID int) ([]CartItem, error) {
	var items []CartItem
	err := r.db.Where(&CartItem{User_ID: ID}).Find(&items).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return items, err
}

// Undirect----------------------------------------------
func (r *repository) AddStock(stock Product_size_stock) (Product_size_stock, error) {

	err := r.db.Create(&stock).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return stock, err
}

func (r *repository) Thumbnail(ID int) (Images2, error) {
	var images Images2
	err := r.db.Where(&Images2{Product_ID: ID}).First(&images).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return images, err
}

func (r *repository) FindImages(ID int) ([]Images2, error) {
	var images []Images2
	err := r.db.Where(&Images2{Product_ID: ID}).Find(&images).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return images, err
}

func (r *repository) FindCartItem(ID int) (CartItem, error) {
	var item CartItem
	err := r.db.Where(&CartItem{ID: ID}).First(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return item, err
}

func (r *repository) DeleteCart(item CartItem) (CartItem, error) {

	err := r.db.Delete(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}

	return item, err
}

func (r *repository) Order(order Orders) (Orders, error) {

	err := r.db.Create(&order).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return order, err
}
func (r *repository) CreateOrderItem(item OrderItem) (OrderItem, error) {

	err := r.db.Create(&item).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return item, err
}

func (r *repository) UserHistory(ID int) ([]Orders, error) {
	var orders []Orders
	err := r.db.Where(&Orders{User_ID: ID}).Find(&orders).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return orders, err
}

func (r *repository) GetOrderItem(ID int) ([]OrderItem, error) {
	var items []OrderItem
	err := r.db.Where(&OrderItem{Order_ID: ID}).Find(&items).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return items, err
}

func (r *repository) AdminOrder() ([]Orders, error) {
	var orders []Orders
	base := r.db.Debug()

	err := base.Find(&orders).Error // Query your results
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return orders, err

}

func (r *repository) GetOrder(ID int) (Orders, error) {
	var orders Orders
	err := r.db.Where(&Orders{ID: ID}).First(&orders).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return orders, err
}

func (r *repository) AdminUpdateOrder(order Orders) (Orders, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		println("=====================")
		println("ERROR WHILE Updating")
		println("=====================")
	}
	return order, err
}

func (r *repository) CountSell(ID int) (int64, error) {
	// var item []Item
	base := r.db.Debug()
	var count int64
	// err := base.Find(&item).Error // Query your results
	err := base.Model(&OrderItem{}).Where(&OrderItem{Product_ID: ID}).Count(&count).Error
	fmt.Println(count)
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return count, err

}

func (r *repository) NewArr() ([]Item, error) {
	var items []Item
	base := r.db.Debug()

	err := base.First(&items).Error // Query your results
	if err != nil {
		println("=====================")
		println("ERROR WHILE F")
		println("=====================")
	}

	return items, err

}

// DEV-Only==============================================
func (r *repository) AddSize(size Size) (Size, error) {

	err := r.db.Create(&size).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return size, err
}
