package utils

import "github.com/gin-gonic/gin"

func ValidateRequestBody(ginContext *gin.Context, dto interface{}) error {
	err := ginContext.ShouldBindJSON(&dto)
	if err != nil {
		return BadRequestError("Dados Inválidos. Preencha os Campos Corretamente!")
	}
	return nil
}
