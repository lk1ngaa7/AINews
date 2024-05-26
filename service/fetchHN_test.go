package service

import (
	"buzzGen/helpers"
	"buzzGen/models"
	"os"
	"reflect"
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
func TestHandleSummary(t *testing.T) {
	type args struct {
		cate string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{"hn"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleSummary(tt.args.cate)
		})
	}
}

func Test_dealByGpt(t *testing.T) {
	type args struct {
		oriData models.TblOriData
	}
	tests := []struct {
		name             string
		args             args
		wantErr          error
		wantSum          string
		wantTitle        string
		wantTranslateStr string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr, gotSum, gotTitle, gotTranslateStr := dealByGpt(tt.args.oriData)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("dealByGpt() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if gotSum != tt.wantSum {
				t.Errorf("dealByGpt() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotTitle != tt.wantTitle {
				t.Errorf("dealByGpt() gotTitle = %v, want %v", gotTitle, tt.wantTitle)
			}
			if gotTranslateStr != tt.wantTranslateStr {
				t.Errorf("dealByGpt() gotTranslateStr = %v, want %v", gotTranslateStr, tt.wantTranslateStr)
			}
		})
	}
}
