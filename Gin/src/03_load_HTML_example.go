package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*")
	r.GET("/demo",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"name":"admin",
			"pwd":"123456",
		})
	})
	r.Run()
}