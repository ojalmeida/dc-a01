package entities

import "time"

type Item struct {
	ID   int       `db:"id"`
	Text string    `db:"col_texto"`
	Date time.Time `db:"col_dt"`
}
