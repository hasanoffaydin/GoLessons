package main

import "fmt"

func main() {
	var customer1 Customer
	customer1 = Customer{id: 1, name: "Aydin", age: 18, adress: Address{city: "Istanbul", district: "AtaSehir"}}
	address := Address{city: "Azerbaycan", district: "Baku"}
	customer2 := Customer{id: 2, name: "Abbas", age: 19, adress: address}
	customer1.age = 31
	customer2.name = "Arif"
	fmt.Println(customer1)
	fmt.Println(customer2)
	customer1.ToString()
	//changeName(&customer1)
	customer1.changeName("Abbas")
	customer1.ToString()
	customer2.ToString()

}

// func changeName(customer *Customer) {
// 	customer.name = "Murat Coskun"
// }
func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

type Customer struct { // struct = class demekdir
	id     int
	name   string
	age    int
	adress Address
}
type Address struct {
	city     string
	district string
}

func (customer Customer) ToString() {
	fmt.Printf("Name : %s,Age : %d\n", customer.name, customer.age)
}
