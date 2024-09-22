package app

import (
	"sync"

	"gorm.io/gorm"
)

type Container struct {
	Database *gorm.DB
}

var (
	container *Container
	once      sync.Once
)

func (c *Container) Init(database *gorm.DB) {
	container.Database = database
}

func GetContainer() *Container {
	once.Do(func() {
		container = &Container{}
	})
	return container
}

func GetDB() *gorm.DB {
	return GetContainer().Database
}
