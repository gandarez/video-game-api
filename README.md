# Video Game API

API para gerenciar consoles, cartuchos e outros itens relacionados a video-games.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem principal do projeto.
- **Docker** e **Docker Compose**: Para execução e integração de ambientes.
- **Kubernetes (Kind)**: Para testes e execução em cluster local.
- **PostgreSQL**: Banco de dados relacional.
- **Redis**: Cache.
- **gRPC**: Comunicação entre serviços.
- **Protobuf**: Definição de APIs gRPC.
- **golangci-lint**: Ferramenta de lint para Go.

## Pré-requisitos

- [Go >= 1.24](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)
- [protoc](https://grpc.io/docs/protoc-installation/)
- (Opcional) [Kind](https://kind.sigs.k8s.io/) para rodar localmente em cluster Kubernetes

## Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis (ajuste conforme seu ambiente):

```env
SERVICE_NAME=video-game-api
SERVER_PORT=17020
GRPC_PORT=17022
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=usuario
DATABASE_PASSWORD=senha
DATABASE_NAME=video_game
REDIS_HOST=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
VENDOR_IGDB_HOST=https://api.igdb.com
VENDOR_TWITCH_HOST=https://id.twitch.tv
VENDOR_TWITCH_CLIENT_ID=seu_client_id
VENDOR_TWITCH_CLIENT_SECRET=seu_client_secret
```

## Como Executar Localmente

### 1. Via Docker Compose

```bash
docker compose up --build
```

A API estará disponível em `http://localhost:17020`.

### 2. Compilando e rodando local (Go)

```bash
make build
./build/video-game-api-linux-amd64
```

> Certifique-se de que o banco de dados e o Redis estejam rodando e configurados conforme o arquivo `.env`.

### 3. Usando Makefile

Principais comandos:

- `make build`: Compila o binário para seu SO.
- `make lint`: Executa o linter.
- `make test`: Executa os testes unitários.
- `make test-integration`: Executa testes de integração via Docker Compose.
- `make build-docker-image`: Gera a imagem Docker do serviço.

### 4. Kubernetes (Kind)

Para rodar em um cluster Kind local:

```bash
make kind-up
make build-docker-image
make kind-load
make kind-apply
```

## Endpoints

- REST: `http://localhost:17020/consoles`
- Health checks: `/liveness`, `/readiness`
- gRPC: Porta `17022` (default)

## Licença

BSD 3-Clause License - Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
Desenvolvido por Carlos Henrique Guardão Gandarez
