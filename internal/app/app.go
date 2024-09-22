package app

import (
	"tgmarket/internal/cache"
	"tgmarket/internal/protobufs"
	"time"

	"gorm.io/gorm"
)

type App struct {
	Database   *gorm.DB
	Cache      map[int64]*cache.User
	LaunchTime uint64
}

var app *App

func GetApp() *App {
	if app == nil {
		app = &App{
			Cache:      make(map[int64]*cache.User),
			LaunchTime: uint64(time.Now().Unix()),
		}
	}
	return app
}

func GetDB() *gorm.DB {
	return GetApp().Database
}

func (app *App) GetUser(id int64) *cache.User {
	userdata, ok := app.Cache[id]
	if ok {
		return userdata
	} else {
		userdata := &cache.User{State: protobufs.UserState_None}
		app.Cache[id] = userdata
		return userdata
	}
}
