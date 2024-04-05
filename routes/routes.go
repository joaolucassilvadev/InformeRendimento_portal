package routes

import (
	"example.com/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes atribui as rotas da aplicação ao roteador Gin.
// @Description A func Routes é responsável por atribuir minhas rotas e passar para meu executável.
func Routes(r *gin.Engine) {
	// roiute
	r.GET("/pdf/:pdf", controller.FuncA)
}

// Swagger configura a rota relacionada ao Swagger.
// @Description Essa rota é responsável por passar minhas rotas Swagger.
func Swagger(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
