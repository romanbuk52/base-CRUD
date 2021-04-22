package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataHandler_GetManByID(t *testing.T) {
	recorder := httptest.NewRecorder()
	tests := []struct {
		name    string
		storage *HumanStorageMock
		w       http.ResponseWriter
		r       *http.Request
	}{
		{
			name: "gratesucces",
			w:    recorder,
			r:    &http.Request{},
			storage: &HumanStorageMock{
				GetFunc: func(id string) (Man, error) {
					return Man{
						ID: id,
					}, nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			dh.GetManByID(tt.w, tt.r)
			fmt.Printf("%d", len(tt.storage.GetCalls()))
		})
	}

}

func TestDataHandler_GetAllMan(t *testing.T) {
	recorder := httptest.NewRecorder()
	tests := []struct {
		name    string
		storage *HumanStorageMock
		w       http.ResponseWriter
		r       *http.Request
	}{
		{
			name: "gratesucces getAll",
			w:    recorder,
			r:    &http.Request{},
			storage: &HumanStorageMock{
				GetAllFunc: func() ([]Man, error) {
					return []Man{}, nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			dh.GetAllMan(tt.w, tt.r)
			fmt.Printf("%d", len(tt.storage.GetAllCalls()))
		})
	}

}

func TestDataHandler_CreateMan(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/man", nil)
	tests := []struct {
		name    string
		storage *HumanStorageMock
		w       http.ResponseWriter
		r       *http.Request
	}{
		{
			name: "test Create",
			w:    recorder,
			r:    req,
			storage: &HumanStorageMock{
				AddFunc: func(Man) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			dh.CreateMan(tt.w, tt.r)
			fmt.Printf("%d", len(tt.storage.AddCalls()))
		})
	}

}

func TestDataHandler_UpdateMan(t *testing.T) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/man", nil)
	tests := []struct {
		name    string
		storage *HumanStorageMock
		w       http.ResponseWriter
		r       *http.Request
	}{
		{
			name: "editman",
			w:    recorder,
			r:    req,
			storage: &HumanStorageMock{
				EditFunc: func(Man) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			dh.UpdateMan(tt.w, tt.r)
			fmt.Printf("%d", len(tt.storage.EditCalls()))
		})
	}

}

func TestDataHandler_DeleteMan(t *testing.T) {
	recorder := httptest.NewRecorder()
	tests := []struct {
		name    string
		storage *HumanStorageMock
		w       http.ResponseWriter
		r       *http.Request
	}{
		{
			name: "test delete",
			w:    recorder,
			r:    &http.Request{},
			storage: &HumanStorageMock{
				DelFunc: func(id string) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			dh.DeleteMan(tt.w, tt.r)
			fmt.Printf("%d", len(tt.storage.DelCalls()))
		})
	}

}
