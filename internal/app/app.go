package app

import (
	"tgmarket/internal/user"
	"time"

	"gorm.io/gorm"
)

type App struct {
	Database   *gorm.DB
	Cache      map[int64]*user.Cache
	LaunchTime uint64
}

var app *App

func GetApp() *App {
	if app == nil {
		app = &App{
			Cache: user.NewCache(), 
			LaunchTime: uint64(time.Now().Unix())
		}
	}
	return app
}

func GetDB() *gorm.DB {
	return GetApp().Database
}

func (app *App) GetUser(id int64) *user.Cache {
	user, ok := app.Cache[id]
	if ok {
		return user
	} else {
		cache := &user.Cache{State: cache.None}
		app.Cache[id] = cache
		return cache
	}
}