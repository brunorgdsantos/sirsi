package controllers

import (
	"Sirsi/src/repositories"
	"Sirsi/src/services"
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
