# GORM

## Fluxo Cli

- Criar os models com structs: internal/domain/models
- Setar configuração com genconfig.Config: internal/domain/config
- Criar as queries: internal/domain/query
- gorm gen -i ./internal/domain/models: gera codigo em internal/domain/generated
- como migrar para o banco com CLI???
- como fazer o controle de migrações CLI???