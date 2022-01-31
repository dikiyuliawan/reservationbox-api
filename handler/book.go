package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"reservationbox-api/hotel"
	"reservationbox-api/price"
	"reservationbox-api/promo"
	"reservationbox-api/room"
	"reservationbox-api/roomtype"

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
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	checkin_date := c.Query("checkin_date")
	checkout_date := c.Query("checkout_date")
	room_qty := c.Query("room_qty")
	room_type_id := c.Query("room_type_id")

	type Result struct {
		ID          int `json:"room_id"`
		Price       int `json:"price"`
		Room_number int `json:"room_number"`
	}

	result := []Result{}

	var total_price int

	err := db.Model(&roomtype.Roomtype{}).Select("roomtypes.name, prices.price, rooms.room_number, rooms.room_status, rooms.id").Joins("left join prices on prices.room_type_id = roomtypes.id").Joins("left join rooms on rooms.room_type_id = roomtypes.id").Where("roomtypes.id = ? AND prices.date BETWEEN ? AND ?", room_type_id, checkin_date, checkout_date).Find(&result)

	for _, pr := range result {
		total_price = pr.Price
	}

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"room_qty":       room_qty,
		"room_type_id":   room_type_id,
		"checkin_date":   checkin_date,
		"checkout_date":  checkout_date,
		"total_price":    total_price,
		"available_room": result,
	})
}

// hotel handler
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

// end hotel handler

// Room type handler
func RoomtypeHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var roomtype []roomtype.Roomtype

	err = db.Find(&roomtype).Error

	if err != nil {
		fmt.Println("Error finding Room Type record")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Succesfully get data room type",
		"data":   roomtype,
	})

}

// end roomtype handler

func PriceHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var price []price.Price

	err = db.Find(&price).Error

	if err != nil {
		fmt.Println("Error get price record")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Successfully get data price",
		"data":   price,
	})

}

func CreatePriceHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	room_type_id, _ := strconv.Atoi(c.PostForm("room_type_id"))
	p, _ := strconv.Atoi(c.PostForm("price"))
	date, _ := time.Parse(time.RFC3339, c.PostForm("date"))

	data := price.Price{
		Date:         date,
		Room_type_id: room_type_id,
		Price:        int(p),
	}

	err = db.Create(&data).Error
	if err != nil {
		fmt.Println("Error creating price record")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Data successfully created",
		"data":   data,
	})
}

func RoomHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	hotel_id, _ := strconv.Atoi(c.PostForm("hotel_id"))
	room_type_id, _ := strconv.Atoi(c.PostForm("room_type_id"))
	room_number, _ := strconv.Atoi(c.PostForm("room_number"))

	data := room.Room{
		Hotel_id:     int(hotel_id),
		Room_type_id: int(room_type_id),
		Room_number:  room_number,
		Room_status:  c.PostForm("room_status"),
	}

	err = db.Create(&data).Error
	if err != nil {
		fmt.Println("Error creating room record")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Data successfully created",
		"data":   data,
	})

}

//promo handler
func PromoHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	type Result struct {
		Name        string
		Promo_name  string
		Promo       int
		Price       int
		Final_price int
	}

	result := []Result{}

	err := db.Model(&roomtype.Roomtype{}).Select("roomtypes.name, prices.price, promos.promo_name, promos.promo, prices.price - promos.promo as final_price").Joins("left join prices on prices.room_type_id = roomtypes.id").Joins("left join promos on promos.room_type_id = roomtypes.id").Find(&result)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func CreatePromoHandler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/reservationbox-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	room_type_id, _ := strconv.Atoi(c.PostForm("room_type_id"))
	promo_name := c.PostForm("promo_name")
	promo_, _ := strconv.Atoi(c.PostForm("promo"))

	data := promo.Promo{
		Room_type_id: room_type_id,
		Promo_name:   promo_name,
		Promo:        promo_,
	}

	err = db.Create(&data).Error

	if err != nil {
		fmt.Println("Error creating promo record")
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Data successfully created",
		"data":   data,
	})

}

//end promo handler
