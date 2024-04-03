package test

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	router := gin.Default()
	return router
}
