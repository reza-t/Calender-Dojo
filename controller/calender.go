package controller

import (
	"calender/model"
	"errors"
	"github.com/google/uuid"
)

var CalenderCtrl ICalender

type calender struct{
	db model.ICalender
}

type ICalender interface {
	Save(calender model.Calender) (err error)
	GetAll() ([]model.Calender, error)
	DeleteByID(id string) error
}

func init() {
	CalenderCtrl = &calender{
		db: model.CalenderModel,
	}
}

func (c *calender) Save(calender model.Calender) (err error) {
	calender.ID = uuid.New().String()
	err = c.db.Insert(calender)
	if err != nil {
		return err
	}
	return nil
}

func (c *calender) GetAll() ([]model.Calender, error) {
	dates, err := c.db.FindAll()
	if err != err {
		return []model.Calender{}, err
	}
	return dates, nil
}

func (c *calender) DeleteByID(id string) error {
	if len(id) < 1 {
		return errors.New("please give Id as prop")
	}
	count, err := c.db.FindAndDeleteByID(id)
	if count < 1 {
		return errors.New("there is not calender to be deleted")
	}
	if err != nil {
		return err
	}
	return nil
}
