package controllers

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/services"
	"Sirsi/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(router *gin.Engine, repo *repositories.UserRepository) {
	service := services.NewUserService(repo)
	controller := &UserController{service: service}

	routes := router.Group("/users")
	{
		routes.POST("/users", controller.CreateUser)
	}
}

// @Tags users
// @Router /users [post]
// @Summary Criar um novo usuário
// @Description Registra um novo usuário na API
// @Accept json
// @Produce json
// @Param user body dtos.User true "Dados do usuário"
// @Success 201 {object} dtos.Message "Usuário criado"
// @Failure 400 {object} dtos.APIError "Erro de validação"
func (c *UserController) CreateUser(ginContext *gin.Context) {
	var userDto dtos.User

	err := utils.ValidateRequestBody(ginContext, &userDto)
	if err != nil {
		ginContext.Error(err)
		return
	}

	err = c.service.CreateUser(userDto.Email, userDto.Password)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, dtos.Message{
		Message: "Usuário criado com sucesso!",
	})
}
