package main

import (
	"github.com/gin-gonic/gin"
	"my_project/comm"
)

func main() {
	r := gin.Default()
	r.Use(comm.Logger())
	r.Use(comm.SetSession())
	r.Use(comm.CheckLogin())

	Bind(r)
	r.Run()
}


