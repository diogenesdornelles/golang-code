package home

import (
    "net/http"
    "github.com/diogenesdornelles/web-server-example/internal/web/handler/home"
)

func ConfigureRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/", home.RootHandler)
}