
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"Get student info successfully",
		})
	})
	r.POST("/post_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"Post student info successfully",
		})
	})
	r.PUT("/update_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"Update student info successfully",
		})
	})
	r.DELETE("/delete_student", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"Delete student info successfully",
		})
	})
	r.Run()
	
}