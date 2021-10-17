package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/add-feed", func(c *gin.Context) {
		fmt.Println(c)
		c.JSON(200, gin.H{
			"result": "OK",
		})
	})
}
