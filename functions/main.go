package main

import "fmt"

func main() {
	add(10, 20)
	total := add(15, 25)
	fmt.Println(total)
	cixma, vurma := addOther(15, 35)
	fmt.Println(cixma, vurma)
	var numbers = []int{10, 20, 30, 40, 50}
	fmt.Println(sum(numbers))
	fmt.Println(newSum(12, 15, 16))
	fmt.Println(advanceSum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(advanceSum(10, 20, 30))
}
func add(a int, b int) int {
	fmt.Println(a + b)
	return a + b
}
func addOther(a int, b int) (int, int) { // iki ve daha artiq return olanda istifade etmek olar;
	return a - b, a * b
	//return a * b
}
func sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func newSum(x int, y int, z int) int {
	return x + y + z
}
func newSum1(x int, y int, z int, h int) int {
	return x + y + z + h
}
func advanceSum(newnumber ...int) int {
	newTotal := 0
	for _, value := range newnumber {
		newTotal += value
	}
	return newTotal
}
