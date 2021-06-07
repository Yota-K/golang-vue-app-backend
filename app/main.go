package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func db_env() {
	fmt.Println("DB_HOST:", os.Getenv("MYSQL_DATABASE_HOST"))
	fmt.Println("DB_NAME:", os.Getenv("MYSQL_DATABASE"))
	fmt.Println("DB_USER:", os.Getenv("MYSQL_USER"))
	fmt.Println("DB_PASS:", os.Getenv("MYSQL_PASSWORD"))
}

func main() {
	db_env()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
