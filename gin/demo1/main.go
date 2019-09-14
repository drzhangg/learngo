package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.POST("/test", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println("idsL",name)
		c.JSON(200,gin.H{
			"data":name,
		})
	})
	router.Run(":8989")
}
