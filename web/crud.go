package web

import (
	"crud-server/model"
	"crud-server/storage"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
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
func (dh *DataHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome in man database"))

}

// GetMan Get man with
func (dh *DataHandler) GetAllMan(w http.ResponseWriter, r *http.Request) {
	//	manID :=

	w.Write([]byte("It works!"))
}

// GetManByID get man by ID
func (dh *DataHandler) GetManByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// w.Write([]byte("Take man with id " + vars["manID"]))
	id := vars["manID"]
	man, ok := dh.StorageData.All[id]
	if !ok {
		dh.SendError(w, http.StatusNotFound, ErrManNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(man)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CreateMan create new man in database
func (dh *DataHandler) CreateMan(w http.ResponseWriter, r *http.Request) {

	var NewMan model.Man // to appoint variables "NewMan" structure "man"

	// println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&NewMan); //записали з тіла запиту в змінну "с"
	err != nil {
		dh.SendError(w, http.StatusBadRequest, err)
		return
	}

	// Generate new ID
	manID := uuid.New().String()
	if _, err := dh.StorageData.All[manID]; err != false {
		w.WriteHeader(http.StatusBadRequest)
	}

	dh.StorageData.All[manID] = NewMan

	w.WriteHeader(http.StatusCreated)
}

// UpdateMan edit man by ID
func (dh *DataHandler) UpdateMan(w http.ResponseWriter, r *http.Request) {

	var editMan model.Man
	vars := mux.Vars(r)
	id := vars["manID"]

	if err := json.NewDecoder(r.Body).Decode(&editMan); err != nil {
		dh.SendError(w, http.StatusBadRequest, err)
		return
	}

	dh.StorageData.All[id] = editMan

	w.WriteHeader(http.StatusCreated)
}

// DeleteMan deleted man for ID
func (dh *DataHandler) DeleteMan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["manID"]

	_, ok := dh.StorageData.All[id]
	if ok == false {
		dh.SendError(w, http.StatusNotFound, ErrManNotFound)
		return
	}

	delete(dh.StorageData.All, id)

	w.WriteHeader(http.StatusOK)
}
