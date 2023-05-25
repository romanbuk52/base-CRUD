package usersstorage

import (
	"time"

	"gorm.io/gorm"
)

// struct users with DB. DBUUser - database user
type DBUUser struct {
	AccesLevel        string
	AccessToken       string
	AccessTokenSecret string
	AvatarURL         string
	Description       string
	ExpiresAt         time.Time
	FirstName         string
	LastName          string
	IDToken           string
	NickName          string `gorm:"primaryKey"`
	Provider          string
	RefreshToken      string
	UserID            string
}

type DBU struct {
	db *gorm.DB
}

func NewDBU(db *gorm.DB) *DBU {
	DataDBU := &DBU{
		db: db,
	}
	if err := db.AutoMigrate(&DBUUser{}); err != nil {
		println(err)
	}

	return DataDBU
}

type DBUHandler struct {
	userStorage *DBU
}

func NewDBUHandler(db *DBU) *DBUHandler {
	return &DBUHandler{db}
}
