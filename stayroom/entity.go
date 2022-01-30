package stayroom

import "time"

type Stayroom struct {
	ID      int
	Stay_id int
	Room_id int
	Date    *time.Time
}
