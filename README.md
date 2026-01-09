
# Platform Game

Projeto de um sistema para gerenciamento de usuários, pontuações e ranking de um jogo de plataforma, desenvolvido em Go.

## Estrutura do Projeto

```
platform-game/
├── go.mod
├── main.go
├── model/
│   ├── user.go
│   ├── score.go
│   └── ranking.go
├── service/
│   ├── user_service.go
│   ├── score_service.go
│   ├── ranking_service.go
│   └── resource_service.go
```

## Descrição dos Componentes

- **main.go**: Ponto de entrada da aplicação. Demonstra o uso dos serviços de usuário, pontuação, ranking e recursos.
- **model/**: Define as estruturas principais:
	- `User`: Usuário do jogo, com validação.
	- `Score`: Pontuação associada a um usuário.
	- `Ranking`: Estrutura para classificação dos usuários.
- **service/**: Implementa a lógica de negócio:
	- `user_service.go`: Filtra e valida usuários.
	- `score_service.go`: Calcula a pontuação total de cada usuário.
	- `ranking_service.go`: Gera o ranking dos usuários com base nas pontuações.
	- `resource_service.go`: Demonstra uso de recursos e defer.

## Funcionalidades

- Cadastro e validação de usuários.
- Registro e soma de pontuações.
- Geração de ranking ordenado por pontos.
- Demonstração de uso de recursos com defer.

## Como Executar

1. Instale o Go ([download](https://golang.org/dl/)).
2. No terminal, navegue até a pasta `platform-game`.
3. Execute:

```bash
go run main.go
```

## Exemplo de Saída

```
Valid Users:
Total score for user John: 25
Total score for user Sarah: 20
Total score for user Sebastian: 0
Recurso Aberto
Processando com Defer
Recurso Fechado
Ranking:
1 lugar - User 1: 25 pontos
2 lugar - User 2: 20 pontos
```

## Requisitos

- Go 1.18 ou superior

