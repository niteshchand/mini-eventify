package models

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Location    string
	Date        time.Time
	Price       float64
	UserID      uint      // event creator
	CreatedAt   time.Time
}
