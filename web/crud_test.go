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
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/man/123", nil)
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
			if tt.expectedError == "" {
				var responseAnswer Man
				assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseAnswer))
				assert.Equal(t, Man{}, responseAnswer)

				return
			}
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
	testStruct := Man{
		FirstName: "tra",
		LastName:  "ta",
		Height:    111,
		Weight:    222,
	}
	data1, err := json.Marshal(testStruct)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "test_Create and test generate UUID",
			r:             httptest.NewRequest(http.MethodPost, "/man", bytes.NewReader(data1)),
			expectedError: "",
			storage: &HumanStorageMock{
				AddFunc: func(m Man) error {
					assert.NotEmpty(t, m.ID)
					m.ID = ""
					assert.Equal(t, testStruct, m)

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

				return
			}

			fmt.Printf("%d", len(tt.storage.AddCalls()))
		})
	}

}

func TestDataHandler_UpdateMan(t *testing.T) {
	testStruct := Man{
		ID:        "123",
		FirstName: "tra",
		LastName:  "ta",
		Height:    111,
		Weight:    222,
	}
	muxVars := map[string]string{
		"manID": "test-UUID",
	}
	data1, err := json.Marshal(testStruct)
	assert.NoError(t, err)

	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "editman",
			r:             mux.SetURLVars(httptest.NewRequest(http.MethodPut, "http://192.168.13.10:8080/man/test-UUID", bytes.NewReader(data1)), muxVars),
			expectedError: "",
			storage: &HumanStorageMock{
				EditFunc: func(m Man) error {
					assert.NotEmpty(t, m.ID)
					assert.Equal(t, testStruct, m)

					return nil
				},
			},
		},
		{
			name:          "test bad data in request, decoder error",
			r:             httptest.NewRequest(http.MethodPut, "http://192.168.13.10:8080/man/123", bytes.NewReader([]byte("test"))),
			expectedError: "invalid character 'e' in literal true (expecting 'r')",
			storage: &HumanStorageMock{
				EditFunc: func(m Man) error {
					return nil
				},
			},
		},
		{
			name:          "test database error",
			r:             mux.SetURLVars(httptest.NewRequest(http.MethodPut, "http://192.168.13.10:8080/man/test-UUID", bytes.NewReader(data1)), muxVars),
			expectedError: "DBError",
			storage: &HumanStorageMock{
				EditFunc: func(m Man) error {
					return errors.New("DBError")
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			recorder := httptest.NewRecorder()
			dh.UpdateMan(recorder, tt.r)
			resultResponse := recorder.Result()

			if tt.expectedError != "" {
				var responseErr responseError
				assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseErr))
				assert.Equal(t, tt.expectedError, responseErr.Description)

				return
			}

			if tt.expectedError == "" {
				return
			}

			fmt.Printf("%d", len(tt.storage.EditCalls()))
		})
	}

}

func TestDataHandler_DeleteMan(t *testing.T) {
	muxVars := map[string]string{
		"manID": "test-UUID",
	}

	tests := []struct {
		name          string
		storage       *HumanStorageMock
		r             *http.Request
		expectedError string
	}{
		{
			name:          "test delete",
			r:             mux.SetURLVars(httptest.NewRequest(http.MethodDelete, "http://192.168.13.10:8080/man/test-UUID", nil), muxVars),
			expectedError: "",
			storage: &HumanStorageMock{
				DelFunc: func(id string) error {
					println("id:", id)
					assert.Equal(t, id, "test-UUID")

					return nil
				},
			},
		},
		{
			name:          "test DB error",
			r:             mux.SetURLVars(httptest.NewRequest(http.MethodDelete, "http://192.168.13.10:8080/man/test-UUID", nil), muxVars),
			expectedError: "man not found",
			storage: &HumanStorageMock{
				DelFunc: func(id string) error {

					return errors.New("man not found")
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			recorder := httptest.NewRecorder()
			dh.DeleteMan(recorder, tt.r)
			resultResponse := recorder.Result()

			if tt.expectedError != "" {
				var responseErr responseError
				assert.NoError(t, json.NewDecoder(resultResponse.Body).Decode(&responseErr))
				assert.Equal(t, tt.expectedError, responseErr.Description)

				return
			}

			if tt.expectedError == "" {
				return
			}

			fmt.Printf("%d", len(tt.storage.DelCalls()))
		})
	}

}
