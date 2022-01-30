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

	// hotel := hotel.Hotel{}
	// hotel.Hotel_name = "Bobobox Alun Alun"
	// hotel.Address = "Jalan Kepatihan No 8, Balonggede, Bandung"

	// err = db.Create(&hotel).Error
	// if err != nil {
	// 	fmt.Println("Error creating hotel record")
	// }

	// roomtype := roomtype.Roomtype{}
	// roomtype.Name = "Earth Double"

	// err = db.Create(&roomtype).Error
	// if err != nil {
	// 	fmt.Println("Error creating roomtype record")
	// }

	// var hotel hotel.Hotel

	// err = db.First(&hotel).Error

	// if err != nil {
	// 	fmt.Println("Error finding hotel record")
	// }

	// fmt.Println("Hotel name", hotel.Hotel_name)
	// fmt.Println("Hotel object %v", hotel)

	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/book", handler.BookHandler)
	router.GET("/hotel", handler.GetHotelHandler)
	router.POST("/hotel", handler.CreateHotelHandler)

	router.Run()
}
