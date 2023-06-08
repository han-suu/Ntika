package item

// import "fmt"

type Service interface {
	FindAll(filter string, sort string) ([]Item, error)

	Create(tagInput ItemInput) (Item, error)
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

	item := Item{
		Name:        itemInput.Name,
		Category:    itemInput.Category,
		Price:       itemInput.Price,
		Description: itemInput.Description,
	}
	newtag, err := s.repository.Create(item)
	return newtag, err
}
