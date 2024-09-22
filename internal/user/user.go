package user

import "tgmarket/internal/protobufs"

type Cache struct {
	State protobufs.UserState
}

func NewCache() map[int64]*Cache {
	return make(map[int64]*Cache)
}
