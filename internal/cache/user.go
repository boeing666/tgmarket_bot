package cache

import (
	"tgmarket/internal/app"
	"tgmarket/internal/models"
	"tgmarket/internal/protobufs"
	"time"
)

type User struct {
	ID          int64
	TelegramID  int64
	State       protobufs.UserState
	MenuMessage string

	ActiveProductID int64
	ActiveMsgID     int
	LastPage        int

	Products map[int64]*models.Product
}

func (u *User) AddProduct(shop int, url string) (*models.Product, error) {
	db := app.GetDB()

	product := models.Product{
		Name:   "Имя не задано",
		URL:    url,
		ShopID: shop,
		UserID: u.ID,
	}

	if err := db.Create(&product).Error; err != nil {
		return nil, err
	}

	u.Products[product.ID] = &product
	return &product, nil
}

func (u *User) RemoveProduct(productID int64) error {
	db := app.GetDB()

	if err := db.Delete(&models.Product{ID: productID}).Error; err != nil {
		return err
	}

	delete(u.Products, productID)
	return nil
}

func (u *User) UpdateProduct(product *models.Product) error {
	db := app.GetDB()

	product.UpdatedAt = time.Now()
	if err := db.Save(&product).Error; err != nil {
		return err
	}

	u.Products[product.ID] = product
	return nil
}
