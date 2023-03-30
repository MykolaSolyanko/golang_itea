package main

import (
	"bufio"
	"errors"
	"fmt"
	c "main/customer"
	p "main/product"
	"os"
	"strconv"
	"strings"
)

const defaultInputAgreement = "Y"

func getProducts() map[string]p.Product {
	return map[string]p.Product{
		"wine": {
			Name:     "wine",
			Quantity: 15,
			Price:    23.99,
		},
		"beer": {
			Name:     "beer",
			Quantity: 23,
			Price:    1.55,
		},
		"cider": {
			Name:     "cider",
			Quantity: 7,
			Price:    1.99,
		},
		"gin": {
			Name:     "gin",
			Quantity: 9,
			Price:    69.99,
		},
		"soda": {
			Name:     "soda",
			Quantity: 29,
			Price:    0.79,
		},
	}
}

func scanTextWithMessage(scanner *bufio.Scanner, message string) string {
	fmt.Println(message)

	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}

	return input
}

func addCustomer(input string, customerStorage *c.Storage) (*c.Customer, error) {
	inputValues := strings.Split(input, ",")

	if (len(inputValues)) != 2 {
		return nil, errors.New("invalid input, please try again")
	}

	email := inputValues[0]
	name := inputValues[1]

	customerStorage.Items = map[string]c.Customer{
		email: {
			Name:  name,
			Email: email,
			Cart:  make(map[string]int),
		},
	}

	return customerStorage.GetCustomer(email), nil
}

func main() {
	productStorage := &p.Storage{}
	customerStorage := &c.Storage{}
	scanner := bufio.NewScanner(os.Stdin)

	productStorage.Items = getProducts()

	fmt.Println("Hello! Welcome to our BV shop.")
	input := scanTextWithMessage(scanner, "Type your email and name (comma-separated) to start shopping")

	customer, err := addCustomer(input, customerStorage)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		fmt.Println("Product list: ")
		productStorage.ViewProducts()

		productName := scanTextWithMessage(scanner, "Type product name to view the details")
		product := productStorage.ViewProductDetails(productName)

		if product == nil {
			fmt.Println("This product is not exists!")
			continue
		}

		input = scanTextWithMessage(scanner, "Do you want to add it to cart ? Y/n")

		if input == defaultInputAgreement {
			input = scanTextWithMessage(scanner, "Type quantity")
			quantity, _ := strconv.Atoi(input)

			if product.Quantity < quantity+customer.Cart[productName] {
				fmt.Println("This is a big quantity, try another")
				continue
			}

			customerStorage.AddToCart(customer.Email, productName, quantity)
		}

		customerStorage.ViewCart(customer.Email)

		input = scanTextWithMessage(scanner, "Do you want to checkout? Y/n")

		if input == defaultInputAgreement {
			customerStorage.Checkout(customer.Email, productStorage)
		}
	}
}
