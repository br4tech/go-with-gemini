# Projeto de Integração Golang com Gemini SDK

Este é um projeto de exemplo que demonstra como integrar Golang com o Gemini, utilizando a SDK fornecida pelo Google.

## Sobre o Gemini

O Gemini é uma nova ferramenta de Inteligência Artificial desenvolvida pelo Google. Ele oferece uma variedade de recursos avançados de IA, incluindo reconhecimento de voz, processamento de linguagem natural e muito mais.

## Pré-requisitos

Antes de começar, certifique-se de ter instalado o seguinte em sua máquina:

- Golang: [Download e instruções de instalação](https://golang.org/dl/)
- SDK do Gemini: [Documentação e instruções de instalação](https://gemini.google.com/sdk)

## Instalação e Configuração

1. Clone este repositório para o seu ambiente local:

```bash
git clone https://github.com/br4tech/go-with-gemini.git
cd go-with-gemini

```


2. Gerar o arquivo do wire:

```bash
go run github.com/google/wire/cmd/wire

```

3. Executar aplicacao:
 
```bash
 go run cmd/wire_gen.go cmd/main.go

```

Obs: 

Caso nao tenha o banco criado execute os comando, abaixo antes de tudo:

```bash

  docker run --name gemini -e POSTGRES_PASSWORD=123456 -d -p 5434:5432 postgres
  
  docker exec -it gemini bash

  psql -U postgres

  CREATE DATABASE geminidb;