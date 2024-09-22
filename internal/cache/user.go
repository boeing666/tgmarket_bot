package cache

import (
	"tgmarket/internal/models"
	"tgmarket/internal/protobufs"
)

type User struct {
	ID       int64
	State    protobufs.UserState
	Products map[int64]models.Product
}
