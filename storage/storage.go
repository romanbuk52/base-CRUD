package storage

import (
	"crud-server/web"
	"sync"
)

type Data struct {
	mu  *sync.RWMutex
	all map[string]web.Man
}

func NewDB() *Data {
	data := &Data{
		all: make(map[string]web.Man),
		mu:  new(sync.RWMutex),
	}
	data.all["577"] = web.Man{
		ID:        "577",
		FirstName: "Andry",
		LastName:  "Kamkin",
		Height:    200,
		Weight:    75,
	}
	return data
}

func (d *Data) Add(m web.Man) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.all[m.ID] = m

	return nil
}

func (d *Data) Get(id string) (web.Man, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	c := d.all[id]

	return c, nil
}

func (d *Data) GetAll() ([]web.Man, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	res := make([]web.Man, 0, len(d.all))
	for _, value := range d.all {
		res = append(res, value)
	}

	return res, nil
}

func (d *Data) Edit(m web.Man) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.all[m.ID] = m

	return nil
}

func (d *Data) Del(id string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	delete(d.all, id)

	return nil
}
