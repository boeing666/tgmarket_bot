package cache

import (
	"tgmarket/internal/app"
	"tgmarket/internal/models"
	"tgmarket/internal/protobufs"
)

type User struct {
	ID          int64
	State       protobufs.UserState
	LastMsgID   int
	MenuMessage string

	Products map[int64]models.Product
}

func (u *User) AddProduct(product models.Product) error {
	db := app.GetDB()

	if err := db.Create(&product).Error; err != nil {
		return err
	}

	u.Products[product.ID] = product
	return nil
}

func (u *User) RemoveProduct(productID int64) error {
	db := app.GetDB()

	if err := db.Delete(&models.Product{ID: productID}).Error; err != nil {
		return err
	}

	delete(u.Products, productID)
	return nil
}

func (u *User) UpdateProduct(product models.Product) (bool, error) {
	if product, exists := u.Products[product.ID]; exists {
		db := app.GetDB()

		if err := db.Save(&product).Error; err != nil {
			return false, err
		}

		u.Products[product.ID] = product
		return true, nil
	}
	return false, nil
}
