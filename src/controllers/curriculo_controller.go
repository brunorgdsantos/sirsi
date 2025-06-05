package controllers

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/services"
	"Sirsi/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CurriculoController struct {
	service *services.CurriculoService
}

func NewCurriculoController(router *gin.Engine, repo *repositories.CurriculoRepository) {
	service := services.NewCurriculoService(repo)
	controller := &CurriculoController{service: service}

	routes := router.Group("/curriculo")
	{
		routes.POST("", controller.CreateCurriculo)

	}
}

// @Tags Create curriculo
// @Router /curriculo [post]
// @Summary Cadastra um novo curriculo
// @Description Cadastra um novo curriculo no banco
// @Accept json
// @Produce json
// @Param curriculo body dtos.Curriculo true "Dados do curriculo"
// @Success 201 {object} dtos.Message "Curriculo cadastrado"
// @Failure 400 {object} dtos.APIError "Erro de validação"
func (c *CurriculoController) CreateCurriculo(ginContext *gin.Context) {
	var curriculoDto dtos.Curriculo

	err := utils.ValidateRequestBody(ginContext, &curriculoDto)
	if err != nil {
		ginContext.Error(err)
		return
	}

	err = c.service.CreateCurriculo(curriculoDto.Nome, curriculoDto.Sobrenome, curriculoDto.Idade, curriculoDto.Email, curriculoDto.Formacao, curriculoDto.Experiencia)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, dtos.Message{
		Message: "Curriculo cadastrado com sucesso!",
	})
}
