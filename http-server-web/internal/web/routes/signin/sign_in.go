package signin

import (
    "net/http"
    "github.com/diogenesdornelles/web-server-example/internal/web/handler/signin"
)

func ConfigureRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /signin", signin.SignInHandler)
    mux.HandleFunc("POST /submit", signin.PostSubmit)
}