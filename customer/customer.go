package customer

import (
	"fmt"
	p "main/product"
)

type Customer struct {
	Name  string
	Email string
	Cart  map[string]int
}

type Repository interface {
	GetCustomer(customerEmail string)
}

type Storage struct {
	Repository
	Items map[string]Customer
}

func (c Storage) GetCustomer(customerEmail string) *Customer {
	customer, ok := c.Items[customerEmail]

	if !ok {
		return nil
	}

	return &customer
}

func (c Storage) ViewCart(customerEmail string) {
	customer := c.GetCustomer(customerEmail)

	fmt.Println("=User ", customer.Email, " cart:")

	for name, value := range customer.Cart {
		fmt.Println("Name - ", name, " quantity - ", value)
	}
}

func (c Storage) AddToCart(customerEmail string, productName string, quantity int) {
	customer := c.GetCustomer(customerEmail)
	customer.Cart[productName] = quantity
}

func (c *Storage) Checkout(customerEmail string, productStorage *p.Storage) {
	fmt.Println("Checkout is completed")

	customer := c.GetCustomer(customerEmail)
	products := productStorage.Items

	for key, value := range customer.Cart {
		product := products[key]
		product.Quantity = product.Quantity - value
		products[key] = product

		fmt.Println("Product - ", key, ", quantity - ", value)

		delete(customer.Cart, key)
	}
}
