package main

import (
	"context"
	"fmt"

	"tennisly.com/mvp/internal"
)

func main() {
	fmt.Println("tennisly")

	ctx := context.Background()

	internal.New(ctx).Run(ctx)
}
