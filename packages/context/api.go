package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func getWorldTime() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println(ctx.Deadline())

	url := "https://timeapi.io/api/time/current/zone?timeZone=Asia%2FKolkata" //"https://worldtimeapi.ortg/api/timezone/America/Toronto"
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		fmt.Println(ctx.Err())
		return
	}

	// fmt.Println("Response body:", resp.Body)
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
