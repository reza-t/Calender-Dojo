package api

import (
	"calender/controller"
	"calender/model"
	"github.com/gin-gonic/gin"
)

//GetCalender return all the meeting dates from the DB
func GetCalender(c *gin.Context) {
	findCalender, err := controller.CalenderCtrl.GetAll()
	if err != nil {
		c.JSON(201, gin.H{"message": err.Error()})
	}
	c.JSON(200, findCalender)
}

//PutCalender save the one meeting to the DB
func PutCalender(c *gin.Context) {
	var cal model.Calender

	if err := c.BindJSON(&cal); err == nil {
		err := controller.CalenderCtrl.Save(cal)

		if err != nil {
			c.JSON(201, gin.H{"message": err.Error()})
		}
		c.JSON(200, gin.H{"result": "OK"})
	}
}

//DeleteCalender find the meeting by ID and remove it from DB
func DeleteCalender(c *gin.Context) {
	id := c.Param("id")
	err := controller.CalenderCtrl.DeleteByID(id)
	if err != nil {
		c.JSON(201, gin.H{"message": err.Error()})
	}
	c.JSON(200, gin.H{"result": "OK"})
}
