package service

import (
	"buzzGen/helpers"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	helpers.InitResource()
	m.Run()
	os.Exit(0)
}

func TestFetchHnData(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FetchHnData(); (err != nil) != tt.wantErr {
				t.Errorf("FetchHnData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
