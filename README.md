# Platform Game

Este é um projeto de um jogo de plataforma simples desenvolvido em Go.

## Estrutura do Projeto

```
platform-game/
├── go.mod
├── main.go
├── model/
│   ├── score.go
│   └── user.go
├── service/
│   ├── score_service.go
│   └── user_service.go
```

## Descrição

- `main.go`: Ponto de entrada da aplicação.
- `model/`: Contém as definições das estruturas de dados principais, como usuário e pontuação.
- `service/`: Implementa a lógica de negócio relacionada a usuários e pontuações.

## Como Executar

1. Certifique-se de ter o Go instalado ([download](https://golang.org/dl/)).
2. No terminal, navegue até a pasta `platform-game`.
3. Execute o comando:

```bash
go run main.go
```

## Requisitos
- Go 1.18 ou superior

