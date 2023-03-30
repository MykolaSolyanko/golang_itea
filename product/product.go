package product

import (
	"fmt"
)

type Product struct {
	Name     string
	Quantity int
	Price    float64
}

type Storage struct {
	Items map[string]Product
}

func (s Storage) ViewProducts() {
	for _, v := range s.Items {
		fmt.Println("Name - ", v.Name, " price - ", v.Price)
	}
}

func (s Storage) ViewProductDetails(productName string) *Product {
	product, ok := s.Items[productName]

	if !ok {
		return nil
	}

	fmt.Println("Name - ", product.Name, " price - ", product.Price, " available - ", product.Quantity)

	return &product
}
