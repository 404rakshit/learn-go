package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	// "newmongo/articles"
)

func main() {
	// connectToMongo()
	// articles.Main()
	// getWorldTime()
	readContextValue()
}

func openConnection(done chan bool) {
	fmt.Println("Attempting Connection...")

	if rand.Intn(100) > 50 {
		fmt.Println("Oops! Hanging Connection!")
		time.Sleep(100000 * time.Hour)
	} else {
		fmt.Println("Connection Established!")
		time.Sleep(2 * time.Second)
	}

	done <- true
}

func openConnectionWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan bool)
	go openConnection(done)

	select {
	case <-done:
		fmt.Println("Connection Successful")
	case <-ctx.Done():
		fmt.Println("Connection Timeout!")
	}

}
