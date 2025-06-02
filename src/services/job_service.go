package services

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *JobService) UpdateJob(id, titulo, descricao, localizacao string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.BadRequestError("ID inválido")
	}

	user := dtos.Job{
		Titulo:      titulo,
		Descricao:   descricao,
		Localizacao: localizacao,
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err = service.repo.UpdateJobsDb(contextServer, objID, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *JobService) DeleteJob(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.BadRequestError("ID inválido")
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err = service.repo.DeleteJobsDb(contextServer, objID)
	if err != nil {
		return err
	}
	return nil
}
