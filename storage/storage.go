package storage

import (
	"crud-server/model"
	"sync"
)

type Data struct {
	Mu  *sync.RWMutex
	All map[string]model.Man
}

func NewDB() *Data {
	data := &Data{
		All: make(map[string]model.Man),
		Mu:  new(sync.RWMutex),
	}
	data.All["577"] = model.Man{
		ID:        "577",
		FirstName: "Andry",
		LastName:  "Kamkin",
		Height:    200,
		Weight:    75,
	}
	return data
	// return &Data{
	// 	all: make(map[string]model.Man),
	// 	mu:  new(sync.RWMutex),
	// }
}

func DelMan() {

}
