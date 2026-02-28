package logger

import (
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// Logger é o logger principal do Zap (para uso avançado).
var Logger *zap.Logger

// Sugar é o logger sugar do Zap (para logs simples com key-value).
var Sugar *zap.SugaredLogger

// InitLogger configura o Zap para terminal e arquivo simultaneamente.
// - level: nível mínimo de log (ex: zapcore.InfoLevel).
// - filePath: caminho do arquivo de log (ex: "app.log").
// Retorna erro se houver problema na configuração.
func InitLogger(level zapcore.Level, filePath string) error {
    // Configuração do encoder (JSON para estruturação)
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Formato de tempo legível

    // Core para arquivo: escreve em JSON no arquivo
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        return err
    }
    fileCore := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.AddSync(file),
        level,
    )

    // Core para console: escreve em JSON no stdout (terminal)
    consoleCore := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.AddSync(os.Stdout),
        level,
    )

    // Combina os cores com NewTee para multi-destino
    core := zapcore.NewTee(fileCore, consoleCore)

    // Cria o logger principal
    Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
    defer Logger.Sync() // Garante flush ao sair

    // Cria o sugar logger
    Sugar = Logger.Sugar()

    return nil
}