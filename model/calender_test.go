package model

import (
	"reflect"
	"testing"
)

const testDataPath = "testData"

func Test_repository_FindAll(t *testing.T) {
	type fields struct {
		path string
	}
	var expected = []Calender{
		{
			ID:    "8248d540-16fe-4f6b-b5d8-e91084c14712",
			What:  "Titan",
			Who:   "Reza",
			Where: "Zoom",
			When:  "Tmr",
			Takes: 1,
		},
		{
			ID:    "8248d540-16fe-4f6b-b5d8-e91084c14713",
			What:  "Demo",
			Who:   "Adnan",
			Where: "Zoom",
			When:  "Tmr",
			Takes: 1,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Calender
		wantErr bool
	}{
		{
			name:    "hasData_allData",
			fields:  fields{path: testDataPath},
			want:    expected,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				path: tt.fields.path,
			}
			got, err := r.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_repository_FindAndDeleteByID(t *testing.T) {
//	type fields struct {
//		path string
//	}
//	type args struct {
//		id string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    int
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &repository{
//				path: tt.fields.path,
//			}
//			got, err := r.FindAndDeleteByID(tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("FindAndDeleteByID() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("FindAndDeleteByID() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_repository_Insert(t *testing.T) {
//	type fields struct {
//		path string
//	}
//	type args struct {
//		calender Calender
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &repository{
//				path: tt.fields.path,
//			}
//			if err := r.Insert(tt.args.calender); (err != nil) != tt.wantErr {
//				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
