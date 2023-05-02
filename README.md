# Clean Architecture Go

## Como funciona

Sistema faz a inserção e listagem de orders contendo o id, price, tax e final price, podendo ser realizada a partir
de diferentes methodos de inserção: Endpoint REST, gRPC e GraphQL. Além de adicionar cada nova inserção na fila do 
RabbitMQ para que outro serviço possa ler. O banco de dados utilizado é o MySQL.

## Como utilizar

Primeiramente suba o banco o cantâiner do docker:

```bash
docker compose up -d
```

Realize a configuração da fila do RabbitMQ e crie a tabela no banco de dados.

Configurações do RabbitMQ e mostrar a mensagem:

```text
Entrar em http://localhost:15672 (Login: guest | Password: guest)
Queues -> Add new queue Name: orders -> Add queue
orders -> bindings -> from exchange: amq.direct -> bind
Get messages -> Get Message(s) -> mensagem!
```

Query para criação da tabela no banco de dados:

```sql
CREATE TABLE orders
(
    id          varchar(255) NOT NULL,
    price       float        NOT NULL,
    tax         float        NOT NULL,
    final_price float        NOT NULL,
    PRIMARY KEY (id)
);
```

## Criar nova mensagem

### Endpoint Rest (POST /order)

Faça um request http contendo as seguintes informações:

URL (POST): http://localhost:8000/order \
Header: Content-Type: application/json

```json
{
  "id": "z",
  "price": 100.5,
  "tax": 0.5
}
```

e receberá a resposta:
```json
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

{"id":"al","price":100.5,"tax":0.5,"final_price":101}
```

### Service CreateOrder gRPC

Levando em conta que possui o evans instalado:

```bash
evans -r repl

# utilizar os comandos:
call CreateOrder
id (TYPE_STRING) => zz
price (TYPE_FLOAT) => 1
tax (TYPE_FLOAT) => 2

# receberá a resposta
{
  "finalPrice": 3,
  "id": "zz",
  "price": 1,
  "tax": 2
}
```

### Mutation createOrder GraphQL

Entre na url: http://localhost:8080 e adicionar a seguinte mutation:

```graphql
mutation createOrder {
    createOrder(input:{id:"zzz", Price: 10, Tax: 2}) {
        id
        Price
        Tax
        FinalPrice
    }
}
```

e irá receber a resposta:

```json
{
  "data": {
    "createOrder": {
      "id": "zzz",
      "Price": 10,
      "Tax": 2,
      "FinalPrice": 12
    }
  }
}
```

## Listar todas as mensagens

### Endpoint Rest (GET /order)

Faça um request http contendo as seguintes informações:

URL (GET): http://localhost:8000/order

e irá receber a resposta:
```json
HTTP/1.1 200 OK
Content-Type: application/json

[
  {
    "id": "z", 
    "price": 100.5,
    "tax": 0.5
  },
  {
    "id": "zz",
    "price": 100.5,
    "tax": 0.5
  },
  {
    "id": "zzz",
    "price": 100.5,
    "tax": 0.5
  }
]
```

### Service ListOrders com gRPC

Levando em conta que possui o evans instalado:

```bash
evans -r repl

# utilizar o comando:
call ListOrders

# receberá a resposta:
{
  "orders": [
    {
      "id": "z", 
      "price": 100.5,
      "tax": 0.5
    },
    {
      "id": "zz",
      "price": 100.5,
      "tax": 0.5
    },
    {
      "id": "zzz",
      "price": 100.5,
      "tax": 0.5
    }
  ]
}
```

### Query ListOrders GraphQL

Entre na url: http://localhost:8080 e adicionar a seguinte query:
```graphql
query listOrders {
    ListOrders {
        id
        Price
        Tax
        FinalPrice
    }
}
```

e irá receber a resposta:
```json
{
  "data": {
    "ListOrders": [
      {
        "id": "z",
        "price": 100.5,
        "tax": 0.5
      },
      {
        "id": "zz",
        "price": 100.5,
        "tax": 0.5
      },
      {
        "id": "zzz",
        "price": 100.5,
        "tax": 0.5
      }
    ]
  }
}
```