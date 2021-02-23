package storage

import (
	"crud-server/model"
	"sync"
)

type DB struct {
	mu  *sync.RWMutex
	all map[string]model.Man
}

type Data struct {
	*DB
}

func NewDB() *Data {
	return &Data{
		&DB{
			all: make(map[string]model.Man),
			mu:  new(sync.RWMutex),
		},
	}
}

func DelMan() {

}
