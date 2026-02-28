package home

import (
	"html/template"
	"net/http"
	"time"

	"github.com/diogenesdornelles/web-server-example/internal/domain/page"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	data := page.HomePage{
		Title:       "Página Inicial",
		Message:     "Bem-vindo ao site!",
		DataCriacao: time.Now(),
		IsSignedIn:  false,
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}