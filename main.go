package main

import (
	"calender/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/v1/calender", api.GetCalender)
	r.PUT("/v1/calender", api.PutCalender)
	r.DELETE("/v1/calender/:id", api.DeleteCalender)
	r.Run()
}
