package item

import "fmt"

// import "fmt"

type Service interface {
	FindAll(filter string, sort string) ([]Item, error)

	Create(tagInput ItemInput) (Item, error)
	// Pap(tagInput Images2Input) (Images2, error)
	// Pap(tagInput string) (Images2, error)
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
		Size_Chart:  itemInput.Size_Chart,
	}
	newtag, err := s.repository.Create(item)

	if err != nil {
		fmt.Println(err)
	}

	search, err := s.repository.Find(newtag)

	for _, based := range itemInput.Images {
		image := Images2{
			Based:      based,
			Product_ID: search.ID,
		}
		_, _ = s.repository.Pap(image)
	}

	return newtag, err
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

// 	return newtag, err
// }
