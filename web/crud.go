package web

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// DataHandler its struct datahandler
type DataHandler struct {
	HumanStorage HumanStorage
}

// HumanStorage methods HumanStorage
type HumanStorage interface {
	Add(Man) error
	Get(id string) (Man, error)
	GetAll() ([]Man, error)
	Edit(Man) error
	Del(id string) error
}

type Man struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
}

// NewDataHandler 1
func NewDataHandler(DB HumanStorage) *DataHandler {
	return &DataHandler{DB}
}

// MainPage write main page
func (dh *DataHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome in man database"))
}

// GetAllMan print all man in DB(map)
func (dh *DataHandler) GetAllMan(w http.ResponseWriter, r *http.Request) {
	// var cacheMap Man
	w.Header().Set("Content-Type", "application/json")
	data, err := dh.HumanStorage.GetAll()
	if err != nil {
		dh.SendError(w, http.StatusInternalServerError, err)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		dh.SendError(w, http.StatusInternalServerError, ErrJsonEncode)
		return
	}
}

// GetManByID get man by ID
func (dh *DataHandler) GetManByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// w.Write([]byte("Take man with id " + vars["manID"]))
	id := vars["manID"]
	man, err := dh.HumanStorage.Get(id)
	if err != nil {
		dh.SendError(w, http.StatusNotFound, ErrManNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(man); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CreateMan create new man in database
func (dh *DataHandler) CreateMan(w http.ResponseWriter, r *http.Request) {

	var NewMan Man // to appoint variables "NewMan" structure "man"

	// println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&NewMan); //записали з тіла запиту в змінну "с"
	err != nil {
		dh.SendError(w, http.StatusBadRequest, err)
		return
	}

	// Generate new ID
	NewMan.ID = uuid.New().String()
	if err := dh.HumanStorage.Add(NewMan); err != nil {
		dh.SendError(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateMan edit man by ID
func (dh *DataHandler) UpdateMan(w http.ResponseWriter, r *http.Request) {

	var editMan Man
	vars := mux.Vars(r)
	editMan.ID = vars["manID"]

	if err := json.NewDecoder(r.Body).Decode(&editMan); err != nil {
		dh.SendError(w, http.StatusBadRequest, err)
		return
	}

	if err := dh.HumanStorage.Edit(editMan); err != nil {
		dh.SendError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteMan deleted man for ID
func (dh *DataHandler) DeleteMan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["manID"]

	if err := dh.HumanStorage.Del(id); err != nil {
		dh.SendError(w, http.StatusInternalServerError, err)
		return
	}

	dh.SendResponse(w, http.StatusOK)
}
