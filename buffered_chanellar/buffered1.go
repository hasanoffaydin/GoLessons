package main

// func main() {
// 	myChannel := make(chan int, 2)

// 	myChannel <- 1
// 	myChannel <- 2
// 	//myChannel <- 3
// 	fmt.Println("End of the main")

// 	fmt.Println(<-myChannel)

// 	myChannel <- 3

// 	fmt.Println(<-myChannel)
// 	fmt.Println(<-myChannel)
// 	//fmt.Println(<-myChannel)

// 	//data := <-myChannel
// 	//fmt.Println(data)
// 	//data1 := <-myChannel
// 	//fmt.Println(data1)

// 	messages := make(chan string, 2)

// 	go func() {
// 		data1 := <-messages
// 		fmt.Println("First : ", data1)

// 		data2 := <-messages
// 		fmt.Println("Second : ", data2)
// 	}()

// 	messages <- "Hello"
// 	messages <- "World"
// 	messages <- "World2"

// 	time.Sleep(1 * time.Second)

// 	fmt.Println("End of the main")
// }
