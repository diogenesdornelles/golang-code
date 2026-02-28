# GO Lessons

## Instalação

- [Go homepage](https://go.dev/doc/install)

1) Baixar .deb
2) sudo tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz
3) nano ~/.bashrc
4) incluir:
```bash
# Configuração do Go
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
```
5) aplicar
```bash
source ~/.bashrc
```

## Ver mais 

- [Effective Go](https://go.dev/doc/effective_go)


## bash

- Adicionar Zsh + Oh My Zsh e plugins:

```bash
sudo apt update && sudo apt install zsh git curl -y
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

```


- Edite arquivo: `nano ~/.zshrc` garantindo que as linhas abaixo estejam ativas:
```bash

# Plugins recomendados para produtividade
plugins=(
    git
    docker
    python
    sudo
    zsh-autosuggestions
    zsh-syntax-highlighting
)

source $ZSH/oh-my-zsh.sh

##### Go #####
# Go - Configuração 
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin

# Go - Aliases
alias gom="go mod tidy"
alias gof="gofmt -w ."

# Go - Função para rodar projeto
gor() {
    if [ -z "$1" ]; then
        go run .
    else
        go run "$1"
    fi
}

# Go - Função para buildar projeto
gob() {
    if [ -z "$1" ]; then
        go build .
    else
        go build "$1"
    fi
}

# Go - Função para iniciar projeto com meu usuário
goi() {
    if [ -z "$1" ]; then
        echo "Erro: Você precisa passar o nome do projeto. Ex: goi meu-projeto"
    else
        go mod init github.com/diogenesdornelles/"$1"
    fi
}

# Go - Função para instalar uma lib externa
gog() {
    if [ -z "$1" ]; then
        echo "Erro: Você precisa passar a URL da lib. Ex: goi lib-terceiro"
    else
        go get "$1"
        go mod tidy
    fi
}

##############
# Oh My Zsh (adicionado pelo script de instalação)
export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="robbyrussell"  # Ou outro tema, como "agnoster"

```
- Salve mudanças: `source ~/.zshrc` 


## CLI

- **go run /nome/.go**: Roda arquivo sem salvar o compilador;
- **go build /nome/.go**: Gerar o compilado do arquivo;
- **go mod init github.com/diogenesdornelles/nome**: Criar novo projeto;
- **go get /url-do-repositorio/**: Adicionar Lib;
- **go mod tidy**: Remover lixo / Sincronizar / Usar ao instalar nova lib;
- **go run .**: Rodar o projeto;
- **go build .**: Gerar o compilado do projeto;

## Watch mode

```bash
go install github.com/air-verse/air@latest
```
```bash
air init
```
```bash
air
```


## URL de libs:

| Nome | Finalidade | URL (via get) |
|:---|:---:|---:|
|Fiber|Framework HTTP (estilo Express/NestJS) rápido e simples.|github.com/gofiber/fiber/v2|
|GORM|ORM para mapeamento de banco de dados (MySQL/PostgreSQL/SQLite).|gorm.io/gorm|
|Stripe Go|SDK Oficial para pagamentos, assinaturas e checkout.|github.com/stripe/stripe-go/v81|
|Validator|Validação de structs (checar se e-mail é válido, campos obrigatórios, etc).|github.com/go-playground/validator/v10|
|UUID|Geração de IDs únicos universais para Ingressos e Transações.|github.com/google/uuid|
|Zap|Logger ultra-rápido para monitorar erros e vendas em tempo real.|go.uber.org/zap|
|Godotenv|Gerenciamento de variáveis de ambiente (Tokens do Pix e Stripe).|github.com/joho/godotenv|
|Golang JWT|Autenticação para diferenciar Compradores de Vendedores.|github.com/golang-jwt/jwt/v5|
|Argon2 / Bcrypt|Hash de senhas (segurança nativa do Go).|golang.org/x/crypto/argon2|
|Zitadel|entrega uma interface de login, gestão de usuários, login social (Google, GitHub) e RBAC (Role-Based Access Control) nativo|go get -u github.com/zitadel/zitadel-go/v3|
|AWS S3|O padrão industrial. Extremamente robusto.|github.com/aws/aws-sdk-go-v2|
|Cloudflare R2|(Usa o SDK da AWS) Sem taxa de saída, o que economiza muito dinheiro.|github.com/aws/aws-sdk-go-v2|
|Supabase Storage|Excelente se você quer algo fácil e rápido como o Better Auth.|github.com/supabase-community/storage-go|
||||

## Keywords

Go has 25 keywords like if and switch that may be used only where the syntax permits; they can’t be used as names: 
- break default func interface select
- case defer go map struct
- chan else goto package switch
- const fallthrough if range type
- continue for import return var

## Predeclared names

In addition, there are about three dozen **predeclared** names like int and true for bui lt-in constants, types and functions:

- **Constants**: true false iota nil

- **Types**: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64 bool byte rune string error

- **Functions**: make len cap new append copy close delete complex real imag panic recover

Thes e names are not reserved, so you may use them in declarations.

## Padrões de nomenclatura 

1. **Pacotes e Pastas (Packages & Directories)**
Em Go, o nome da pasta e o nome do pacote devem ser iguais. Sempre minúsculas. Nome simples e descritivo. Sem plural e sem nomes genéricos. Hífen é proibido em pacotes. Coisas como user_manager é desencorajado. Se usar, prefira user ou usermgt

2. **Arquivos (Files)**
Minúsculas e usando snake case, se composto.

3. **Variáveis, Constantes, Funções e Métodos**
CamelCase para exportações e pascalCase para uso no próprio pacote

4. **Acrônimos (O "Estilo ID")**
Essa é uma regra que confunde muitos iniciantes. Se um nome contém um acrônimo (como API, ID, HTTP, URL), ele deve manter a capitalização consistente.

5. **Interfaces**
A regra do "er" (Reader, Writer, Closer) é focada em interfaces de propósito único. Elas são a base da composição em Go. 


## Estrutura API (View opcionalmente)

Árvore de pastas e arquivos:

```text
project/
├── cmd/
│   └── api/
│       └── main.go           # Ponto de entrada (Bootstrap da aplicação)
├── internal/
│   ├── api/
│   │   ├── handler/         # Controllers (recebem HTTP, validam entrada)
│   │   ├── middleware/      # Auth, Logging, CORS
│   │   └── route/           # Definição das rotas do Fiber
│   ├── domain/               # O CORAÇÃO: Structs de dados e Interfaces
│   ├── repository/           # Camada de Dados (GORM / Consultas SQL)
│   ├── service/              # Regras de Negócio (Onde a mágica acontece)
│   └── platform/             # Configurações (Banco de dados, Stripe, Pix)
├── pkg/                      # Código utilitário que pode ser exportado
├── public/                   # Arquivos públicos em dev mode
|__ templates/                # Templates html
├── .env                      # Variáveis sensíveis (NÃO comitar)
├── .gitignore
├── go.mod
├── go.sum
└── README.md

```

Comando para iniciar:

```bash
mkdir -p cmd/api internal/api/handler internal/api/middlewar internal/api/route internal/domain internal/repository internal/service internal/platform pkg public templates
touch .env .gitignore cmd/api/main.go internal/platform/config.go
sudo chown -R $USER:$USER .
```

## Static Files ou Cloud

Desenvolvimento use static

```go
app.Static("/static", "./public")
```

Produção cloud storage

Não coloque a lógica de upload dentro do seu Handler (Controller). Crie um "Provider" ou "Uploader" na pasta internal/platform.

O fluxo ideal:

O Handler recebe o arquivo (c.FormFile("imagem")).

O Handler passa o arquivo para o Service.

O Service chama o S3/Cloudflare Platform para fazer o upload.

O S3 retorna a URL (ex: https://cdn.evento.com/foto.jpg).

O Service salva essa URL no banco via Repository.

## Ferramentas de Observabilidade:



[Observabilidade](/https://marcospereirajr.com.br/observability-with-open-telemetry-and-golang-8ae2d9e994af)


## Intellisense, Debugger, Linter

Para ter a melhor experiência possível, você só precisa de uma coisa:

- Instale a extensão oficial "Go" (da Google) no VS Code.

- Assim que abrir um arquivo .go, o VS Code vai perguntar no canto inferior direito: "Analysis tools missing. Install?".

- Clique em Install All.

Isso vai baixar o gopls (Intellisense), delve (Debugger) e o staticcheck (Linter).