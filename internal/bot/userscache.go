package bot

import (
	"tgmarket/internal/app"
	"tgmarket/internal/cache"
	"tgmarket/internal/models"
	"tgmarket/internal/protobufs"
)

type usersCache struct {
	users map[int64]*cache.User
}

var userscache usersCache

func initUsersCache() {
	userscache.users = make(map[int64]*cache.User)
}

func (u usersCache) getUser(id int64) (*cache.User, error) {
	userdata, ok := u.users[id]
	if ok {
		return userdata, nil
	}

	db := app.GetDB()

	user := models.User{TelegramID: id}
	err := db.Where(user).Assign(user).FirstOrCreate(&user).Error
	if err != nil {
		return nil, err
	}

	products, err := getUserProducts(user.ID)
	if err != nil {
		return nil, err
	}

	userdata = &cache.User{
		ID:        user.ID,
		State:     protobufs.UserState_None,
		LastMsgID: 0,
		Products:  products,
	}

	u.users[id] = userdata

	return userdata, nil
}

func getUserProducts(userID int64) (map[int64]models.Product, error) {
	var products []models.Product
	db := app.GetDB()

	if err := db.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}

	productMap := make(map[int64]models.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	return productMap, nil
}
