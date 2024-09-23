package app

import (
	"time"

	"gorm.io/gorm"
)

type App struct {
	Database   *gorm.DB
	LaunchTime uint64
}

var app *App

func GetApp() *App {
	if app == nil {
		app = &App{
			LaunchTime: uint64(time.Now().Unix()),
		}
	}
	return app
}

func GetDB() *gorm.DB {
	return GetApp().Database
}
