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

func (service *JobService) CreateJob(title, description, area, category, location string, requirements []string, postedDate string) error {
	user := dtos.Vacancies{
		Title:        title,
		Description:  description,
		Area:         area,
		Category:     category,
		Location:     location,
		Requirements: requirements,
		PostedAt:     postedDate,
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err := service.repo.CreateJobsDb(contextServer, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *JobService) UpdateJob(id, title, description, location string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.BadRequestError("ID inválido")
	}

	user := dtos.Vacancies{
		Title:       title,
		Description: description,
		Location:    location,
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
