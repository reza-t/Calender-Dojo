package controller

import (
	"calender/model"
	"testing"
)

type mockCalenderModel struct {}

var mockInsert func (model.Calender) error
var mockFindAll func () ([]model.Calender, error)
var mockDeleteByID func (string) (int, error)


func (m mockCalenderModel) Insert(calender model.Calender) error {
	return mockInsert(calender)
}

func (m mockCalenderModel) FindAll() ([]model.Calender, error) {
	return mockFindAll()
}

func (m mockCalenderModel) FindAndDeleteByID(id string) (int, error) {
	return mockDeleteByID(id)
}

func Test_calender_DeleteByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		mockDeleteById  func (string) (int, error)
		args    args
		wantErr bool
	}{
		{
			name: "IdNotProvided_error",
			mockDeleteById: func(string) (int, error) {
				return 0, nil
			},
			args: args{id: ""},
			wantErr: true,
		},
		{
			name: "deleted_numberOfDeletedRecord",
			mockDeleteById: func(string) (int, error) {
				return 1, nil
			},
			args: args{id: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		mockDeleteByID = tt.mockDeleteById
		t.Run(tt.name, func(t *testing.T) {
			ca := &calender{db: mockCalenderModel{}}
			if err := ca.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//func Test_calender_GetAll(t *testing.T) {
//	tests := []struct {
//		name    string
//		want    []model.Calender
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ca := &calender{}
//			got, err := ca.GetAll()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_calender_Save(t *testing.T) {
//	type args struct {
//		calender model.Calender
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ca := &calender{}
//			if err := ca.Save(tt.args.calender); (err != nil) != tt.wantErr {
//				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}