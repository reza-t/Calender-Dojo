package controller

import (
	"calender/model"
	"errors"
	"github.com/google/uuid"
)

var CalenderCtrl ICalender

type calender struct{}

type ICalender interface {
	Save(calender model.Calender) (err error)
	GetAll() ([]model.Calender, error)
	DeleteByID(id string) error
}

func init() {
	CalenderCtrl = &calender{}
}

func (*calender) Save(calender model.Calender) (err error) {
	calender.ID = uuid.New().String()
	err = model.CalenderModel.Insert(calender)
	if err != nil {
		return err
	}
	return nil
}

func (*calender) GetAll() ([]model.Calender, error) {
	dates, err := model.CalenderModel.FindAll()
	if err != err {
		return []model.Calender{}, err
	}
	return dates, nil
}

func (*calender) DeleteByID(id string) error {
	if len(id) < 1 {
		return errors.New("please give Id as prop")
	}
	count, err := model.CalenderModel.FindAndDeleteByID(id)
	if count < 1 {
		return errors.New("there is not calender to be deleted")
	}
	if err != nil {
		return err
	}
	return nil
}
