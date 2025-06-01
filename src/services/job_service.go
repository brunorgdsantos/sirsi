package services

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/utils"
	"context"
)

type JobService struct {
	repo *repositories.JobRepository
}

func NewJobService(repo *repositories.JobRepository) *JobService {
	return &JobService{repo: repo}
}

func (s *JobService) GetAllJobs(ctx context.Context) ([]interface{}, error) {
	return s.repo.GetAll(ctx)
}

func (service *JobService) CreateJob(titulo, descricao, localizacao string) error {
	user := dtos.Job{
		Titulo:      titulo,
		Descricao:   descricao,
		Localizacao: localizacao,
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err := service.repo.CreateJobsDb(contextServer, user)
	if err != nil {
		return err
	}
	return nil
}
