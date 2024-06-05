package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go f1() // islemesine go qerar verir
	// go f2()
	// fmt.Println("End of the main")
	// time.Sleep(1 + time.Second)

	// wg.Add(3) // olarsa deadlocker hatasi verir
	// wg.Add(1) // olarsa sadece 1 dene go prosesini gozleyir
	wg := sync.WaitGroup{} // wait group go + funksiyalarin bitmesini gozleyir
	wg.Add(5)              // go proseslerin sayini bildirir

	go func() {
		defer wg.Done() // go prosesi bitdikden sonra prosesi sayir
		fmt.Println("F1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("F2")
	}()

	startTime := time.Now()

	go func() {
		fmt.Println("f1")
		time.Sleep(3 * time.Second)
		defer wg.Done()

	}()
	go func() {
		fmt.Println("f2")
		time.Sleep(3 * time.Second)
		defer wg.Done()

	}()
	go func() {
		fmt.Println("f3")
		time.Sleep(3 * time.Second)
		defer wg.Done()

	}()

	wg.Wait() // olmasa go proseslerinin bitmesini gozlemir program biten kimi dayanir

	fmt.Println("Passed time : ", time.Now().Sub(startTime))
	//fmt.Println("End of the main")

}
func f1() {
	fmt.Println("F1")
	time.Sleep(3 + time.Second) //prosesin yatmasi
}
func f2() {
	fmt.Println("F2")
}
