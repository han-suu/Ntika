package item

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(filter string, sort string) ([]Item, error)
	Create(item Item) (Item, error)
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
