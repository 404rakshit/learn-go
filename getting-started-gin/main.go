package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping", handleGet)
	router.GET("/count", handleCount)
	router.Run() // listen and serve on 0.0.0.0:8080
}

func handleGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleCount(c *gin.Context) {
	var n int = 0
	for i := 0; i < 100000000; i++ {
		n += 1
	}
	c.JSON(200, gin.H{"count": n})
}
