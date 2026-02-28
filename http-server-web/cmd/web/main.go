package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/diogenesdornelles/web-server-example/internal/web/routes/home"
	"github.com/diogenesdornelles/web-server-example/internal/web/routes/signin"
	"github.com/diogenesdornelles/web-server-example/internal/platform"
)

var AppPort = ":7080"

/*
Configura o servidor HTTP, registrando as rotas e iniciando o servidor com timeouts adequados para produção.
*/
func main() {

	// Load configurations
    config := platform.LoadConfig()
	AppPort = config.AppPort
	// A ServeMux is a request multiplexer. It matches the URL of each
	// incoming request against a list of registered patterns and calls
	// the handler for the best pattern match.

    // Exemplo: log diferente baseado no modo
    if config.Mode == platform.Dev {
        log.Println("Running in development mode")
    } else {
        log.Println("Running in production mode")
    }

	mux := http.NewServeMux()


	// Registrar rotas
	home.ConfigureRoutes(mux)
	signin.ConfigureRoutes(mux)

	// Configure the HTTP server with timeouts for production readiness
	server := &http.Server{
		Addr:         AppPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Start the server
	fmt.Printf("Server starting on port %s\n", AppPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
