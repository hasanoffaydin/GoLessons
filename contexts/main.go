package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	id   int64
	name string
}

var productChanel = make(chan Product)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	go getProductDetailsFromExternalService(10)
	select {
	case productDetail := <-productChanel:
		fmt.Println("Product getirildi ", productDetail)
	case <-ctx.Done():
		fmt.Println("Timeout occurred, context cancelled")
	}
	fmt.Println("End of the main")

	newctx := context.Background()
	newctx = context.WithValue(newctx, "correlation-id", "abc123")
	F1(newctx)

}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(10 * time.Second)
	productChanel <- Product{
		id:   productId,
		name: "Paltar yuyan",
	}
}

func F1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	F2(ctx)
}

func F2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	F3(ctx)
}

func F3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))

}
