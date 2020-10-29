package main

import (
	"sometest/router"
	"github.com/gin-gonic/gin"
	
	_"sometest/logger"
	"github.com/golang/glog"
	
	
)

func main() {

	glog.Infof("This is an info message")
	glog.Errorf("This is an error message ")
	glog.Warningf("This is a warning message")
	
	
	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
