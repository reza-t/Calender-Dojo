package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const filename = "calender"


type Calender struct {
	ID    string  `json:"id"`
	What  string  `json:"what"`
	Who   string  `json:"who"`
	Where string  `json:"where"`
	When  string  `json:"when"`
	Takes float32 `json:"takes"`
}

type repository struct{
	path string
}

func (r *repository) Insert(calender Calender) error {
	f, err := os.OpenFile(r.path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	b, err := json.Marshal(calender)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	f.WriteString("\n")

	return nil
}

func (r *repository) FindAll() ([]Calender, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return []Calender{}, err
	}
	defer file.Close()

	var calender []Calender
	var cal Calender

	scanner := bufio.NewScanner(file)

	// GetAll the file line by line and append it to the calender array
	for scanner.Scan() {
		err = json.Unmarshal(scanner.Bytes(), &cal)
		calender = append(calender, cal)
	}

	if err := scanner.Err(); err != nil {
		return []Calender{}, err
	}

	return calender, nil
}

func (r *repository) FindAndDeleteByID(id string) (int, error) {
	getCalender, err := r.FindAll()
	if err != nil {
		return 0, err
	}

	var deletedCount = 0

	var calenders []Calender
	//Find the record and remove it
	for i, v := range getCalender {
		if v.ID != id {
			fmt.Println(v)
			fmt.Println(i)
			calenders = append(calenders, v)
		} else {
			deletedCount ++
		}
	}

	err = os.Remove(r.path)
	if err != nil {
		return 0, err
	}

	for _, v := range calenders {
		r.Insert(v)
	}

	return deletedCount, nil
}

type ICalender interface {
	Insert(calender Calender) error
	FindAll() ([]Calender, error)
	FindAndDeleteByID(string) (int, error)
}

var CalenderModel ICalender

func init() {
	CalenderModel = &repository{
		path: filename,
	}
}
