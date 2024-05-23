package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @Summary Retorna um arquivo PDF com o cpf fornecido.
// Essaa rota retorna o conteudo de um arquivo pdf com o nome especificado
// @Produce application/cpf
// @Param pdf path string true "Numero do cpf"
// @Success 200 {file} application/pdf
// @Failure 404 {object} map[string]string
// @Router /pdf/{pdf} [get]
func FuncA(ctx *gin.Context) {
	pdf := ctx.Param("pdf")

	filePath := filepath.Join("/home/joao/leraquivo/InformeRendimento_portal/paginas/", pdf+".pdf")
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Arquivo não encontrado"})
		return
	}

	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", pdf))

	ctx.Data(http.StatusOK, "application/pdf", fileBytes)
}
