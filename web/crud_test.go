package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDataHandler_GetManByID(t *testing.T) {
	request := httptest.NewRequest("http.MethodGet", "http://localhost:8080/man/123", nil)
	vars := map[string]string{
		"manID": "abcd",
	}
	request = mux.SetURLVars(request, vars)

	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "succes",
			r:             &http.Request{},
			expectedError: "",
			storage: &HumanStorageMock{
				GetFunc: func(id string) (Man, error) {
					return Man{
						ID: id,
					}, nil
				},
			},
		},
		{
			name:          "DB_error",
			r:             request,
			expectedError: ErrManNotFound.Error(),
			storage: &HumanStorageMock{
				GetFunc: func(id string) (Man, error) {
					return Man{}, errors.New("test")
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			recorder := httptest.NewRecorder()

			dh.GetManByID(recorder, tt.r)

			resultResponse := recorder.Result()
			if tt.expectedError != "" {
				var responseErr responseError
				assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseErr))
				assert.Equal(t, tt.expectedError, responseErr.Description)

				return
			}

			var responseAnswer Man
			assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseAnswer))
			assert.Equal(t, Man{}, responseAnswer)

			fmt.Printf("%d", len(tt.storage.GetCalls()))
		})
	}

}

func TestDataHandler_GetAllMan(t *testing.T) {
	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "succes",
			r:             &http.Request{},
			expectedError: "",
			storage: &HumanStorageMock{
				GetAllFunc: func() ([]Man, error) {
					return []Man{}, nil
				},
			},
		},
		{
			name:          "succes",
			r:             &http.Request{},
			expectedError: "DBError",
			storage: &HumanStorageMock{
				GetAllFunc: func() ([]Man, error) {
					return []Man{}, errors.New("DBError")
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			recorder := httptest.NewRecorder()
			dh.GetAllMan(recorder, tt.r)
			resultResponse := recorder.Result()
			if tt.expectedError != "" {
				var responseErr responseError
				assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseErr))
				assert.Equal(t, tt.expectedError, responseErr.Description)

				return
			}
			fmt.Printf("%d", len(tt.storage.GetAllCalls()))
		})
	}

}

func TestDataHandler_CreateMan(t *testing.T) {
	data1, err := json.Marshal(Man{})
	assert.NoError(t, err)

	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "test Create",
			r:             httptest.NewRequest(http.MethodPost, "/man", nil),
			expectedError: "",
			storage: &HumanStorageMock{
				AddFunc: func(Man) error {
					return nil
				},
			},
		},
		{
			name:          "test bad data in request, decoder error",
			r:             httptest.NewRequest(http.MethodPost, "/man", bytes.NewBuffer([]byte("test"))),
			expectedError: "invalid character 'e' in literal true (expecting 'r')",
			storage: &HumanStorageMock{
				AddFunc: func(Man) error {
					return nil

				},
			},
		},
		{
			name:          "test database error",
			r:             httptest.NewRequest(http.MethodPost, "/man", bytes.NewBuffer(data1)),
			expectedError: "DBError",
			storage: &HumanStorageMock{
				AddFunc: func(Man) error {
					return errors.New("DBError")
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			recorder := httptest.NewRecorder()
			dh.CreateMan(recorder, tt.r)
			resultResponse := recorder.Result()
			if tt.expectedError != "" {
				var responseErr responseError
				assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseErr))
				assert.Equal(t, tt.expectedError, responseErr.Description)
				// t.Fatal(tt.expectedError, responseErr.Description)

				return
			}
			var responseAnswer Man
			assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseAnswer))
			assert.Equal(t, Man{}, responseAnswer)

			fmt.Printf("%d", len(tt.storage.AddCalls()))
		})
	}

}

func TestDataHandler_UpdateMan(t *testing.T) {
	req0 := httptest.NewRequest(http.MethodPost, "/man", nil)

	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "editman",
			r:             req0,
			expectedError: "",
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
			recorder := httptest.NewRecorder()
			dh.UpdateMan(recorder, tt.r)
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
