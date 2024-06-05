package main

import "fmt"

type IShippable interface {
	shippingCost() int
}

type Book struct {
	desi int
}
type Electronic struct {
	desi int
}
type Flower struct {
	desi int
}

func (book Book) shippingCost() int {
	return 5 + book.desi*2
}
func (electronic Electronic) shippingCost() int {
	return 5 + electronic.desi*3
}
func (flowe Flower) shippingCost() int {
	return 5 + flowe.desi*4
}
func main() {
	book1 := Book{desi: 10}
	book2 := Book{desi: 20}

	books := []Book{book1, book2}
	fmt.Println(book1.shippingCost())
	fmt.Println(book2.shippingCost())
	fmt.Println(calculateTotalShippingCostOfBook(books))

	electronic1 := Electronic{desi: 15}
	electronic2 := Electronic{desi: 25}
	electronics := []Electronic{electronic1, electronic2}
	fmt.Println(electronic1.shippingCost())
	fmt.Println(electronic2.shippingCost())
	fmt.Println(calculateTotalShippingCostOfElectronic(electronics))

	var product IShippable
	product = Book{desi: 30}
	fmt.Println(product.shippingCost())
	product = Electronic{desi: 8}
	fmt.Println(product.shippingCost())
	products := []IShippable{Electronic{desi: 15}, Book{desi: 14}, Flower{desi: 18}}
	fmt.Println(calculateTotalShippingCost(products))
}
func calculateTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.shippingCost()
	}
	return total
}
func calculateTotalShippingCostOfBook(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.shippingCost()
	}
	return total
}
func calculateTotalShippingCostOfElectronic(electronic []Electronic) int {
	total := 0
	for _, electronic := range electronic {
		total += electronic.shippingCost()
	}
	return total
}
