package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()
	
	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(200,
			gin.H{
				"message": "pong",
			})
	})

	ginRouter.Run(":8080") 
}