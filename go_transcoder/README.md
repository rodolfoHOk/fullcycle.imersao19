# Imersão FullCycle 19 - Plataforma de streaming de vídeos

## Descrição

Repositório do microsserviço Golang (conversor de vídeos)

## Tecnologias utilizadas

- Go Lang
- PostgreSQL
- Docker

## Rodar a aplicação

Levante os containers:

```bash
docker-compose up -d
```

Através do PGAdmin, crie as tabelas do arquivo `db.sql` no banco de dados `converter`.

Acesse o container o rodando o comando:

```bash
docker compose exec -it go_app_dev bash
```

Use os comandos `go run .cmd/splitchunks/main.go` e `go run .cmd/videoconverter/main.go` para rodar a aplicação, conforme a aula.

# Guia de comandos utilizados

- go mod init github.com/rodolfoHOk/fullcycle.imersao19/go_transcoder
- go run ./cmd/videoconverter/main.go
- docker-compose up -d
- go mod tidy

## Readme principal

- [README.md](../README.md)
