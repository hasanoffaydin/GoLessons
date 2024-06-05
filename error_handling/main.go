package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var x int
	var y float32
	var pointer *int
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(pointer)

	var newPointer *int
	if newPointer == nil {
		fmt.Println("New Pointer hec bir deyer saxlamir")
	}

	FileData, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(FileData)
	}

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	var productService = ProductService{}
	err = productService.Add(Product{id: 1, name: "", price: 3000})
	if err != nil {
		fmt.Println(err)
	}
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Bolen 0 ola bilmez")
	}
	return x / y, nil
}

type Product struct {
	id    int
	name  string
	price int
}
type ProductService struct {
}

func (productService ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Productun adi bos ola bilmez", code: "1001"}
	}
	if product.price < 0 {
		return ValidationError{text: "Price 0 dan boyuk olmalidir", code: "2002"}
	}
	fmt.Println("Urun Eklendi")
	return nil
}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Xeta : %s, Code : %s", validationError.text, validationError.code)
}
