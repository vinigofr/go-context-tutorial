package main

import (
	"context"
	"fmt"
	"time"
)

// Comportamento usando deadlines e context

func main() {
	fmt.Println("Go Context Tutorial")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doSomethinCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("OH NO!!! I've exceeded the deadline")
	}
	time.Sleep(2 * time.Second)
}

func doSomethinCool(ctx context.Context) {
	// Simulador de longo processo rodando.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("Doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
