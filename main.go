package main

import "github.com/gin-gonic/gin"

func main() {
	// Crie um router com o Gin
	router := gin.Default()

	// Capture o request GET
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Executamos nosso servidor na porta 8080
	router.Run()
}
