package services

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/utils"
)

type CurriculoService struct {
	repo *repositories.CurriculoRepository
}

func NewCurriculoService(repo *repositories.CurriculoRepository) *CurriculoService {
	return &CurriculoService{repo: repo}
}

func (service *CurriculoService) CreateCurriculo(nome, sobrenome, formacao, experiencia string) error {
	curriculo := dtos.Curriculo{
		Nome:        nome,
		Sobrenome:   sobrenome,
		Formacao:    formacao,
		Experiencia: experiencia,
	}

	contextServer := utils.CreateContextServerWithTimeout()

	err := service.repo.Create(contextServer, curriculo)
	if err != nil {
		return err
	}
	return nil
}
