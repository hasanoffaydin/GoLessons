package main

import "fmt"

func main() {
	myChannel := make(chan string)
	done := make(chan bool)
	go func() {
		myChannel <- "Hello World"
	}()
	//message, isOpen := <-myChannel
	//fmt.Println(message, isOpen)

	go func() { // bu halda deadlock xetasi bas vermir cunki main funksiyasi bitir ve go islemir
		newMessage := <-myChannel
		fmt.Println(newMessage)
		done <- true
	}()

	<-done
	fmt.Println("End of the main")

}
