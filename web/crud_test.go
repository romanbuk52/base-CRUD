package web

import (
	"net/http"
	"testing"
)

func TestDataHandler_GetManByID(t *testing.T) {
	tests := []struct {
		name   string
		storage *HumanStorageMock
		w http.ResponseWriter
		r *http.Request
	}{
		{
			name: "gratesucces",
			storage: &HumanStorageMock{
							GetFunc: func(id string) (Man, error) {
								
								panic("mock out the Get method")
							},
						
		},
		// TODO: Add test cases.
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dh := &DataHandler{
				HumanStorage: tt.storage,
			}
			dh.GetManByID(tt.args.w, tt.args.r)
			fmt.Printf("%d", len(tt.storage.GetCalls()))
		})
	}
}
