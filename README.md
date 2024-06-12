# Fastx Exchange

### Pre-requisites

- Go >= 1.2
- PostgreSQL 16
- Node JS >= 20
- Redis
- Docker
- make

### Need to install some packages before setting project
- [Golangci-lint](https://golangci-lint.run/welcome/install/): Used to detect errors and check code style in Go code.
- [Atlas](https://atlasgo.io/docs):Use Handle for Database Migrations
- [Air](https://github.com/air-verse/air): Auto reload when code change

### First time setup
- install dotenv
```sh
$ npm install -g dotenv-cli
```


- copy .env file from .env.example and then change database configuration
```sh
$ cp .env.example .env
```

### Install app
```sh
$ go mod tidy
```
- Create Migration file from Model
```sh
$ make diff_migration
```
- Create Migration file
```sh
$ make create_migration
```

- Run seed data 
```sh
$ make db-seed
```

### Start app
```sh
$ make dev
```