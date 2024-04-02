package main

import (
	"example.com/docs"
	"example.com/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/your_username/your_project/docs" // Substitua com o caminho do seu projeto
)

// @title Sua API Swagger
// @description Esta é uma API de exemplo utilizando Swagger e Gin
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	service := gin.Default()

	// Configurando informações do Swagger
	docs.SwaggerInfo.Title = "Api Informe de rendimento"
	docs.SwaggerInfo.Description = "Esta api tem um get a qual lhe devolve o informe de rendimento mediante ao cpf"
	docs.SwaggerInfo.Version = "1.0"
	routes.Routes(service)
	// Rota para servir a documentação Swagger
	service.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rota para a funcionalidade de retorno de PDF

	service.Run()
}
