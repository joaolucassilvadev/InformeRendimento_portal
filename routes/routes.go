package routes

import (
	"example.com/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine) {
	r.GET("/pdf/:pdf", controller.FuncA)
}

func Swagger(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
