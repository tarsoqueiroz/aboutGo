# Go configuration with Viper

## About

> `https://github.com/spf13/viper`

Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:

- setting defaults
- reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
- live watching and re-reading of config files (optional)
- reading from environment variables
- reading from remote config systems (etcd or Consul), and watching changes
- reading from command line flags
- reading from buffer
- setting explicit values

## Environment setup for test

```sh
# Golang struct path
mkdir cmd internal pkg api web configs init scripts deployments test docs tools
```

## First test

> [Carregando configurações com Viper](https://aprendagolang.com.br/2021/11/04/carregando-configuracoes-com-viper/)

```sh
# Initializing project
go mod init confviper

# Installing depencies
go get github.com/spf13/viper

# Edit files .GO and ...
go run main.go
## or
go run ./cmd/first.go
```

## Handling Go configuration with Viper

> [Handling Go configuration with Viper](https://blog.logrocket.com/handling-go-configuration-viper/)

```sh
# Create the .env file on root path

# Edit file ./cmd/env.go

# Run file
go run ./cmd/env.go
```
