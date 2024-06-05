package main

import "fmt"

func main() { //function literal
	//f1()

	func() {
		fmt.Println("New Function")
	}()

	topla := func(x int, y int) int {
		return x + y
	}
	//fmt.Println(z)

	cix := func(a int, b int) int {
		return b - a
	}
	fmt.Println(cix(4, 10))
	fmt.Println(calculator(8, 9, topla, cix))
}

func f1(x int, y int) int {
	return x * y
}
func f2(x int, y int) int {
	return x + y
}
func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
