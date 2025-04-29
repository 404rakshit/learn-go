package main

import (
	"context"
	"fmt"
)

func readContextValue() {
	ctx := context.WithValue(context.Background(), "username", "alex")
	ctx = context.WithValue(ctx, "password", 123)

	result := ctx.Value("username")
	password := ctx.Value("password")

	if result == nil {
		fmt.Println("User Not Found!")
		return
	}

	var username string = result.(string)

	fmt.Printf("\nHello %s!", username)
	fmt.Printf("\nPassword %d!", password)
}
