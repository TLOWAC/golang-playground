package storage

import "module/types"

type Storage interface {
	Get(int) *types.User
}
