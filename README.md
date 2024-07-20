# API Rest Celebrities

Este é um projeto de um CRUD API REST com o objetivo de estudos, para gerenciamento de celebridades, desenvolvido em Go, que se conecta a um banco de dados PostgreSQL.

## Descrição

A API REST Celebrities permite o cadastro, atualização, remoção e consulta de informações sobre celebridades. Este projeto utiliza Go, Gorilla Mux para roteamento, GORM para ORM e PostgreSQL como banco de dados.

## Bibliotecas Utilizadas

- **github.com/gorilla/mux**: Um roteador e despachante de solicitações HTTP poderoso e flexível.
- **github.com/jackc/pgx/v5**: Um driver PostgreSQL e biblioteca de conexões.
- **gorm.io/gorm**: O ORM (Object Relational Mapping) para Go, usado para facilitar a interação com o banco de dados PostgreSQL.
- **gorm.io/driver/postgres**: Driver PostgreSQL para GORM.
- **github.com/jackc/pgpassfile**: Biblioteca para lidar com arquivos pgpass.
- **github.com/jackc/pgservicefile**: Biblioteca para lidar com arquivos de serviço PostgreSQL.
- **github.com/jinzhu/inflection**: Biblioteca para manipulação de inflexões em strings.
- **github.com/jinzhu/now**: Biblioteca para manipulação de datas e horários.
- **golang.org/x/crypto**: Coleção de pacotes criptográficos.
- **golang.org/x/sync**: Pacotes suplementares para sincronização.
- **golang.org/x/text**: Manipulação de texto (caracteres e strings).

## Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install)

## Instalação

Clone este repositório:

```
git clone https://github.com/rafaelsouza-develop/api-rest-celebridades.git
```
Acesse a pasta do projeto no terminal:

```
cd api-rest-celebridades
```
Para subir o container docker execute o comando a baixo:

```
docker-compose up -d
```

Para subir a aplicação execute o comando a baixo:

```
go run main.go
```
