package main

// func main() {
// 	myChannel := make(chan int)

// 	go func() {
// 		for i := 1; i <= 10; i++ {
// 			myChannel <- i
// 			fmt.Println("Send data : ", i)
// 			time.Sleep(1 * time.Second)
// 		}

// 		close(myChannel) // kanali baglayir
// 	}()

// 	for {
// 		data, isOpen := <-myChannel
// 		if isOpen == false {
// 			break
// 		}
// 		fmt.Println("Received data : ", data)
// 	}
// }
