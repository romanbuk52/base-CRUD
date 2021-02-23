package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

type DataHandler struct {
}

// NewDataHandler 1
func (d *DataHandler) NewDataHandler() {
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
