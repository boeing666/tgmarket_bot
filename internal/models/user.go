package models

import "time"

type User struct {
	ID         int64 `gorm:"primaryKey"`
	TelegramID int64
	CreatedAt  time.Time
	Products   []Product
}
