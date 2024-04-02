package routes

import (
	"example.com/controller"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/pdf/:pdf", controller.FuncA)
}
