package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	router := setupRouter()

	err := router.Run(":5000")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}

}
