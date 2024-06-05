package main

import "fmt"

func main() {
	defer fmt.Println("Hello") // funksiya bitdikden sonra onunde yazildigi setiri isledir
	fmt.Println("World")

	defer fmt.Println("1 ci defer") // tersine bas verir proses deferler cox olanda asagidan yuxariya dogru gedir
	defer fmt.Println("2 ci defer")
	defer fmt.Println("3 ci defer")
	fmt.Println("Main Funksiyasi")

	// var condition = true
	// fmt.Println("1-ci defer")
	// if condition {
	// 	return
	// }
	// defer fmt.Println("2-ci defer")

	x := 10
	y := 20
	defer fmt.Println("x deyeri : ", x)
	defer fmt.Println("y deyeri : ", y)
	x = 100
	y = 200
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)

	defer cleaneUp()
	var condition = true
	if condition {
		panic("An error occured...") // panic olanda run islemi dayanir
	}
}

func cleaneUp() {
	fmt.Println("Cleane Up Worked....")
}
