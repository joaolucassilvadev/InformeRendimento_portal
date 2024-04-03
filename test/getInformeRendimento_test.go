package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/controller"
)

func TestGetInformeRendimento(t *testing.T) {
	r := SetRouter()
	r.GET("/pdf/:pdf", controller.FuncA)
	// bom aqui passamos os parametros para nossa requisição, falamos que ela é um get
	req, _ := http.NewRequest("GET", "/pdf/62020735350", nil)
	// Bom aqui nós gravamos a resposta do Newrequest e passamos esse valor para o w
	w := httptest.NewRecorder()
	// bom aqui simulamos o servidor assim passando o w e o req a qual o req faz o request e o w grava o resultado da solicitação
	r.ServeHTTP(w, req)
	// bom aqui ferificamos se o new request é igual a 200, botamos esse code para pegarmos o codigo da nossa requisição http
	if w.Code != http.StatusOK {
		t.Errorf("Era esperado status 200 porém foi obtido %d", w.Code)
	}

	// Verifique se o conteúdo do corpo da resposta não está vazio
	if len(w.Body.Bytes()) == 0 {
		t.Error("esperava-se conteúdo no corpo da resposta, mas estava vazio")
	}
}
