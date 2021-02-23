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
	data.All["57734684"] = model.Man{
		FirstName: "Andry",
		LastName:  "Pupkin",
	}
	//return &Data{
	//	all: make(map[string]model.Man),
	//	mu:  new(sync.RWMutex),
	//}
	return data
}

func DelMan() {

}
