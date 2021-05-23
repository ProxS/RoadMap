package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var port string = ":8081"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	r := gin.New()
	r.GET("/:message", func(c *gin.Context) {
		message := c.Param("message")
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	r.Run(port) // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
