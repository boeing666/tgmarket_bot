package app

import (
	"tgmarket/internal/protobufs"
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
			Cache:      user.NewCache(),
			LaunchTime: uint64(time.Now().Unix()),
		}
	}
	return app
}

func GetDB() *gorm.DB {
	return GetApp().Database
}

func (app *App) GetUser(id int64) *user.Cache {
	userdata, ok := app.Cache[id]
	if ok {
		return userdata
	} else {
		cache := &user.Cache{State: protobufs.UserState_None}
		app.Cache[id] = cache
		return cache
	}
}
