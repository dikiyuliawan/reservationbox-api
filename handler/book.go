package handler

import (
	"fmt"
	"net/http"

	"reservationbox-api/hotel"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Diki Yuliawan",
		"project": "Bobobox Technical Test",
	})
}

func BookHandler(c *gin.Context) {
	checkin_date := c.Query("checkin_date")
	checkout_date := c.Query("checkout_date")
	room_qty := c.Query("room_qty")
	room_type_id := c.Query("room_type_id")

	c.JSON(http.StatusOK, gin.H{"checkin_date": checkin_date, "checkout_date": checkout_date, "room_qty": room_qty, "room_type_id": room_type_id})
}

// get list hotel
func GetHotelHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var hotel []hotel.Hotel

	err = db.Find(&hotel).Error

	if err != nil {
		fmt.Println("Error finding hotel record")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   hotel,
	})

}

// create new data hotel
func CreateHotelHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	data := hotel.Hotel{
		Hotel_name: c.PostForm("hotel_name"),
		Address:    c.PostForm("address"),
	}

	db.Create(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": "Data successfully created",
		"data":   data,
	})
}
