package main

import "fmt"

func main() {
	var a int
	var p *int
	a = 10

	p = &a // adress tutan deyisken

	fmt.Println(a)
	fmt.Println(p)  // adresi print edecek
	fmt.Println(*p) // adresde olan deyeri aliriq

	*p = 20
	fmt.Println(*p)
	fmt.Println(a)

	var d = 10
	var b int
	var c *int
	b = d
	c = &d
	*c = 20
	fmt.Println(d, b)

	var h int = 10
	var g *int
	g = &h
	fmt.Println(h)
	newAdd12(h)
	fmt.Println(h)
	add12(g)
	fmt.Println(h)

	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)
}
func changeValue(numbers []int) {
	numbers[0] = 100
}
func add12(x *int) { // pass by reference
	*x = *x + 12
}
func newAdd12(x int) { // pass by value
	x = x + 12
}
