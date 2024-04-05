package routes

import (
	"example.com/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Description A func Routes é responsabel por atribuir minha rota e passar para meu executavel
func Routes(r *gin.Engine) {
	// roiute
	r.GET("/pdf/:pdf", controller.FuncA)
}

// @Description Essa rota é reponsavel por passar minhas rota swagger
func Swagger(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
