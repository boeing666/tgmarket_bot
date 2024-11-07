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

	FilterName      string
	FiltredProducts map[int64]*models.Product
}

func (u *User) AddProduct(shop int, url string, productid string) (*models.Product, error) {
	/*db := app.GetDB()

	mm := parser.MegaMarket()
	productInfo, err := mm.GetProductInfo(productid)
	if err != nil {
		return nil, err
	}

	productOffers, err := mm.GetOffers(productid)
	if err != nil {
		return nil, err
	}

	productPrice := 0
	productBonus := 0
	if len(productOffers.Offers) > 0 {
		productPrice = productOffers.Offers[0].FinalPrice
		productBonus = productOffers.Offers[0].BonusAmountFinalPrice
		for _, offer := range productOffers.Offers {
			if productPrice > offer.FinalPrice {
				productPrice = offer.FinalPrice
			}
			if productBonus > offer.BonusAmountFinalPrice {
				productBonus = offer.BonusAmountFinalPrice
			}
		}

	}

	product := models.Product{
		Name:      productInfo.Goods.Title,
		Price:     productPrice,
		Bonus:     productBonus,
		ProductID: productid,
		URL:       url,
		ShopID:    shop,
		UserID:    u.ID,
	}

	if err := db.Create(&product).Error; err != nil {
		return nil, err
	}

	u.Products[product.ID] = &product
	return &product, nil*/
	return nil, nil
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

func (user *User) FindProductByProductID(productID string) *models.Product {
	for _, product := range user.Products {
		if product.ProductID == productID {
			return product
		}
	}
	return nil
}
