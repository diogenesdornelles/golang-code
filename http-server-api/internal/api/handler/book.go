package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/diogenesdornelles/http-server-example/internal/domain/page"
	"github.com/diogenesdornelles/http-server-example/internal/domain/user"
)

// Parse templates globalmente no pacote handlers
var Templates = template.Must(template.ParseGlob("templates/*.html"))

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	fmt.Fprintf(w, "Hello, %s %s!", name, surname)
}




func GetPage(w http.ResponseWriter, r *http.Request) {
	data := page.Data{
		Title:       "Página Inicial",
		Message:     "Bem-vindo ao site!",
		DataCriacao: time.Now(),
	}
	if err := Templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}

func PostSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}
	Username := r.FormValue("username")

	// Password := r.FormValue("password")
	// fmt.Fprintf(w, "Username: %s, Password: %s", Username, Password)

	if err := Templates.ExecuteTemplate(w, "lista.html", user.User{Name: Username}); err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}
