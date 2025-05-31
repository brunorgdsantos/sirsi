package main

import (
	_ "Sirsi/docs"
	"Sirsi/src/controllers"
	"Sirsi/src/repositories"
	"Sirsi/src/utils/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title API de Autenticação e Tarefas SIRSI
// @version 1.0
// @description API do Sistema de Recrutamento e Seleção Inteligente (SIRSI)
// @host localhost:8001
// @BasePath /
func main() {

	uri, dbname := "xxx", "sirsi_database"
	repoUser, errUser := repositories.NewUserRepository(uri, dbname, "users")
	//repoTask, errTask := repositories.NewTaskRepository(uri, dbname, "tasks")

	/*if errUser != nil || errTask != nil {
		log.Fatalf("Erro no repositorio ao iniciar: errUser=%v,errTask=%v ", errUser,
			errTask)
		return
	} */

	if errUser != nil {
		log.Fatalf("Erro no repositorio ao iniciar: errUser=%v ", errUser)
		return
	}

	server := gin.Default()
	server.Use(middlewares.ErrorMiddlewareHandler())

	controllers.NewUserController(server, repoUser)

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	// @description Value: Bearer abc... (Bearer+space+token)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	server.Run(":8001")
}
