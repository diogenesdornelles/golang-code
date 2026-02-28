## SQLC ORM + PG + Testes + Migrations

### Instalação
- dotenv: `go install github.com/joho/godotenv`
- driver: `go install github.com/jackc/pgx/v5`
- querybuilder: `go install github.com/Masterminds/squirrel`
- migrations: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- testes: `go install github.com/pashagolub/pgxmock`
- Observabilidade: `go install go.opentelemetry.io/otel`
- sqlc: `sudo snap install sqlc` ou `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

### Configuração Inicial
- Preencher schema em `db/temp/schema.sql`
- Preencher queries em `db/queries/query.sql`
- Criar `sqlc.yaml` em root e configurar caminhos para arquivos e código gerado
- Gerar código Go: `sqlc generate`

### Migrations
- Gerar SQL de migração inicial: `migrate create -ext sql -dir db/migration -seq init_schema`
- Jogar código de criação de tabelas em `.up.sql` e rollback em `.down.sql`
- Fazer migração: `migrate -path ./db/migration -database "postgres://postgres:1234@localhost:5432/testdb?sslmode=disable" up`
- Para futuras migrações (ex.: adicionar coluna): `migrate create -ext sql -dir ./db/migration -seq add_age_column`
- Jogar código de alteração em `.up.sql` e rollback em `.down.sql` (usar `ALTER TABLE`)
- Rodar migração: `migrate -path ./db/migration -database "postgres://postgres:1234@localhost:5432/testdb?sslmode=disable" up`
- Para migração específica: `migrate -path ./db/migration up/down <numero>`

### Queries e Schema
- Após mudanças no banco, atualizar manualmente as queries em `db/queries` (ex.: incluir coluna `age`)
- Atualizar schema em `db/temp/schema.sql` (ex.: adicionar `age` na tabela `authors`)
- Regenerar código Go: `sqlc generate`
- Lembrete: Código Go (`sqlc generate`) deve estar sincronizado com o banco (`migrate`)

### Testes
- Usar pgxmock para mockar queries em testes unitários (evita dependência de DB real)
- Exemplo básico em teste:
  ```go
  func TestCreateAuthor(t *testing.T) {
      mock, err := pgxmock.NewPool()
      if err != nil {
          t.Fatal(err)
      }
      defer mock.Close()

      mock.ExpectExec("INSERT INTO authors").WithArgs("John Doe").WillReturnResult(pgxmock.NewResult("INSERT", 1))

      // Seu código de teste aqui, usando mock em vez de conexão real
  }
  ```