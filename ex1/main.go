package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "api-key", "my-super-secret-api-key")
}

func doSomethingCool(ctx context.Context) {
	// Caso a chave não exista, retornatá <nil>
	apiKey := ctx.Value("api-key")
	fmt.Println(apiKey)
}

func main() {
	fmt.Println("Go context Tutorial")

	ctx := context.Background()
	ctx = enrichContext(ctx)

	doSomethingCool(ctx)
}
