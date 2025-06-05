package services

import (
	"Sirsi/src/dtos"
	"Sirsi/src/repositories"
	"Sirsi/src/utils"
	"fmt"
)

type CurriculoService struct {
	repo *repositories.CurriculoRepository
}

func NewCurriculoService(repo *repositories.CurriculoRepository) *CurriculoService {
	return &CurriculoService{repo: repo}
}

func (service *CurriculoService) CreateCurriculo(nome, sobrenome string, idade int, email string, formacao []dtos.Education, experiencia []string) error {
	curriculo := dtos.Curriculo{
		Nome:        nome,
		Sobrenome:   sobrenome,
		Idade:       idade,
		Email:       email,
		Formacao:    formacao,
		Experiencia: experiencia,
	}
	fmt.Print(curriculo)

	contextServer := utils.CreateContextServerWithTimeout()

	err := service.repo.Create(contextServer, curriculo)
	if err != nil {
		return err
	}
	return nil
}
