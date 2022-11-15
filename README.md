# Golang com RabbitMQ

[Golang](https://go.dev) é uma linguagem simplesmente fantástica, super performática, intuitiva e com vários benefícios. E por que não utilizá-la para criar aplicações performáticas e resilientes?

Este é o objetivo desse simples projeto! Uma aplicação onde é feita a criação de um pedido através de uma requisição HTTP, porém o processamento é feito de forma assíncrona, usando a plataforma de mensageria [RabbitMQ](https://www.rabbitmq.com) para garantir resiliência.

---

## Multithread

Com Golang é possível criar várias threads para realizar diversas operações. E pensando em explorar esse recurso, ele foi utilizado na criação dos workers que consomem as mensagens para criação dos pedidos. Ou seja, podemos ter a quantidade de workers que desejarmos em nossa aplicação.

Para especificar a quantidade de Workers, basta alterar o parâmetro da função `MakeOrderCreateWorker` no arquivo de [inicialização](src/main.go).

---

## Banco de Dados

Para armazenar os pedidos criados, foi utilizado o [Sqlite3](https://www.sqlite.org/index.html) por ser um banco de dados mais simples e fácil de ser configurado.

Será necessário criar um arquivo `order.db` na raiz no projeto e criar a tabela `orders` para que seja possível armazenar os dados. Segue o script sql para a criação da mesma:

```sql
CREATE TABLE orders (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    date DATETIME DEFAULT CURRENT_TIMESTAMP
)
```

---

## Instalando as dependências

Observação: é necessário ter o Go instalado na máquina.

Para instalar as dependências do projeto, execute o seguinte comando:

```bash
go get ./...
```

---

## Inicializando a aplicação

O projeto usa o RabbitMQ, então será necessário que ele esteja rodando na sua máquina. A string de conexão utilizada é: `amqp://example:123456@localhost:5672/`

Mas pode ser alterada nas seguintes funções: `MakeOrderCreateWorker` e `MakeOrderServer`
<br/>

### Para inicializar a aplicação, execute o seguinte comando:

```bash
go run src/main.go
```

A aplicação será inicializada na porta `3000`


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

---

## Projeto em desenvolvimento!

Este projeto ainda está em fase de desenvolvimento e serão feitas melhorias no algoritmo, além da criação de testes de unidade e melhorias na configuração das variáveis de ambiente e banco de dados. Também será feita a conteinerização da aplicação.