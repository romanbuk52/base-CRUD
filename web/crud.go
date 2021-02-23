package web

import (
	"crud-server/storage"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DataHandler struct {
	StorageData *storage.Data
}

// NewDataHandler 1
func NewDataHandler(DB *storage.Data) *DataHandler {
	return &DataHandler{DB}
}

// MainPage write main page
func (d *DataHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome in man database"))

}

// GetMan Get man with UUID
func (d *DataHandler) GetMan(w http.ResponseWriter, r *http.Request) {
	//	manID :=

	w.Write([]byte("It works!"))
}

// GetManById .
func (d *DataHandler) GetManById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//id, _ := strconv.Atoi(vars["manID"])
	w.Write([]byte("Take man with id " + vars["manID"]))
	err := json.NewEncoder(w).Encode(d.StorageData.All[vars["manID"]])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.Write((d.StorageData.All[vars["manID"]]))
}

//func (d *storage.Data) CreateMan(w http.ResponseWriter, r *http.Request) {
//
//	var NewMan man // to appoint variables "NewMan" structure "man"
//
//	err := json.NewDecoder(r.Body).Decode(&NewMan) //записали з тіла запиту в змінну "с"
//
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	manID := uuid.New().String()
//
//	manDB[manID] := NewMan
//
//	w.WriteHeader(http.StatusCreated)
//}
//
//func GetAllMan(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("getall"))
//}
//func EditMan(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("edit"))
//}
//func DeleteMan(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("delete"))
//}
