# Projeto de Integração Golang com Gemini SDK: Explorando o Poder da IA do Google

Este projeto de exemplo tem como objetivo demonstrar, de forma prática e direta, como integrar a linguagem de programação Golang com o Gemini, a poderosa ferramenta de Inteligência Artificial desenvolvida pelo Google, utilizando a SDK oficial.

## Desvendando o Gemini: A Inteligência Artificial do Google ao Seu Alcance

O Gemini representa a mais recente inovação do Google no campo da Inteligência Artificial. Essa plataforma versátil oferece uma ampla gama de recursos avançados, incluindo:

* **Reconhecimento de Voz:** Permite a interação com sistemas através da fala.
* **Processamento de Linguagem Natural (PLN):** Capacidade de entender e gerar texto em linguagem humana.
* **Visão Computacional:** Habilidade de interpretar e analisar informações visuais.
* **E muito mais!**

Este projeto focará em como utilizar algumas dessas capacidades dentro de uma aplicação Golang.

## Preparando o Terreno: Pré-requisitos Essenciais

Antes de mergulhar no código, certifique-se de que o seu ambiente de desenvolvimento possui as seguintes ferramentas instaladas e configuradas:

* **Golang:** A linguagem de programação que impulsiona este projeto.
    * **Download e Instruções de Instalação:** Você pode encontrar o guia completo para instalação no site oficial: [https://golang.org/dl/](https://golang.org/dl/)

* **SDK do Gemini para Golang:** A biblioteca que facilita a comunicação entre sua aplicação Golang e os serviços do Gemini.
    * **Documentação e Instruções de Instalação:** Para instalar e configurar a SDK corretamente, siga as instruções detalhadas na documentação oficial: [https://ai.google.dev/gemini-api/docs/sdks?hl=pt-br#go-quickstart](https://ai.google.dev/gemini-api/docs/sdks?hl=pt-br#go-quickstart)

## Mãos à Obra: Instalação e Configuração do Projeto

Siga estes passos para configurar o projeto em sua máquina local:

1.  **Clone o Repositório:** Utilize o Git para copiar o código fonte do projeto para o seu computador. Abra seu terminal ou prompt de comando e execute:

    ```bash
    git clone [https://github.com/br4tech/go-with-gemini.git](https://github.com/br4tech/go-with-gemini.git)
    cd go-with-gemini
    ```

2.  **Gere o Código Wire:** O Wire é uma ferramenta para injeção de dependências em Go. Execute o seguinte comando para gerar o código necessário:

    ```bash
    go run [github.com/google/wire/cmd/wire](https://github.com/google/wire/cmd/wire)
    ```

3.  **Execute a Aplicação:** Você tem duas opções para executar o projeto, utilizando Docker ou diretamente no seu ambiente local:

    **a) Utilizando Docker (Recomendado para um ambiente isolado):**

    Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina. Execute os seguintes comandos na raiz do projeto:

    ```bash
    docker-compose build
    docker-compose up app
    ```

    Este comando irá construir a imagem Docker e iniciar o container da aplicação.

    **b) Executando Sem Docker (Requer PostgreSQL instalado localmente):**

    Se você optar por não usar o Docker, siga estes passos:

    * **Inicie o Banco de Dados PostgreSQL:** Assumindo que você tem o PostgreSQL instalado e configurado, inicie o servidor. O comando abaixo inicia um container Docker do PostgreSQL para facilitar, caso você não tenha um rodando localmente:

        ```bash
        docker run --name gemini -e POSTGRES_PASSWORD=123456 -d -p 5434:5432 postgres
        ```

    * **Crie o Banco de Dados:** Conecte-se ao servidor PostgreSQL e crie o banco de dados `geminidb`. Se você usou o container Docker, pode fazer isso com os seguintes comandos:

        ```bash
        docker exec -it gemini bash
        psql -U postgres
        CREATE DATABASE geminidb;
        \q # Para sair do psql
        exit # Para sair do container
        ```

    * **Execute a Aplicação Golang:** Finalmente, execute a aplicação Golang:

        ```bash
        go run cmd/wire_gen.go cmd/main.go
        ```

## Interagindo com a Aplicação: Criando um Produto

Com a aplicação em execução (geralmente na porta `8080`), você pode interagir com ela para criar novos produtos. Utilize o `curl` no seu terminal para enviar uma requisição HTTP POST:

```bash
curl --location 'localhost:8080/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Nome do Produto",
    "code": "CODIGO_UNICO",
    "image": "URL_DA_IMAGEM_DO_PRODUTO"
}'
```

Substitua os valores entre as aspas com as informações do produto que você deseja criar. Por exemplo:

```bash
curl --location 'localhost:8080/product' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Produto Fantástico",
    "code": "PROD001",
    "image": "[https://exemplo.com/imagem_produto.jpg](https://exemplo.com/imagem_produto.jpg)"
}'
```

## Explorando a API com o Postman

Para facilitar a exploração e o teste da API do projeto, uma Collection do Postman está disponível. Você pode importá-la utilizando o seguinte link:

[https://app.getpostman.com/join-team?invite_code=b90555d01a13ab0e308495a2e081d08a4434bf28414e9f3aeb735f4e5bd2f3ff](https://app.getpostman.com/join-team?invite_code=b90555d01a13ab0e308495a2e081d08a4434bf28414e9f3aeb735f4e5bd2f3ff)

Ao acessar este link, você poderá importar a Collection para o seu Postman e encontrar requisições de exemplo para interagir com a API de forma mais visual e organizada.