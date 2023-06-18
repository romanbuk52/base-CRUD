package usersstorage

import "net/http"

type DBUcrud interface {
	Add(DBUUser) error
	Get(NickName string) (DBUUser, error)
	GetAll() ([]DBUUser, error)
	Edit(NickName string) (DBUUser, error)
	Delete(NickName string) error
}

func (dbu *DBUHandler) Add(w http.ResponseWriter, r *http.Request) {

}
