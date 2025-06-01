package services

import (
	"Sirsi/src/repositories"
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
