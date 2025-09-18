# Desafio Pós Go Expert Clean Architecture

## Sumário

- [Executando o projeto](#executando-o-projeto)
- [Requisitos](#requisitos)

### Executando o projeto

[[Sumário](#sumário)]

Abra um terminal no diretório do projeto e execute:

```bash
# Subir o mysql e o rabbitmq
docker compose up -d

# Entrar no diretório:
cd cmd/ordersystem

# Rodar o projeto
go run main.go wire_gen.go
```

Serviços do servidor

| Tipo    | Porta |
| ------- | ----- |
| Web     | 8000  |
| GRPC    | 50051 |
| GraphQL | 8080  |

### Requisitos

[[Sumário](#sumário)]

- [x] Automigrations;
- [x] Adicionar a listagem de orders no diretório api;
- [x] Criar o usecase de listagem das orders;
- [x] Disponibilizar no endpoint REST GET /orders;
- [x] Disponibilizar no GRPC como ListOrders;
- [x] Disponibilizar a listagem das orders como query no GraphQL.
