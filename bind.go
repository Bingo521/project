package main

import (
	"github.com/gin-gonic/gin"
	"my_project/handler"
	"my_project/logs"
)

const (
	GET  = 1
	POST = 2
	ANY  = 3
)

type Handler struct {
	Type int64
	Path string
	Func func(c *gin.Context)
}

var funcArr = []*Handler{
	&Handler{
		Type: ANY,
		Path: "/ping",
		Func: Ping,
	},
	&Handler{
		Type: ANY,
		Path: "/login",
		Func: Login,
	},
	&Handler{
		Type: ANY,
		Path: "/get_schools",
		Func: GetSchools,
	},
	&Handler{
		Type: ANY,
		Path: "/set_school",
		Func:SetSchool,
	},
	&Handler{
		Type: ANY,
		Path: "/get_course/current",
		Func: GetCurCourse,
	},
	&Handler{
		Type: ANY,
		Path: "/get_course/by_week",
		Func: GetCourseByWeek,
	},
}

func Bind(r *gin.Engine) {
	for _, handler := range funcArr {
		if handler.Type == GET {
			r.GET(handler.Path, handler.Func)
		} else if handler.Type == POST {
			r.POST(handler.Path, handler.Func)
		} else if handler.Type == ANY {
			r.GET(handler.Path, handler.Func)
			r.POST(handler.Path, handler.Func)
		}
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context) {
	h := handler.NewLogin()
	c.JSON(200, h.Execute(c))
}

func GetSchools(c *gin.Context) {
	logs.Info("get school")
	h := handler.NewGetSchoolHandler()
	c.JSON(200, h.Execute(c))
}

func SetSchool(c *gin.Context){
	h := handler.NewSetSchoolHandler(c)
	c.JSON(200, h.Execute())
}

func GetCurCourse(c *gin.Context){
	h := handler.NewGetCurCurrentCourseHandler(c)
	c.JSON(200,h.Execute())
}

func GetCourseByWeek(c * gin.Context){
	h := handler.NewGetCourseByWeekNumHandler(c)
	c.JSON(200,h.Execute())
}