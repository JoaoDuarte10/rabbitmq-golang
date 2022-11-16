# Golang com RabbitMQ

[Golang](https://go.dev) é uma linguagem simplesmente fantástica, super performática, intuitiva e com vários benefícios. E por que não utilizá-la para criar aplicações performáticas e resilientes?

Este é o objetivo desse simples projeto! Uma aplicação onde é feita a criação de um pedido através de uma requisição HTTP, porém o processamento é feito de forma assíncrona, usando a plataforma de mensageria [RabbitMQ](https://www.rabbitmq.com) para garantir resiliência.

---

## Multithread

Com Golang é possível criar várias threads para realizar diversas operações. E pensando em explorar esse recurso, ele foi utilizado na criação dos workers que consomem as mensagens para criação dos pedidos. Ou seja, podemos ter a quantidade de workers que desejarmos em nossa aplicação.

Para especificar a quantidade de Workers, basta alterar o parâmetro da função `MakeOrderCreateWorker` no arquivo de [inicialização](src/main.go).

---

## Tecnologias Utilizadas

- Conteinerização:
  - [Docker](https://www.docker.com)
- Linguagem:
  - [Golang](https://go.dev)
- Banco de dados:
  - [Sqlite3](https://www.sqlite.org/index.html)
- Mensageria:
  - [RabbitMQ](https://www.rabbitmq.com)
- Logs:
  - [Grafana Loki](https://grafana.com/oss/loki/)
- Dashboards:
  - [Grafana](https://grafana.com)


---

## Inicializando a aplicação

Este projeto utiliza o Docker, então toda a infraestrutura necessária será criada no momento em que os containeres forem iniciados.

### Obervação: É necessário ter o Docker instalado na máquina.

Para criar os containeres, execute o seguinte comando:

```bash
docker-compose up -d
```

<br/>

A aplicação principal será inicializada na porta `3000`

Para acessar o Grafana, acesse: `http://localhost:3333`

---

## Rotas

Rotas existentes na aplicação:

```bash
# Criação de pedidos: 
/order

# Recuperação dos pedidos criados:
/orders
```

Exemplo da rota para criação do pedido:

```bash
curl --request POST \
  --url http://localhost:3000/order \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "",
	"price": ,
	"date": ""
}'
```

Exemplo da rota para recuperação dos pedidos:

```bash
curl --request GET \
  --url http://localhost:3000/orders \
  --header 'Content-Type: application/json'
```
