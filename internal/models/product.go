package models

import (
	"time"
)

type Product struct {
	ID int64 `gorm:"primaryKey"`

	Name      string `gorm:"type:varchar(256);"`
	URL       string `gorm:"type:varchar(256);"`
	ProductID string `gorm:"type:varchar(32);"`

	Price  int
	Bonus  int
	ShopID int

	MinPrice   int
	MinBonuses int

	CreatedAt time.Time
	UpdatedAt time.Time

	UserID int64
}
