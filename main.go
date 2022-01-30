package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "Diki Yuliawan",
			"project": "Bobobox Technical Test",
		})
	})

	router.Run()
}
