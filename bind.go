package main

import (
	"github.com/gin-gonic/gin"
	"my_project/handler"
)

const (
	GET = 1
	POST = 2
	ANY = 3
)

type Handler struct {
	Type int64
	Path string
	Func func(c *gin.Context)
}

var funcArr = []*Handler{
	&Handler{
		Type:ANY,
		Path:"/ping",
		Func:Ping,
	},
	&Handler{
		Type:ANY,
		Path:"/login",
		Func:Login,
	},
}

func Bind(r *gin.Engine){
	for _,handler := range funcArr{
		if handler.Type == GET{
			r.GET(handler.Path,handler.Func)
		}else if handler.Type == POST{
			r.POST(handler.Path,handler.Func)
		}else if handler.Type == ANY{
			r.GET(handler.Path,handler.Func)
			r.POST(handler.Path,handler.Func)
		}
	}
}

func Ping(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context){
	h:=handler.NewLogin(c)
	c.JSON(200, h.Execute())
}


