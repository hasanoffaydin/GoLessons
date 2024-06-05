package main

import "fmt"

var z = 10

func main() {
	var condition = true
	if condition {
		//block scope
		var x = 10
		fmt.Println(x)
	}
	//fmt.Println(x)

	var condition1 = true //function scope
	if condition1 {
		fmt.Println(condition1)
	}
	fmt.Println(condition1)

	test() // function scope
	//fmt.Println(x,y)

	fmt.Println(z) // global scope

}
func test() {
	var x = 10
	var y = 20
	fmt.Println(x, y)
	fmt.Println(z)
}
