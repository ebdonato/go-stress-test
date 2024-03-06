# Desafio Técnico Stress Test

Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.

Execução do Teste:

Realizar requests HTTP para a URL especificada.
Distribuir os requests de acordo com o nível de concorrência definido.
Garantir que o número total de requests seja cumprido.
Geração de Relatório:

Apresentar um relatório ao final dos testes contendo:
Tempo total gasto na execução
Quantidade total de requests realizados.
Quantidade de requests com status HTTP 200.
Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
Execução da aplicação:
Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
`docker run <sua imagem docker> —url=<http://google.com> —requests=1000 —concurrency=10`

## Aplicação

Para construir a imagem a partir do Dockerfile:

```shell
docker build -t NOME_DA_IMAGEM .
```

Posteriormente para executar:

```shell
docker run --rm -it NOME_DA_IMAGEM —-url=http://google.com —-requests=1000 —-concurrency=10
```

ou

```shell
docker run --rm -it NOME_DA_IMAGEM —u=http://google.com —r=1000 —c=10
```

As opções de entrada de parâmetros que podem ser passadas via linha de comando:

```shell
  -u, --url string        URL to send requests to (default "https://www.fullcycle.com.br")
  -r, --requests int      Number of requests to send (default 100)
  -c, --concurrency int   Number of requests to send concurrently (default 10)
  -h, --help              Prints help
```

|Comando|Atalho|Descrição|Valor Padrão|
|---|---|---|---|
|url|u|URL do serviço a ser testado|<https://www.fullcycle.com.br>|
|requests|r|Número total de requests|100|
|concurrency|c|Número de chamadas simultâneas|10|
|help|h|Exibe ajuda|Falso|
