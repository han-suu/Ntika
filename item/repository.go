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
	// ============================
	Pap(img Images2) (Images2, error)
	AddSize(size Size) (Size, error)
	AddStock(size Product_size_stock) (Product_size_stock, error)
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

func (r *repository) Pap(img Images2) (Images2, error) {

	err := r.db.Create(&img).Error

	if err != nil {
		println("=====================")
		println("ERROR WHILE CREATING")
		println("=====================")
	}

	return img, err
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
