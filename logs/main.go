package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"time"
	"github.com/diogenesdornelles/logs/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var url = "http://example.com"


func testZapLogger() {
	fmt.Println("####", "Testando logger package", "#####")
    // Inicializa o logger profissional
    err := logger.InitLogger(zapcore.InfoLevel, "app.log")
    if err != nil {
        log.Fatalf("Erro ao inicializar logger: %v", err)
    }

    // Agora use logger.Sugar para logs
    logger.Sugar.Infow("Aplicação iniciada",
        "version", "1.0",
        "env", "production",
    )
    logger.Sugar.Errorf("Erro exemplo: %s", "detalhes do erro")

}


func zapLibrary() {
	fmt.Println("####", "zap library terminal", "#####")
	loggerTerminal, _ := zap.NewProduction()
	defer loggerTerminal.Sync() // flushes buffer, if any
	sugarTerminal := loggerTerminal.Sugar()
	sugarTerminal.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugarTerminal.Infof("Failed to fetch URL: %s", url)

	fmt.Println("####", "zap library file", "#####")
	// Abrir arquivo para logs do Zap
	file, err := os.OpenFile("zap.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log do Zap: %v", err)
	}
	defer file.Close()
	
	// Configurar Zap para escrever no arquivo
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// Multioutput permite enviar logs para múltiplos destinos (ex: console + arquivo) usando io.MultiWriter. Isso é útil para ter logs persistentes e visíveis durante a execução.
func multiOutput() {
	fmt.Println("####", "multi output", "#####")
	// 1. Abrir (ou criar) o arquivo. O flag O_APPEND evita apagar o que já existe.
	// O_RDWR: leitura/escrita | O_CREATE: cria se não existir
	file, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}
	defer file.Close()

	// Cria um escritor que replica a mensagem para ambos os destinos
	multi := io.MultiWriter(os.Stdout, file)

	log.SetOutput(multi)
	log.Println("Log visível em dois lugares!")
}

func structuredLogging() {
	fmt.Println("####", "structured logging terminal", "#####")
	// Configura o logger para produzir JSON
	// A opção nil usa as configurações padrão, mas podemos personalizar (ex: timestamps, níveis, etc.)
	// opção os.Stdout para enviar para console, ou um arquivo para persistência
	handlerTerminal := slog.NewJSONHandler(os.Stdout, nil)
	loggerTerminal := slog.New(handlerTerminal)

	// Exemplo de log estruturado usando log/slog
	loggerTerminal.Error("Conexão falhou",
		"database", "postgres",
		"attempt", 3,
		"error", "timeout",
	)
	fmt.Println("####", "structured logging file", "#####")
	// Abrir arquivo para logs estruturados
	file, err := os.OpenFile("structured.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log estruturado: %v", err)
	}
	defer file.Close()

	// Configura o handler para escrever no arquivo em JSON
	handlerFile := slog.NewJSONHandler(file, nil)
	loggerFile := slog.New(handlerFile)

	// Exemplo de log estruturado
	loggerFile.Error("Conexão falhou",
		"database", "postgres",
		"attempt", 3,
		"error", "timeout",
	)
}

func loggingWithContext() {
	fmt.Println("####", "logging with context", "#####")
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	// Criando um logger com contexto adicional
	ctxLogger := logger.With(
		"request_id", "abc123",
		"user_id", 42,
	)

	ctxLogger.Info("Processando requisição com contexto")

}

func main() {
	// 1. Abrir (ou criar) o arquivo. O flag O_APPEND evita apagar o que já existe.
	// O_RDWR: leitura/escrita | O_CREATE: cria se não existir
	file, err := os.OpenFile("system.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}
	defer file.Close()

	// 2. Definir o arquivo como o destino do log global
	log.SetOutput(file)

	log.Println("Este log será salvo no arquivo system.log, não no console!")

	// Se quiser desabilitar completamente o log, redirecionamos para io.Discard
	// log.SetOutput(io.Discard)

	multiOutput()
	structuredLogging()
	loggingWithContext()
	zapLibrary()
	testZapLogger()
}
