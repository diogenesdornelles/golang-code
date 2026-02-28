package signin

import (
	"html/template"
	"net/http"
	"time"
	"github.com/diogenesdornelles/web-server-example/internal/domain/signin"
	"github.com/diogenesdornelles/web-server-example/internal/domain/page"
)


func SignInHandler(w http.ResponseWriter, r *http.Request) {


	tmpl := template.Must(template.ParseFiles("templates/signin.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}


func PostSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}
	Email := r.FormValue("email")
	Password := r.FormValue("password")

	SignedUser := signin.User{
		Email: Email,
		Password: Password,
	}

	data := page.HomePage{
		Title:       "Página Inicial",
		Message:     "Bem-vindo ao site!",
		DataCriacao: time.Now(),
		IsSignedIn:  bool(SignedUser.Email != "" && SignedUser.Password != ""),
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}
