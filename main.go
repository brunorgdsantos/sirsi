package main

import (
	"Sirsi/src/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	uri, dbname := "xxx", "sirsi_database"
	repoUser, errUser := repositories.NewUserRepository(uri, dbname, "users")
	repoTask, errTask := repositories.NewTaskRepository(uri, dbname, "tasks"))

	if errUser != nil || errTask != nil {
		log.Fatalf("Erro no repositorio ao iniciar: errUser=%v,errTask=%v ", errUser,
			errTask)
		return
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin"})
	})

	app.Run(":8001")

}
