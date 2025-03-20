# Projeto de Integração Golang com Gemini SDK

Este é um projeto de exemplo que demonstra como integrar Golang com o Gemini, utilizando a SDK fornecida pelo Google.

## Sobre o Gemini

O Gemini é uma nova ferramenta de Inteligência Artificial desenvolvida pelo Google. Ele oferece uma variedade de recursos avançados de IA, incluindo reconhecimento de voz, processamento de linguagem natural e muito mais.

## Pré-requisitos

Antes de começar, certifique-se de ter instalado o seguinte em sua máquina:

- Golang: [Download e instruções de instalação](https://golang.org/dl/)
- SDK do Gemini: [Documentação e instruções de instalação](https://ai.google.dev/gemini-api/docs/sdks?hl=pt-br#go-quickstart)

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

Com docker:

```bash
 docker-compose build

 docker-compose up app

```

Sem docker
 
```bash
 docker start gemini

 go run cmd/wire_gen.go cmd/main.go

```

Obs: 

Caso nao tenha o banco criado execute os comando, abaixo antes de tudo:

```bash

  docker run --name gemini -e POSTGRES_PASSWORD=123456 -d -p 5434:5432 postgres
  
  docker exec -it gemini bash

  psql -U postgres

  CREATE DATABASE geminidb;
  ```

4. Para criar um produto:

```bash
curl --location 'localhost:8080/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Produto 2",
    "code": "XXX2",
    "image": "https://www.google.com/url?sa=i&url=https%3A%2F%2Fprotelimp.com.br%2Fproduto%2Fcasa-perfume-agradable-500ml%2F&psig=AOvVaw3TU3W8R725EEPsNTzGP8VK&ust=1715091303808000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCOCnzJCb-YUDFQAAAAAdAAAAABAE"
}'
```
