package handler

import (
	"log"
	"net/http"

	"reservationbox-api/hotel"

	"github.com/gin-gonic/gin"
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

func HotelHandler(c *gin.Context) {
	var hotelInput hotel.HotelInput

	err := c.ShouldBindJSON(&hotelInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"hotel_name": hotelInput.Hotel_name,
		"address":    hotelInput.Address,
	})
}
