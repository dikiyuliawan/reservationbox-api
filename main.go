package main

import (
	"fmt"
	"log"
	"reservationbox-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("Database connection success")

	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/book", handler.BookHandler)
	router.POST("/hotel", handler.HotelHandler)

	router.Run()
}
