package main

import (
	"context"
	"fmt"

	"tennisy.com/mvp/internal"
)

func main() {
	fmt.Println("tennisy")

	ctx := context.Background()

	internal.New(ctx).Run(ctx)
}
