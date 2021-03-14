package storageSQL

import (
	"crud-server/web"

	"gorm.io/gorm"
)

// Human it`s gorm structure
type Human struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"firstName"`
	LastName  string `gorm:"lastName"`
	Height    int    `gorm:"height"`
	Weight    int    `gorm:"weight"`
}

// // asWebModel converter database model >> webModel
// func asWebModel(h *Human) struct {
// 	m.ID = h.ID
// 	m.FirstName = h.FirstName
// 	m.LastName = h.LastName
// 	m.Height = h.Height
// 	m.Weight = h.Weight

// 	return
// }

// asWebModel converter database model >> webModel
func asWebModel(h *Human) (m web.Man) {
	m.ID = h.ID
	m.FirstName = h.FirstName
	m.LastName = h.LastName
	m.Height = h.Height
	m.Weight = h.Weight

	return
}

// fromWebModel converter webModel >> database model
func fromWebModel(m web.Man) (h Human) {
	h.ID = m.ID
	h.FirstName = m.FirstName
	h.LastName = m.LastName
	h.Height = m.Height
	h.Weight = m.Weight

	return
}

// Data 000
type Data struct {
	// mu  *sync.RWMutex
	db *gorm.DB
}


// NewDB it`s new *Data
func NewDB(db *gorm.DB) *Data {
	data := &Data{
		db: db,
	}
	if err := db.AutoMigrate(&Human{}); err != nil {
		println(err)
	}
	return data
}

// Add func add nem Man
func (d *Data) Add(m web.Man) error {
	newHuman := fromWebModel(m)
	res := d.db.Create(&newHuman)

	return res.Error
}

// Get func get man by id witch SQL
func (d *Data) Get(id string) (m web.Man, er error) {
	var readMan Human

	res := d.db.First(readMan, id)
	m = asWebModel(&readMan)

	return m, res.Error
}

func (d *Data) GetAll() (m []web.Man, er error) {
	var readHumans Human
	var convertMan web.Man
	result := d.db.Find(&readHumans)
	convertMan = asWebModel(&readHumans)
	sliceOut := make([]web.Man, 0, int(result.RowsAffected))
	sliceOut = append(sliceOut, convertMan)

	return sliceOut, result.Error
}

func (d *Data) Edit(m web.Man) error {
	newHuman := fromWebModel(m)
	result := d.db.Model(newHuman.ID).Updates(newHuman)

	return result.Error
}

func (d *Data) Del(id string) error {
	var human Human
	result := d.db.Delete(human, id)

	return result.Error
}
