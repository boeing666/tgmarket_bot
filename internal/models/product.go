package models

import (
	"time"
)

type Product struct {
	ID int64 `gorm:"primaryKey"`

	Name string `gorm:"type:varchar(64);"`
	URL  string `gorm:"type:varchar(256);"`

	ShopID int

	MinPrice   uint
	MinBonuses uint

	CreatedAt time.Time
	UpdatedAt time.Time

	UserID int64
}
