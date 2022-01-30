package price

import "time"

type Price struct {
	ID           int
	Date         time.Time
	Room_type_id int
	Price        int
}
