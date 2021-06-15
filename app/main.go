package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Yota-K/golang-vue-app-backend/database"
	"github.com/Yota-K/golang-vue-app-backend/log"
)

func db_env() {
	database.Database_init()
}

func main() {
	db_env()

	r := gin.Default()
	r.Use(Logger())

	// 標準的なリクエストの例
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// パスパラメータの例
	r.GET("/user/:id", func(c *gin.Context) {
		id:= c.Param("id")
		c.JSON(200, gin.H{
			"user_id": id,
		})
	})

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path, method := c.Request.URL.Path, c.Request.Method
		log.InfoLog(log.Fields{"path": path, "method": method}, "APIの実行始めます。")
		c.Next()
		log.InfoLog(log.Fields{"path": path, "method": method}, "APIの実行終わります。")
	}
}
