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
	"os"
)

// @title API de Autenticação e Tarefas SIRSI
// @version 1.0
// @description API do Sistema de Recrutamento e Seleção Inteligente (SIRSI)
// @host localhost:8001
// @BasePath /
func main() {

	uri := os.Getenv("MONGO_URI")
	dbname := "sirsi_database"
	//uri, dbname := "mongodb+srv://sirsi:12345@clustersirsi.nibjz9g.mongodb.net/?retryWrites=true&w=majority&appName=ClusterSirsi", "sirsi_database"
	repoUser, errUser := repositories.NewUserRepository(uri, dbname, "users")
	repoTestes, errTeste := repositories.NewJobRepository(uri, dbname, "vagas")
	repoCurriculo, errTeste := repositories.NewCurriculoRepository(uri, dbname, "curriculos")

	if errUser != nil {
		log.Fatalf("Erro no repositorio ao iniciar: errUser=%v ", errUser)
		return
	}

	server := gin.Default()

	//server.LoadHTMLGlob("src/templates/*.html") //Carrega templates HTML
	server.Static("/static", "src/static")

	server.Use(middlewares.ErrorMiddlewareHandler())

	controllers.NewUserController(server, repoUser)

	if errTeste != nil {
		log.Fatalf("Erro no repositorio ao iniciar: errUser=%v ", errUser)
		return
	}

	jobController := controllers.NewJobController(repoTestes)
	server.GET("/jobs", jobController.GetJobsApiSwagger)
	server.GET("/jobs/vagas", jobController.GetJobsPage)

	controllers.NewJobControllerCreate(server, repoTestes)
	controllers.NewJobControllerUpdate(server, repoTestes)
	controllers.NewJobControllerDelete(server, repoTestes)
	controllers.NewCurriculoController(server, repoCurriculo)

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	// @description Value: Bearer abc... (Bearer+space+token)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	server.Run(":8001")
}
