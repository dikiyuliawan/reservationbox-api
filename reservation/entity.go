package reservation

import "time"

type Reservation struct {
	ID                int
	Order_id          int
	Customer_name     string
	Booked_room_count int
	checkin_date      *time.Time
	checkout_date     *time.Time
	hotel_id          int
}
