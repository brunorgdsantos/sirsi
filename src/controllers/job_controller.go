package controllers

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/services"
	"Sirsi/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobController struct {
	service *services.JobService
}

func NewJobController(repo *repositories.JobRepository) *JobController {
	service := services.NewJobService(repo)
	return &JobController{service}
}

func NewJobControllerCreate(router *gin.Engine, repo *repositories.JobRepository) {
	service := services.NewJobService(repo)
	controller := &JobController{service: service}

	routes := router.Group("/cadastrar/vagas")
	{
		routes.POST("", controller.CreateJobs)
	}
}

func NewJobControllerUpdate(router *gin.Engine, repo *repositories.JobRepository) {
	service := services.NewJobService(repo)
	controller := &JobController{service: service}

	routes := router.Group("/atualizar/vagas")
	{
		routes.PUT("", controller.UpdateJobs)
	}
}

func NewJobControllerDelete(router *gin.Engine, repo *repositories.JobRepository) {
	service := services.NewJobService(repo)
	controller := &JobController{service: service}

	routes := router.Group("/deletar/vagas")
	{
		routes.DELETE("", controller.DeleteJobs)
	}
}

// @Tags Read jobs
// @Router /jobs [get]
// @Summary Retorna vagas
// @Description Retorna todas as vagas diponíveis no Banco (Colletion Testes)
// @Accept json
// @Produce json
// @Success 201 {object} dtos.Message "Vagas retornadas com sucesso!"
// @Failure 400 {object} dtos.APIError "Erro ao retornar as vagas"
func (ctrl *JobController) GetJobsApiSwagger(c *gin.Context) {
	jobs, err := ctrl.service.GetAllJobs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Erro ao buscar vagas"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

// @Tags Create Jobs
// @Router /cadastrar/vagas [post]
// @Summary Criar uma nova vaga
// @Description Registra uma novo vaga no banco
// @Accept json
// @Produce json
// @Param job body dtos.Job true "Dados da vaga"
// @Success 201 {object} dtos.Message "Vaga criado com sucesso"
// @Failure 400 {object} dtos.APIError "Erro de validação"
func (c *JobController) CreateJobs(ginContext *gin.Context) {
	var jobDto dtos.Job

	err := utils.ValidateRequestBody(ginContext, &jobDto)
	if err != nil {
		ginContext.Error(err)
		return
	}

	err = c.service.CreateJob(jobDto.Titulo, jobDto.Descricao, jobDto.Localizacao)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, dtos.Message{
		Message: "Vaga criada com sucesso!",
	})
}

// @Tags Update Jobs
// @Router /atualizar/vagas [put]
// @Summary atualizar uma vaga
// @Description Atualziar uma vaga no banco
// @Accept json
// @Produce json
// @Param job body dtos.Job true "Dados da vaga"
// @Success 201 {object} dtos.Message "Vaga atualizada com sucesso"
// @Failure 400 {object} dtos.APIError "Erro de validação"
func (c *JobController) UpdateJobs(ginContext *gin.Context) {
	var jobDto dtos.Job

	err := utils.ValidateRequestBody(ginContext, &jobDto)
	if err != nil {
		ginContext.Error(err)
		return
	}

	err = c.service.UpdateJob(jobDto.Id, jobDto.Titulo, jobDto.Descricao, jobDto.Localizacao)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, dtos.Message{
		Message: "Vaga atualizada com sucesso!",
	})
}

// @Tags Deletar Jobs
// @Router /deletar/vagas [delete]
// @Summary Deletar uma vaga
// @Description Deletar uma vaga no banco
// @Accept json
// @Produce json
// @Param job body dtos.JobDelete true "Dados da vaga"
// @Success 201 {object} dtos.Message "Vaga atualizada com sucesso"
// @Failure 400 {object} dtos.APIError "Erro de validação"
func (c *JobController) DeleteJobs(ginContext *gin.Context) {
	var jobDto dtos.Job

	err := utils.ValidateRequestBody(ginContext, &jobDto)
	if err != nil {
		ginContext.Error(err)
		return
	}

	err = c.service.DeleteJob(jobDto.Id)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, dtos.Message{
		Message: "Vaga deletada com sucesso!",
	})
}

func (ctrl *JobController) GetJobsPage(c *gin.Context) {
	jobs, err := ctrl.service.GetAllJobs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar vagas"})
		return
	}

	c.HTML(http.StatusOK, "jobs.html", gin.H{
		"title": "Vagas Disponíveis",
		"jobs":  jobs,
	})
}
